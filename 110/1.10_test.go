package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDivideTemperature(t *testing.T) {
	tt := []struct {
		Name     string
		Arr      []float32
		Expected map[int][]float32
	}{
		{
			Name:     "Empty",
			Arr:      []float32{},
			Expected: map[int][]float32{},
		},
		{
			Name: "One element",
			Arr:  []float32{-30.33},
			Expected: map[int][]float32{
				-30: {-30.33},
			},
		},
		{
			Name: "More elements",
			Arr:  []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5},
			Expected: map[int][]float32{
				-20: {-25.4, -27.0, -21.0},
				10:  {13, 19.0, 15.5},
				20:  {24.5},
				30:  {32.5},
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			got := DivideTemperature(tc.Arr)
			assert.Equal(t, tc.Expected, got)
		})
	}
}
