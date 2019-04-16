package jira

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"syscall"
	"text/template"

	"github.com/trivago/tgo/tcontainer"

	"github.com/DATA-DOG/godog/gherkin"

	"golang.org/x/crypto/ssh/terminal"

	"github.com/cbush06/kosher/common"
	"github.com/cbush06/kosher/config"
	"github.com/cbush06/kosher/report"

	"github.com/andygrunwald/go-jira"
)

// Jira represents a connection to a Jira server.
type Jira struct {
	base64Credentials          string
	settings                   *config.Settings
	cukeReport                 *report.CucumberReport
	hostPath                   string
	jiraPriorities             []jira.Priority
	jiraLabels                 []string
	jiraAuth                   *jira.BasicAuthTransport
	jiraConn                   *jira.Client
	jiraProjs                  *jira.ProjectList
	jiraIssueType              *jira.IssueType
	jiraProject                *jira.Project
	jiraAffectsVersion         *jira.Version
	summaryTemplate            *template.Template
	descriptionTemplate        *template.Template
	acceptanceCriteriaTemplate *template.Template
}

// IssueContext provides the data context used in the Summary and Description templates for Jira issues.
type IssueContext struct {
	Feature    *report.CukeFeature
	Element    *report.CukeElement
	FailedStep *report.CukeStep
}

// Send connects to the configured Jira server, retrieves the user's credentials
// via CLI, and creates new issues for failed tests in the CucumberReport.
func (j *Jira) Send(settings *config.Settings, cukeReport *report.CucumberReport) error {
	j.settings = settings
	j.cukeReport = cukeReport
	j.hostPath = settings.Settings.GetString("integrations.jira.host")

	// load labels
	if settings.Settings.IsSet("integrations.jira.labels") {
		j.jiraLabels = strings.Split(settings.Settings.GetString("integrations.jira.labels"), ",")
	} else {
		j.jiraLabels = []string{}
	}

	if err := j.retrieveCredentials(); err != nil {
		return fmt.Errorf("error encountered while retrieving Jira credentials: %s", err)
	}

	if err := j.connect(); err != nil {
		return fmt.Errorf("error encountered while verifying connection to Jira server: %s", err)
	}

	if err := j.chooseProject(); err != nil {
		return fmt.Errorf("error encountered while listing Jira projects: %s", err)
	}

	if err := j.chooseIssueType(); err != nil {
		return fmt.Errorf("error encountered while listing Jira issue types: %s", err)
	}

	if err := j.getAffectsVersion(); err != nil {
		return fmt.Errorf(`error encountered while getting "Affects Version": %s`, err)
	}

	if err := j.loadTemplates(); err != nil {
		return fmt.Errorf("error encountered while loading Jira templates: %s", err)
	}

	if err := j.createIssues(); err != nil {
		return fmt.Errorf("error encountered while creating issues for failed scenarios: %s", err)
	}

	return nil
}

func (j *Jira) loadTemplates() error {
	var err error

	if j.summaryTemplate, err = GetSummaryTemplate(j.settings); err != nil {
		return fmt.Errorf("error encountered loading Jira summary template: %s", err)
	}

	if j.descriptionTemplate, err = GetDescriptionTemplate(j.settings); err != nil {
		return fmt.Errorf("error encountered loading Jira description template: %s", err)
	}

	return nil
}

func (j *Jira) retrieveCredentials() error {
	var (
		username string
		password string
		bytePwd  []byte
		err      error
	)

	fmt.Println("Enter Jira credentials...")

	// Get Username
	fmt.Print("Username: ")
	consoleScanner := bufio.NewScanner(os.Stdin)
	consoleScanner.Scan()
	username = consoleScanner.Text()

	// Get Password
	fmt.Print("Password: ")
	if bytePwd, err = terminal.ReadPassword(int(syscall.Stdin)); err != nil {
		return fmt.Errorf("error encountered retrieving Jira password: %s", err)
	}
	password = string(bytePwd)

	j.jiraAuth = &jira.BasicAuthTransport{
		Username: username,
		Password: password,
	}

	fmt.Print("\n")

	return nil
}

func (j *Jira) connect() error {
	var (
		self *jira.User
		resp *jira.Response
		err  error
	)

	if j.jiraConn, err = jira.NewClient(j.jiraAuth.Client(), j.hostPath); err != nil {
		return fmt.Errorf("error encountered while connecting to Jira: %s", err)
	}

	// Query for self record to confirm connection
	if self, resp, err = j.jiraConn.User.GetSelf(); err != nil {
		if resp.StatusCode == 401 || resp.StatusCode == 403 {
			return fmt.Errorf("invalid username / password combination")
		}
		return fmt.Errorf("error encountered retrieving self: %s", err)
	}

	fmt.Printf("Successfully connected to Jira as [%s]\n", self.DisplayName)

	// Load available priorities
	if j.jiraPriorities, _, err = j.jiraConn.Priority.GetList(); err != nil {
		return fmt.Errorf("error encountered retrieving available Jira priorities: %s", err)
	}

	return nil
}

func (j *Jira) chooseProject() error {
	var (
		projectIdx = -1
		err        error
	)

	if j.jiraProjs, _, err = j.jiraConn.Project.GetList(); err != nil {
		return fmt.Errorf("error encountered while listing available Jira projects: %s", err)
	}

	fmt.Println("\n                JIRA PROJECTS")
	fmt.Println("_____________________________________________")
	for i, nextProj := range *j.jiraProjs {
		fmt.Printf("[%d]\t%s(%s)\n", i+1, nextProj.Name, nextProj.Key)
	}
	fmt.Println("_____________________________________________")

	// Get project selection
	consoleScanner := bufio.NewScanner(os.Stdin)

	for projectIdx < 1 || projectIdx > len(*j.jiraProjs) {
		fmt.Print("\nSelect Project: ")
		consoleScanner.Scan()

		if projectIdx, err = strconv.Atoi(consoleScanner.Text()); err != nil || projectIdx < 1 || projectIdx > len(*j.jiraProjs) {
			fmt.Println("Invalid project selection, please enter a number from the list above")
		}
	}
	projectIdx--

	// Store selection
	if j.jiraProject, _, err = j.jiraConn.Project.Get(((*j.jiraProjs)[projectIdx]).ID); err != nil {
		return fmt.Errorf("error encountered while retrieving full representation of project [%s]: %s", ((*j.jiraProjs)[projectIdx]).Key, err)
	}
	fmt.Print("\n")

	return nil
}

func (j *Jira) chooseIssueType() error {
	var (
		issueTypeIdx = -1
		err          error
	)

	fmt.Println("\n             JIRA ISSUE TYPES")
	fmt.Println("_____________________________________________")
	for i, nextType := range j.jiraProject.IssueTypes {
		fmt.Printf("[%d]\t%s\n", i+1, nextType.Name)
	}
	fmt.Println("_____________________________________________")

	// Get project selection
	consoleScanner := bufio.NewScanner(os.Stdin)

	for issueTypeIdx < 1 || issueTypeIdx > len(j.jiraProject.IssueTypes) {
		fmt.Printf("\nSelect Issue Type: ")
		consoleScanner.Scan()

		if issueTypeIdx, err = strconv.Atoi(consoleScanner.Text()); err != nil || issueTypeIdx < 1 || issueTypeIdx > len(j.jiraProject.IssueTypes) {
			fmt.Println("Invalid issue type selection, please enter a number from the list above")
		}
	}
	issueTypeIdx--

	// Store selection
	j.jiraIssueType = &j.jiraProject.IssueTypes[issueTypeIdx]

	fmt.Print("\n")

	return nil
}

func (j *Jira) getAffectsVersion() error {
	// Get project selection
	consoleScanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter \"Affects Version\": ")
	consoleScanner.Scan()
	affectsVersion := consoleScanner.Text()

	fmt.Print("\n")

	// Try to get "Affects Version"
	if len(affectsVersion) > 0 {
		for _, v := range j.jiraProject.Versions {
			if strings.EqualFold(affectsVersion, v.Name) {
				j.jiraAffectsVersion = &v
				break
			}
		}

		if j.jiraAffectsVersion == nil {
			return fmt.Errorf("Affects Version [%s] entered, but not found in list of available project versions", affectsVersion)
		}
	}

	return nil
}

func (j *Jira) createIssues() error {
	var (
		issuesCreated int
		issuesSkipped int
	)

	godogDialect := gherkin.GherkinDialectsBuildin().GetDialect(j.settings.Settings.GetString("cucumberDialect"))

	scenarioKeywords := godogDialect.ScenarioKeywords()
	sort.Strings(scenarioKeywords)

	scenarioOutlineKeywords := godogDialect.ScenarioOutlineKeywords()
	sort.Strings(scenarioOutlineKeywords)

	if j.cukeReport.StepsFailed < 1 {
		fmt.Println("No failed steps to send to Jira...")
	} else {
		for _, feature := range j.cukeReport.Features {
			for _, element := range feature.Elements {
				if common.StringSliceContainsFold(scenarioKeywords, element.Type) || common.StringSliceContains(scenarioOutlineKeywords, element.Type) {
					if element.StepsFailed > 0 {
						if created, err := j.createIssue(&feature, &element); err != nil {
							fmt.Printf("error encountered creating issue: %s\n", err)
						} else if created {
							issuesCreated++
						} else {
							issuesSkipped++
						}
					}
				}
			}
		}
	}

	fmt.Printf("Jira Issues Created: %d; Test Failures Skipped: %d\n", issuesCreated, issuesSkipped)

	return nil
}

func (j *Jira) createIssue(feature *report.CukeFeature, element *report.CukeElement) (bool, error) {
	var (
		doCreate     bool
		createdIssue *jira.Issue
		err          error
	)

	issueContext := &IssueContext{
		Feature:    feature,
		Element:    element,
		FailedStep: getFailedStep(element),
	}

	var summaryBytes bytes.Buffer
	if err = j.summaryTemplate.Execute(&summaryBytes, issueContext); err != nil {
		return false, fmt.Errorf("error encountered applying summary template to scenario: %s", err)
	}
	summary := summaryBytes.String()

	var descriptionBytes bytes.Buffer
	if err = j.descriptionTemplate.Execute(&descriptionBytes, issueContext); err != nil {
		return false, fmt.Errorf("error encountered applying description template to scenario: %s", err)
	}
	description := descriptionBytes.String()

	doCreate = getYesOrNo(fmt.Sprintf("Create [%s] (Y/n): ", summary))
	if doCreate {
		newIssue := &jira.Issue{
			Fields: &jira.IssueFields{
				Summary:     summary,
				Description: description,
				Project: jira.Project{
					ID:  j.jiraProject.ID,
					Key: j.jiraProject.Key,
				},
				Type:     *j.jiraIssueType,
				Labels:   j.jiraLabels,
				Unknowns: tcontainer.NewMarshalMap(),
				Priority: j.choosePriority(),
			},
		}

		// Add affectsVersion if possible
		if j.jiraAffectsVersion != nil {
			newIssue.Fields.Unknowns.Set("versions", []jira.Version{*j.jiraAffectsVersion})
		}

		if createdIssue, _, err = j.jiraConn.Issue.Create(newIssue); err != nil {
			return false, fmt.Errorf("error encountered while creating new issue: %s", err)
		}

		fmt.Printf("\tIssue [%s] created...\n", createdIssue.Key)
	}

	fmt.Print("\n")

	return doCreate, nil
}

func (j *Jira) choosePriority() *jira.Priority {
	var (
		priorityIdx = -1
		err         error
	)

	fmt.Println("\tChoose priority...")

	for i, p := range j.jiraPriorities {
		fmt.Printf("\t\t[%d] %s\n", i+1, p.Name)
	}

	// Get project selection
	consoleScanner := bufio.NewScanner(os.Stdin)

	// Get priority selection
	for priorityIdx < 1 || priorityIdx > len(j.jiraPriorities) {
		fmt.Printf("\tEnter priority selection: ")
		consoleScanner.Scan()

		if priorityIdx, err = strconv.Atoi(consoleScanner.Text()); err != nil || priorityIdx < 1 || priorityIdx > len(j.jiraPriorities) {
			fmt.Println("\t\tInvalid priority selection, please enter a number from the list above")
		}
	}
	priorityIdx--

	return &j.jiraPriorities[priorityIdx]
}

func getYesOrNo(query string) bool {
	var (
		consoleScanner  = bufio.NewScanner(os.Stdin)
		response        string
		trimmedResponse string
	)

	fmt.Print(query)
	consoleScanner.Scan()

	response = strings.ToUpper(consoleScanner.Text())

	if len(response) < 1 {
		fmt.Printf("Invalid response: %s; please enter Y or N\n\n", response)
		return getYesOrNo(query)
	}

	trimmedResponse = response[:1]
	if !strings.EqualFold(response, "Y") && !strings.EqualFold(response, "N") {
		fmt.Printf("Invalid response: %s; please enter Y or N\n\n", response)
		return getYesOrNo(query)
	}

	return strings.EqualFold(trimmedResponse, "Y")
}

func getFailedStep(element *report.CukeElement) *report.CukeStep {
	if element.StepsFailed < 1 {
		return nil
	}

	for _, step := range element.Steps {
		if strings.EqualFold(step.Result.Status, "failed") {
			return &step
		}
	}

	return nil
}
