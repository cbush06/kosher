package jira

import (
	"bufio"
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"

	"github.com/cbush06/kosher/config"
	"github.com/cbush06/kosher/report"
)

// Jira represents a connection to a Jira server.
type Jira struct {
	base64Credentials string
	settings          *config.Settings
	cukeReport        *report.CucumberReport
	username          string
	password          string
	hostPath          string
}

// Send connects to the configured Jira server, retrieves the user's credentials
// via CLI, and creates new issues for failed tests in the CucumberReport.
func (j *Jira) Send(settings *config.Settings, cukeReport *report.CucumberReport) error {
	j.settings = settings
	j.cukeReport = cukeReport
	j.hostPath = settings.Settings.GetString("integrations.jira.host")

	if err := j.retrieveCredentials(); err != nil {
		return fmt.Errorf("error encountered while retrieving Jira credentials: %s", err)
	}

	if err := j.verifyConnection(); err != nil {
		return fmt.Errorf("error encountered while verifying connection to Jira server: %s", err)
	}

	return nil
}

func (j *Jira) retrieveCredentials() error {
	var (
		bytePwd []byte
		err     error
	)

	fmt.Println("Enter Jira credentials...")

	// Get Username
	fmt.Print("Username: ")
	consoleScanner := bufio.NewScanner(os.Stdin)
	consoleScanner.Scan()
	j.username = consoleScanner.Text()

	// Get Password
	fmt.Print("Password: ")
	if bytePwd, err = terminal.ReadPassword(syscall.Stdin); err != nil {
		return fmt.Errorf("error encountered retrieving Jira password: %s", err)
	}
	j.password = string(bytePwd)

	fmt.Print("\n")

	return nil
}

func (j *Jira) verifyConnection() error {
	var (
		requestURL *url.URL
		response   *http.Response
		err        error
	)

	requestURL, _ = url.Parse(j.hostPath)
	requestURL.Path = path.Join(requestURL.Path, "/rest/api/2/project")

	request, _ := http.NewRequest("GET", requestURL.String(), bytes.NewBufferString(""))
	request.SetBasicAuth(j.username, j.password)
	request.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	if response, err = client.Do(request); err != nil {
		return fmt.Errorf("error encountered while verifying connection to Jira server: %s", err)
	}

	if response.StatusCode != 200 {
		return fmt.Errorf("failed to verify connection; status code [%d] returned", response.StatusCode)
	}

	return nil
}
