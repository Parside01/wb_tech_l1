package main

import (
	"fmt"
	"io"
	"os"
	"sync"
)

func MoveArrayToChannel(arr []int, in chan int) {
	for _, i := range arr {
		in <- i
	}
	close(in)
}

func ProcessMessagesToChannel(in <-chan int, out chan int) {
	for msg := range in {
		out <- msg * 2
	}
	close(out)
}

func PrintProcessedMessages(out chan int, w io.Writer) {
	for msg := range out {
		if _, err := fmt.Fprint(w, msg); err != nil {
			panic(err)
		}
	}
}

func main() {
	arrSize := 100
	arr := make([]int, arrSize)
	for i := range arr {
		arr[i] = i
	}

	wg := sync.WaitGroup{}

	in := make(chan int)
	out := make(chan int)

	wg.Add(1)
	go func() {
		defer wg.Done()
		MoveArrayToChannel(arr, in)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		ProcessMessagesToChannel(in, out)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		PrintProcessedMessages(out, os.Stdout)
	}()

	wg.Wait()
}
