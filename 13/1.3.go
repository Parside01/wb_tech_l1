package main

import (
	"fmt"
	"sync"
)

func SumSquaresConcurrency() int {
	results := make(chan int)
	numbers := []int{2, 4, 6, 8, 10}
	wg := sync.WaitGroup{}

	for _, num := range numbers {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			results <- num * num
		}(num)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	sum := 0
	for square := range results {
		sum += square
	}
	return sum
}

func main() {
	fmt.Println(SumSquaresConcurrency())
}
