package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSetsIntersection(t *testing.T) {
	tt := []struct {
		Name   string
		Left   []int
		Right  []int
		Expect []int
	}{
		{
			Name:   "Empty sets",
			Left:   []int{},
			Right:  []int{},
			Expect: []int{},
		},
		{
			Name:   "Default sets",
			Left:   []int{1, 2, 3, 4, 5},
			Right:  []int{5, 4, 3, 2, 1},
			Expect: []int{1, 2, 3, 4, 5},
		},
	}

	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			got := SetsIntersection(tc.Left, tc.Right)
			assert.Equal(t, tc.Expect, got)
		})
	}
}
