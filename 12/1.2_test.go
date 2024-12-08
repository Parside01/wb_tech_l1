package main

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateSquaresConcurrency(t *testing.T) {
	var out bytes.Buffer
	expect := "4 16 36 64 100 "

	CalculateSquaresConcurrency(&out)

	assert.Equal(t, expect, out.String())
}
