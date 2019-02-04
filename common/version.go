package common

import (
	"fmt"
	"strings"
)

// Version represents a fully qualified version number
type Version struct {
	Major  int
	Minor  int
	Patch  int
	Suffix string
}

// Version returns the version information in string form (e.g. 1.02.55-DEV)
func (v *Version) Version() string {
	if strings.EqualFold(v.Suffix, "prod") {
		return fmt.Sprintf("%d.%02d.%02d", v.Major, v.Minor, v.Patch)
	}
	return fmt.Sprintf("%d.%02d.%02d-%s", v.Major, v.Minor, v.Patch, v.Suffix)
}
