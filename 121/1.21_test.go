package main

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdapter(t *testing.T) {
	adapter := NewAdapter(&Human{})
	var out bytes.Buffer
	adapter.Speak(&out)

	assert.Equal(t, out.String(), "Hello")
}
