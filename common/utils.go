package common

import (
	"sort"
	"strings"
)

// StringSliceContains returns true if the SORTED slice `s` contains `v`.
func StringSliceContains(s []string, v string) bool {
	idx := sort.SearchStrings(s, v)
	return idx < len(s) && s[idx] == v
}

// StringSliceContainsFold returns true if the SORTED slice `s` contains `v`, ignoring case.
func StringSliceContainsFold(s []string, v string) bool {
	vUpper := strings.ToUpper(v)
	sUpper := s[:]
	for i, next := range sUpper {
		sUpper[i] = strings.ToUpper(next)
	}
	idx := sort.SearchStrings(sUpper, vUpper)
	return idx < len(sUpper) && sUpper[idx] == vUpper
}
