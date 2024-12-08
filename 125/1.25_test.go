package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestSleep(t *testing.T) {
	duration := time.Millisecond * 500
	start := time.Now()
	time.Sleep(duration)
	sleepTime := time.Since(start)

	assert.GreaterOrEqual(t, sleepTime, duration)
}
