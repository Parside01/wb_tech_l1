package main

import (
	"fmt"
	"strings"
)

func ReverseWords(s string) string {
	words := strings.Split(s, " ")
	l, r := 0, len(words)-1
	for l < r {
		words[l], words[r] = words[r], words[l]
		l++
		r--
	}
	return strings.Join(words, " ")
}

func main() {
	fmt.Println(ReverseWords("snow dog sun"))
}
