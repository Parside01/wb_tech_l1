package main

import (
	"bytes"
	"fmt"
	"github.com/brianvoe/gofakeit/v7"
	"testing"
)

func TestAction(t *testing.T) {
	casesCount := 10
	tt := make([]struct {
		name   string
		expect string
	}, casesCount)

	for i := range casesCount {
		name := gofakeit.Name()
		tt[i].name = name
		tt[i].expect = fmt.Sprintf("Hello, my name is %s", name)
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var out bytes.Buffer
			action := Action{
				Human: Human{
					Name: tc.name,
				},
			}

			action.SayHello(&out)

			if out.String() != tc.expect {
				t.Errorf("got %s, want %s", out.String(), tc.expect)
			}
		})
	}
}
