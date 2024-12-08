package main

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v7"
)

func QuickSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	quickSortImp(arr, 0, len(arr)-1)
}

func quickSortImp(arr []int, left, right int) {
	i, j := left, right

	// Чтобы нельзя было создать тест, который будет выполняться за O(n^2).
	someElem := arr[gofakeit.IntN(len(arr))]

	for i <= j {
		for arr[i] < someElem {
			i++
		}
		for arr[j] > someElem {
			j--
		}
		if i <= j {
			arr[i], arr[j] = arr[j], arr[i]
			i++
			j--
		}
	}
	if j > left {
		quickSortImp(arr, left, j)
	}
	if i < right {
		quickSortImp(arr, i, right)
	}
}

func main() {
	arr := []int{5, 1, 3, 5}
	QuickSort(arr)
	fmt.Println(arr)
}
