package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	tt := []struct {
		Name     string
		Arr      []int
		Target   int
		Expected int
	}{
		{
			Name:     "Empty",
			Arr:      []int{},
			Target:   0,
			Expected: -1,
		},
		{
			Name:     "Default",
			Arr:      []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			Target:   5,
			Expected: 5,
		},
		{
			Name:     "No element in array",
			Arr:      []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			Target:   16,
			Expected: -1,
		},
	}
	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			got := BinarySearch(tc.Arr, tc.Target)
			assert.Equal(t, tc.Expected, got)
		})
	}
}
