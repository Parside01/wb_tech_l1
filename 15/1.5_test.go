package main

import (
	"bytes"
	"context"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestStartMessageConsumer(t *testing.T) {
	in := make(chan string)
	var buf bytes.Buffer
	messages := make([]string, 100)
	for i := range messages {
		messages[i] = gofakeit.Sentence(5)
	}
	go StartMessageConsumer(context.Background(), in, &buf)

	for _, msg := range messages {
		in <- msg
	}
	close(in)
	expect := strings.Join(messages, "\n") + "\n"
	assert.Equal(t, expect, buf.String())
}
