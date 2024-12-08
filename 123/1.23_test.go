package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeleteElement(t *testing.T) {
	tt := []struct {
		Name     string
		Slice    []int
		Index    int
		Expected []int
	}{
		{
			Name:     "Empty",
			Slice:    []int{},
			Index:    0,
			Expected: []int{},
		},
		{
			Name:     "Great index",
			Slice:    []int{1, 2, 3},
			Index:    4,
			Expected: []int{1, 2, 3},
		},
		{
			Name:     "Less index",
			Slice:    []int{1, 2, 3},
			Index:    -1,
			Expected: []int{1, 2, 3},
		},
		{
			Name:     "Normally test",
			Slice:    []int{1, 2, 3},
			Index:    1,
			Expected: []int{1, 3},
		},
	}
	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			got := DeleteElement(tc.Slice, tc.Index)
			assert.Equal(t, tc.Expected, got)
		})
	}
}
