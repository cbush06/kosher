package common

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVersion(t *testing.T) {
	tests := []struct {
		version  *Version
		expected string
	}{
		{
			&Version{1, 2, 3, ""},
			"1.2.3",
		},
		{
			&Version{1, 2, 3, "DEV"},
			"1.2.3-DEV",
		},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("Get-Version-for-[%v]", tc.version), func(t *testing.T) {
			actual := tc.version.Version()
			assert.Equalf(t, tc.expected, actual, "expected [%s] but got [%s]", tc.expected, actual)
		})
	}
}
