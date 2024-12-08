package main

import (
	"bytes"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGoroutineStopUseChannel(t *testing.T) {
	var out bytes.Buffer
	expect := "Stopped"
	GoroutineStopUseChannel(&out, expect)
	assert.Equal(t, expect, out.String())
}

func TestGoroutineStartUseCloseChannel(t *testing.T) {
	var out bytes.Buffer

	messagesCount := gofakeit.IntRange(1, 100)
	messages := make([]string, messagesCount)
	stopMessage := "Stopped"
	expect := ""
	for i := range messages {
		messages[i] = gofakeit.Sentence(10)
		expect += messages[i]
	}
	expect += stopMessage

	GoroutineStopUseCloseChannel(&out, messages, stopMessage)
	assert.Equal(t, expect, out.String())
}

func TestGoroutineStopUseVariables(t *testing.T) {
	var out bytes.Buffer

	stopMessage := "Stopped"
	GoroutineStopUseVariables(&out, stopMessage)

	assert.Equal(t, stopMessage, out.String())
}

func TestGoroutineStopUseContext(t *testing.T) {
	var out bytes.Buffer

	stopMessage := "Stopped"
	GoroutineStopUseContext(&out, stopMessage)
	assert.Equal(t, stopMessage, out.String())
}

func TestGoroutineStopUseGoexit(t *testing.T) {
	var out bytes.Buffer
	stopMessage := "Stopped"
	GoroutineStopUseGoexit(&out, stopMessage)
	assert.Equal(t, stopMessage, out.String())
}

func TestGoroutineStartUseCond(t *testing.T) {
	var out bytes.Buffer
	stopMessage := "Stopped"
	GoroutineStopUseCond(&out, stopMessage)
	assert.Equal(t, stopMessage, out.String())
}
