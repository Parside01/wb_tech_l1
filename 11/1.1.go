package main

import (
	"fmt"
	"io"
)

type Human struct {
	Name string
}

func (h *Human) SayHello(w io.Writer) {
	_, err := fmt.Fprintf(w, "Hello, my name is %s", h.Name)
	if err != nil {
		panic(err)
	}
}

type Action struct {
	Human
}
