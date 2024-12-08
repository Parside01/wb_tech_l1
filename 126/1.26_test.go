package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsUniqChars(t *testing.T) {
	tt := []struct {
		Name   string
		Input  string
		Expect bool
	}{
		{
			Name:   "Empty",
			Input:  "",
			Expect: true,
		},
		{
			Name:   "Lower string",
			Input:  "abcd",
			Expect: true,
		},
		{
			Name:   "Upper string",
			Input:  "ABCD",
			Expect: true,
		},
		{
			Name:   "Different string",
			Input:  "abCdefAaf",
			Expect: false,
		},
		{
			Name:   "Not unique chars",
			Input:  "aabcd",
			Expect: false,
		},
	}
	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			got := IsUniqChars(tc.Input)
			assert.Equal(t, tc.Expect, got)
		})
	}
}
