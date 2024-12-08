package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTypeIdentification(t *testing.T) {
	tt := []struct {
		Name       string
		Value      any
		ExpectType string
	}{
		{
			Name:       "int",
			Value:      10,
			ExpectType: "int",
		},
		{
			Name:       "string",
			Value:      "Hello, world!",
			ExpectType: "string",
		},
		{
			Name:       "bool",
			Value:      true,
			ExpectType: "bool",
		},
		{
			Name:       "chan int",
			Value:      make(chan int),
			ExpectType: "chan int",
		},
		{
			Name:       "chan string",
			Value:      make(chan string),
			ExpectType: "chan string",
		},
		{
			Name:       "chan bool",
			Value:      make(chan bool),
			ExpectType: "chan bool",
		},
	}
	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			got := TypeIdentification(tc.Value)
			assert.Equal(t, tc.ExpectType, got)
		})
	}
}
