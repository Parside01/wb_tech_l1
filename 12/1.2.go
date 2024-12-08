package main

import (
	"fmt"
	"io"
	"os"
	"sync"
)

// По условию задачи нужно подставить сюда os.Stdout.
func CalculateSquaresConcurrency(w io.Writer) {
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

	for square := range results {
		if _, err := fmt.Fprintf(w, "%d ", square); err != nil {
			panic(err)
		}
	}
}

func main() {
	CalculateSquaresConcurrency(os.Stdout)
}
