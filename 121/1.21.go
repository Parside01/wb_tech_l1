package main

import (
	"fmt"
	"io"
)

type Speakable interface {
	Speak(w io.Writer)
}

type Human struct{}

func NewAdapter(human *Human) Speakable {
	return &Adapter{human}
}

func (h *Human) Hello(w io.Writer) {
	if _, err := fmt.Fprint(w, "Hello"); err != nil {
		panic(err)
	}
}

type Adapter struct {
	*Human
}

func (a *Adapter) Speak(w io.Writer) {
	a.Hello(w)
}
