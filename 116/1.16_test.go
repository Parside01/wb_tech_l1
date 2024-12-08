package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQuickSort(t *testing.T) {
	tt := []struct {
		Name   string
		Arr    []int
		Expect []int
	}{
		{
			Name:   "Empty",
			Arr:    []int{},
			Expect: []int{},
		},
		{
			Name:   "One element",
			Arr:    []int{1},
			Expect: []int{1},
		},
		{
			Name:   "Non sorted",
			Arr:    []int{5, 3, 2, 8, 10, 11},
			Expect: []int{2, 3, 5, 8, 10, 11},
		},
		{
			Name:   "Reverse sort",
			Arr:    []int{5, 4, 3, 2, 1, 0, -1},
			Expect: []int{-1, 0, 1, 2, 3, 4, 5},
		},
	}
	for _, tc := range tt {
		t.Run(tc.Name, func(tt *testing.T) {
			QuickSort(tc.Arr)
			assert.Equal(t, tc.Expect, tc.Arr)
		})
	}
}
