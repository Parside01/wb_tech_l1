package main

import "fmt"

func Reverse(s string) string {
	res := []rune(s)
	l, r := 0, len(res)-1
	for l < r {
		res[l], res[r] = res[r], res[l]
		l++
		r--
	}
	return string(res)
}

func main() {
	fmt.Println(Reverse("главрыба"))
}
