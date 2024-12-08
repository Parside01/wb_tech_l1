package main

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestDistanceTo(t *testing.T) {
	tt := []struct {
		Name   string
		Point  *Point
		Other  *Point
		Expect float64
	}{
		{
			Name:   "Negative coords",
			Point:  NewPoint(-1, -1),
			Other:  NewPoint(-4, -5),
			Expect: 5.0,
		},
		{
			Name:   "Mixed coords",
			Point:  NewPoint(1, -1),
			Other:  NewPoint(-1, 1),
			Expect: math.Sqrt(8),
		},
		{
			Name:   "Large coords",
			Point:  NewPoint(1000, 1000),
			Other:  NewPoint(2000, 2000),
			Expect: math.Sqrt(2000000),
		},
		{
			Name:   "Same coords",
			Point:  NewPoint(-1, -1),
			Other:  NewPoint(-1, -1),
			Expect: math.Sqrt(0),
		},
	}

	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			got := tc.Point.DistanceTo(tc.Other)
			assert.Equal(t, tc.Expect, got)
		})
	}
}
