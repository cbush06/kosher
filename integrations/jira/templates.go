package jira

import (
	"log"
	"text/template"

	"github.com/cbush06/kosher/config"
	"github.com/spf13/afero"
)

const (
	defaultSummaryTemplate     = `{{.Feature.Name}}: {{.Element.Name}}`
	defaultDescriptionTempalte = ``
)

// GetSummaryTemplate returns the default template for a Jira issue summary.
func GetSummaryTemplate(s *config.Settings) (*template.Template, error) {
	templateText := defaultSummaryTemplate

	if s.Settings.IsSet("jira.summaryTemplate") {
		summaryTemplateFileName := s.Settings.GetString("jira.summaryTemplate")
		summaryTemplateFilePath, _ := s.FileSystem.ConfigDir.RealPath(summaryTemplateFileName)

		if fileExists, err := afero.Exists(s.FileSystem.ConfigDir, summaryTemplateFileName); err != nil {
			log.Printf("error encountered while retrieving file [%s]; using default Jira summary template\n", summaryTemplateFilePath)
		} else if !fileExists {
			log.Printf("file [%s] does not exist\n", summaryTemplateFilePath)
		} else {
			var (
				summaryTemplateFileBytes []byte
				err                      error
			)

			if summaryTemplateFileBytes, err = afero.ReadFile(s.FileSystem.ConfigDir, summaryTemplateFileName); err != nil {
				log.Printf("error encountered reading file [%s]; using default Jira summary template\n", summaryTemplateFilePath)
			} else {
				templateText = string(summaryTemplateFileBytes)
			}
		}
	}

	return template.New("summaryTemplate").Parse(templateText)
}
