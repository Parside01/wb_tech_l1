package main

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsertToSet(t *testing.T) {
	tt := []struct {
		Name      string
		Args      []string
		ExpectOut string
	}{
		{
			Name:      "Empty set",
			Args:      []string{},
			ExpectOut: "{  }",
		},
		{
			Name:      "Insert to empty set",
			Args:      []string{"dog"},
			ExpectOut: "{ dog }",
		},
		{
			Name:      "Double insert same value",
			Args:      []string{"dog", "dog"},
			ExpectOut: "{ dog }",
		},
		{
			Name:      "Insert more unique values",
			Args:      []string{"dog", "cat", "bird"},
			ExpectOut: "{ dog, cat, bird }",
		},
		{
			Name:      "Double insert more unique values",
			Args:      []string{"cat", "cat", "bird", "bird", "dog", "dog"},
			ExpectOut: "{ cat, bird, dog }",
		},
	}

	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			set := NewSet()
			for _, arg := range tc.Args {
				set.Insert(arg)
			}

			var out bytes.Buffer
			set.Println(&out)

			assert.Equal(t, tc.ExpectOut, out.String())
		})
	}
}
