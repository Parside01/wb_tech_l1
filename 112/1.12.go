package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type Set struct {
	set map[string]struct{}
}

func NewSet() *Set {
	return &Set{
		set: make(map[string]struct{}),
	}
}

func (s *Set) Insert(value string) {
	if _, found := s.set[value]; !found {
		s.set[value] = struct{}{}
	}
}

func (s *Set) Remove(value string) {
	delete(s.set, value)
}

func (s *Set) Contains(value string) bool {
	_, found := s.set[value]
	return found
}

func (s *Set) Println(w io.Writer) {
	keys := make([]string, 0, len(s.set))
	for k := range s.set {
		keys = append(keys, k)
	}

	if _, err := fmt.Fprintf(w, "{ %s }", strings.Join(keys, ", ")); err != nil {
		panic(err)
	}
}

func main() {
	strs := []string{"cat", "cat", "dog", "cat", "tree"}
	strs = []string{"cat", "cat", "bird", "bird", "dog", "dog"}
	set := NewSet()

	for _, str := range strs {
		set.Insert(str)
		set.Println(os.Stdout)
	}

	set.Println(os.Stdout)
}
