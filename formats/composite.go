package formats

import (
	"io"

	"github.com/cbush06/godog"
	"github.com/cbush06/godog/gherkin"
)

// NewCompositeFormatterFunc builds the formatter constructor function used by godog when registering this formatter
func NewCompositeFormatterFunc(formatters []string) godog.FormatterFunc {
	return func(suite string, out io.Writer) godog.Formatter {
		var newFormatters = make([]godog.Formatter, 0)
		for _, nextFormat := range formatters {
			formatter := godog.FindFmt(nextFormat)(suite, out)
			newFormatters = append(newFormatters, formatter)
		}

		return &CompositeFormat{
			formats: newFormatters,
		}
	}
}

// CompositeFormat allows using more than one formatter at a time
type CompositeFormat struct {
	formats []godog.Formatter
}

// Feature calls the Feature functions on each formatter composing this CompositeFormat.
func (c *CompositeFormat) Feature(ft *gherkin.Feature, p string, src []byte) {
	for _, f := range c.formats {
		f.Feature(ft, p, src)
	}
}

// Node calls the Node functions on each formatter composing this CompositeFormat.
func (c *CompositeFormat) Node(n interface{}) {
	for _, f := range c.formats {
		f.Node(n)
	}
}

// Defined calls the Defined functions on each formatter composing this CompositeFormat.
func (c *CompositeFormat) Defined(s *gherkin.Step, sd *godog.StepDef) {
	for _, f := range c.formats {
		f.Defined(s, sd)
	}
}

// Failed calls the Failed functions on each formatter composing this CompositeFormat.
func (c *CompositeFormat) Failed(s *gherkin.Step, sd *godog.StepDef, err error) {
	for _, f := range c.formats {
		f.Failed(s, sd, err)
	}
}

// Passed calls the Passed functions on each formatter composing this CompositeFormat.
func (c *CompositeFormat) Passed(s *gherkin.Step, sd *godog.StepDef) {
	for _, f := range c.formats {
		f.Passed(s, sd)
	}
}

// Skipped calls the Skipped functions on each formatter composing this CompositeFormat.
func (c *CompositeFormat) Skipped(s *gherkin.Step, sd *godog.StepDef) {
	for _, f := range c.formats {
		f.Skipped(s, sd)
	}
}

// Undefined calls the Undefined functions on each formatter composing this CompositeFormat.
func (c *CompositeFormat) Undefined(s *gherkin.Step, sd *godog.StepDef) {
	for _, f := range c.formats {
		f.Undefined(s, sd)
	}
}

// Pending calls the Pending functions on each formatter composing this CompositeFormat.
func (c *CompositeFormat) Pending(s *gherkin.Step, sd *godog.StepDef) {
	for _, f := range c.formats {
		f.Pending(s, sd)
	}
}

// Summary calls the Summary functions on each formatter composing this CompositeFormat.
func (c *CompositeFormat) Summary() {
	for _, f := range c.formats {
		f.Summary()
	}
}
