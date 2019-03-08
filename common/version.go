package common

import (
	"fmt"
)

// Version represents a fully qualified version number
type Version struct {
	Major  int
	Minor  int
	Patch  int
	Suffix string
}

// Version returns the version information in string form (e.g. 1.02.55)
func (v *Version) Version() string {
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
}
