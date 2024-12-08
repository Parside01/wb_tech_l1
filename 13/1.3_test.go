package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSumSquaresConcurrency(t *testing.T) {
	expect := 4 + 16 + 36 + 64 + 100
	assert.Equal(t, expect, SumSquaresConcurrency())
}
