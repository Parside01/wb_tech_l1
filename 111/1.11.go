package main

import (
	"fmt"
	"sort"
)

func SetsIntersection(right []int, left []int) []int {
	set := make(map[int]bool)

	for _, v := range right {
		set[v] = true
	}

	res := []int{}

	for _, v := range left {
		if set[v] {
			res = append(res, v)
		}
	}
	sort.Ints(res)
	return res
}

func main() {
	fmt.Println(SetsIntersection([]int{2, 1, 2, 3, 4}, []int{1, 2, 3, 4}))
}
