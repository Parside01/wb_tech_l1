package main

import "fmt"

func DeleteElement(slice []int, index int) []int {
	if index >= len(slice) || index < 0 {
		return slice
	}
	return append(slice[:index], slice[index+1:]...)
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(DeleteElement(slice, 1))
}
