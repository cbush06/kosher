package common

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParseDate(t *testing.T) {
	tests := []struct {
		date     string
		format   string
		expected time.Time
	}{
		{
			"February 06, 2019",
			"MMMM DD, YYYY",
			time.Date(2019, time.February, 06, 0, 0, 0, 0, time.Local),
		},
		{
			"Feb 06, 2019",
			"MMM DD, YYYY",
			time.Date(2019, time.February, 06, 0, 0, 0, 0, time.Local),
		},
		{
			"02-06-2019",
			"MM-DD-YYYY",
			time.Date(2019, time.February, 06, 0, 0, 0, 0, time.Local),
		},
		{
			"02-06-19",
			"MM-DD-YY",
			time.Date(2019, time.February, 06, 0, 0, 0, 0, time.Local),
		},
		{
			"Wednesday, Feb 06, 2019",
			"DDDD, MMM DD, YYYY",
			time.Date(2019, time.February, 06, 0, 0, 0, 0, time.Local),
		},
		{
			"Wed, Feb 06, 2019",
			"DDD, MMM DD, YYYY",
			time.Date(2019, time.February, 06, 0, 0, 0, 0, time.Local),
		},
		{
			"02-06-2019",
			"",
			time.Time{},
		},
		{
			"02-06-2019",
			"MM-qq-YYYY",
			time.Time{},
		},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("Parse-[%s]-with-pattern-[%s]", tc.date, tc.format), func(t *testing.T) {
			actual := ParseDate(tc.date, tc.format)
			assert.Equalf(t, tc.expected, actual, fmt.Sprintf("expected [%s] but got [%s]", tc.expected.Local(), actual.Local()))
		})
	}
}

func TestFormatDate(t *testing.T) {
	tests := []struct {
		date     time.Time
		format   string
		expected string
	}{
		{
			time.Date(2019, time.February, 06, 0, 0, 0, 0, time.Local),
			"MMMM DD, YYYY",
			"February 06, 2019",
		},
		{
			time.Date(2019, time.February, 06, 0, 0, 0, 0, time.Local),
			"MMM DD, YYYY",
			"Feb 06, 2019",
		},
		{
			time.Date(2019, time.February, 06, 0, 0, 0, 0, time.Local),
			"MM-DD-YYYY",
			"02-06-2019",
		},
		{
			time.Date(2019, time.February, 06, 0, 0, 0, 0, time.Local),
			"MM-DD-YY",
			"02-06-19",
		},
		{
			time.Date(2019, time.February, 06, 0, 0, 0, 0, time.Local),
			"DDDD, MMM DD, YYYY",
			"Wednesday, Feb 06, 2019",
		},
		{
			time.Date(2019, time.February, 06, 0, 0, 0, 0, time.Local),
			"DDD, MMM DD, YYYY",
			"Wed, Feb 06, 2019",
		},
		{
			time.Date(2019, time.February, 06, 0, 0, 0, 0, time.Local),
			"",
			"",
		},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("Format-[%s]-with-pattern-[%s]", tc.date.Local(), tc.format), func(t *testing.T) {
			actual := FormatDate(tc.date, tc.format)
			assert.Equalf(t, tc.expected, actual, fmt.Sprintf("expected [%s] but got [%s]", tc.expected, actual))
		})
	}
}
