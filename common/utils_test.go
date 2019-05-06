package common

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringSliceContains(t *testing.T) {
	tests := []struct {
		slice    []string
		query    string
		expected bool
	}{
		{
			[]string{
				"ABC", "DEF", "GHI",
			},
			"ABC",
			true,
		},
		{
			[]string{
				"ABC", "DEF", "GHI",
			},
			"JKL",
			false,
		},
		{
			[]string{
				"ABC", "DEF", "GHI",
			},
			"abc",
			false,
		},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("Does-[%v]-Contain-[%s]", tc.slice, tc.query), func(t *testing.T) {
			assert.Equal(t, tc.expected, StringSliceContains(tc.slice, tc.query))
		})
	}
}

func TestStringSliceContainsFold(t *testing.T) {
	tests := []struct {
		slice    []string
		query    string
		expected bool
	}{
		{
			[]string{
				"ABC", "DEF", "GHI",
			},
			"ABC",
			true,
		},
		{
			[]string{
				"ABC", "DEF", "GHI",
			},
			"abc",
			true,
		},
		{
			[]string{
				"ABC", "DEF", "GHI",
			},
			"jkl",
			false,
		},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("Does-[%v]-Contain-[%s]", tc.slice, tc.query), func(t *testing.T) {
			assert.Equal(t, tc.expected, StringSliceContainsFold(tc.slice, tc.query))
		})
	}
}
