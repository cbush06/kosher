package jira

import (
	"fmt"
	"log"
	"text/template"

	"github.com/cbush06/kosher/config"
	"github.com/spf13/afero"
)

const (
	defaultSummaryTemplate     = `{{.Feature.Name}}: {{.Element.Name}}`
	defaultDescriptionTemplate = `h2. Issue:
{{.FailedStep.Result.Error}}

h2. Feature Title:
{{.Feature.Name}}

h2. Scenario Title:
{{.Element.Name}}

h2. Scenario Description
{{.Element.Description}}

h2. Failed Step / Actual Result:
{color:red}*{{.FailedStep.GetTrimmedKeyword}}* {{.FailedStep.Name}}{color}
{color:red}{{.FailedStep.Result.Error}}

h2. Steps / Expected Result:
{{range .Element.Steps}}# *{{.GetTrimmedKeyword}}* {{.Name}}
{{end}}`
)

// GetSummaryTemplate returns the user-specified or default template for a Jira issue summary.
func GetSummaryTemplate(s *config.Settings) (*template.Template, error) {
	templateText := defaultSummaryTemplate

	if s.Settings.IsSet("integrations.jira.summaryTemplate") {
		summaryTemplateFileName := s.Settings.GetString("integrations.jira.summaryTemplate")
		if retrievedTemplate, err := tryRetrieveTemplateText(s, summaryTemplateFileName); err == nil {
			templateText = retrievedTemplate
		} else {
			log.Printf("error encountered retrieving Jira summary template: %s\n", err)
			log.Println("using default Jira summary template")
		}
	}

	return template.New("summaryTemplate").Parse(templateText)
}

// GetDescriptionTemplate returns the user-specified or default template for a Jira issue description.
func GetDescriptionTemplate(s *config.Settings) (*template.Template, error) {
	templateText := defaultDescriptionTemplate

	if s.Settings.IsSet("integrations.jira.descriptionTemplate") {
		descriptionTemplateName := s.Settings.GetString("integrations.jira.descriptionTemplate")
		if retrievedTemplate, err := tryRetrieveTemplateText(s, descriptionTemplateName); err == nil {
			templateText = retrievedTemplate
		} else {
			log.Printf("error encountered retrieving Jira description template: %s\n", err)
			log.Println("using default Jira description template")
		}
	}

	return template.New("descriptionTemplate").Parse(templateText)
}

func tryRetrieveTemplateText(s *config.Settings, fileName string) (string, error) {
	filePath, _ := s.FileSystem.ConfigDir.RealPath(fileName)

	if fileExists, err := afero.Exists(s.FileSystem.ConfigDir, fileName); err != nil {
		return "", fmt.Errorf("error encountered while retrieving file [%s]: %s", filePath, err)
	} else if !fileExists {
		return "", fmt.Errorf("file [%s] does not exist", filePath)
	} else {
		var (
			fileBytes []byte
			err       error
		)

		if fileBytes, err = afero.ReadFile(s.FileSystem.ConfigDir, fileName); err != nil {
			return "", fmt.Errorf("error encountered reading file [%s]: %s", filePath, err)
		}

		return string(fileBytes), nil
	}
}
