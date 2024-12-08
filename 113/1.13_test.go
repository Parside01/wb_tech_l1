package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSwap(t *testing.T) {
	tt := []struct {
		Name   string
		First  int
		Second int
	}{
		{
			Name:   "Equal values",
			First:  1,
			Second: 1,
		},
		{
			Name:   "Different values",
			First:  10,
			Second: 9,
		},
	}

	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			a, b := tc.First, tc.Second
			Swap(&tc.First, &tc.Second)
			assert.Equal(t, a, tc.Second)
			assert.Equal(t, b, tc.First)
		})
	}
}
