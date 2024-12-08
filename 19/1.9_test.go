package main

import (
	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMoveArrayToChannel(t *testing.T) {
	expect := make([]int, 1000)
	gofakeit.Slice(&expect)

	in := make(chan int)
	var got []int

	go MoveArrayToChannel(expect, in)
	for i := range in {
		got = append(got, i)
	}

	assert.Equal(t, expect, got)
}

func TestProcessMessagesToChannel(t *testing.T) {
	arr := make([]int, 1000)
	gofakeit.Slice(&arr)

	in := make(chan int)
	out := make(chan int)

	var got []int

	go MoveArrayToChannel(arr, in)
	go ProcessMessagesToChannel(in, out)

	for i := range out {
		got = append(got, i)
	}
	for i := range arr {
		assert.Equal(t, arr[i]*2, got[i])
	}
}
