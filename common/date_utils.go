package common

import (
	"strings"
	"time"
)

// ParseDate parses a given string to a `time.Time` using the `dateFormat` value specified in the settings file
func ParseDate(date string, dateFormat string) time.Time {
	var (
		err    error
		goDate time.Time
	)

	if len(dateFormat) < 1 {
		return time.Time{}
	}

	if goDate, err = time.ParseInLocation(convertDateFormatToGoFormat(dateFormat), date, time.Local); err != nil {
		return time.Time{}
	}

	return goDate
}

// FormatDate converts a `time.Time` to a string using the `dateFormat` value specified in the settings file
func FormatDate(t time.Time, dateFormat string) string {
	if len(dateFormat) < 1 {
		return ""
	}
	return t.Format(convertDateFormatToGoFormat(dateFormat))
}

func convertDateFormatToGoFormat(dateFormat string) string {
	dateFormat = strings.Replace(dateFormat, "MMMM", "January", 1)
	dateFormat = strings.Replace(dateFormat, "MMM", "Jan", 1)
	dateFormat = strings.Replace(dateFormat, "MM", "01", 1)
	dateFormat = strings.Replace(dateFormat, "YYYY", "2006", 1)
	dateFormat = strings.Replace(dateFormat, "YY", "06", 1)
	dateFormat = strings.Replace(dateFormat, "DDDD", "Monday", 1)
	dateFormat = strings.Replace(dateFormat, "DDD", "Mon", 1)
	dateFormat = strings.Replace(dateFormat, "DD", "02", 1)
	return dateFormat
}
