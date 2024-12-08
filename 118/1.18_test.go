package main

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestCounterInc(t *testing.T) {
	counter := NewCounter()
	wg := sync.WaitGroup{}
	expect := 0
	operationsCount := 100

	for i := 0; i < operationsCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Inc()
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Dec()
		}()
	}
	wg.Wait()
	assert.Equal(t, expect, counter.Value())
}
