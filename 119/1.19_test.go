package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverse(t *testing.T) {
	tt := []struct {
		Name   string
		Input  string
		Expect string
	}{
		{
			Name:   "Empty",
			Input:  "",
			Expect: "",
		},
		{
			Name:   "Russian",
			Input:  "главрыба",
			Expect: "абырвалг",
		},
		{
			Name:   "English",
			Input:  "hello",
			Expect: "olleh",
		},
		{
			Name:   "Chinese",
			Input:  "你好",
			Expect: "好你",
		},
		{
			Name:   "Japanese",
			Input:  "こんにちは",
			Expect: "はちにんこ",
		},
		{
			Name:   "Arabic",
			Input:  "مرحبا",
			Expect: "ابحرم",
		},
		{
			Name:   "Emoji",
			Input:  "😊🌍",
			Expect: "🌍😊",
		},
	}
	for _, tc := range tt {
		t.Run(tc.Name, func(t *testing.T) {
			got := Reverse(tc.Input)
			assert.Equal(t, tc.Expect, got)
		})
	}
}
