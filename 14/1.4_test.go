package main

import (
	"bytes"
	"context"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWorker(t *testing.T) {
	var buf bytes.Buffer
	msgs := make([]string, 15)
	in := make(chan string)

	for i := range msgs {
		msgs[i] = gofakeit.Sentence(5)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go Worker(ctx, &buf, in)
	expect := ""
	for _, msg := range msgs {
		expect += msg
		in <- msg
	}

	cancel()
	assert.Equal(t, expect, buf.String())
}
