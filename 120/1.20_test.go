package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseWord(t *testing.T) {
	tt := []struct {
		Name   string
		Input  string
		Expect string
	}{
		{
			Name:   "Empty",
			Input:  "",
			Expect: "",
		},
		{
			Name:   "Normally string",
			Input:  "snow dog sun",
			Expect: "sun dog snow",
		},
	}

	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			got := ReverseWords(tc.Input)
			assert.Equal(t, tc.Expect, got)
		})
	}
}
