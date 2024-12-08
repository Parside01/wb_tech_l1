package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/brianvoe/gofakeit/v7"
	"io"
	"os"
	"sync"
	"time"
)

var LifeTime time.Duration
var Rate time.Duration

func StartMessagePublisher(ctx context.Context, out chan string, rate time.Duration) {
	ticker := time.NewTicker(rate)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			out <- gofakeit.Sentence(5)
		}
	}
}

func StartMessageConsumer(ctx context.Context, in <-chan string, writer io.Writer) {
	for {
		select {
		case <-ctx.Done():
			return
		case msg, ok := <-in:
			if !ok {
				return
			}
			if _, err := fmt.Fprintln(writer, msg); err != nil {
				panic(err)
			}
		}
	}
}

func main() {
	flag.DurationVar(&LifeTime, "timeout", time.Second*5, "Programme lifetime")
	flag.DurationVar(&Rate, "rate", time.Millisecond*20, "Rate of workers")
	flag.Parse()

	wg := sync.WaitGroup{}
	in := make(chan string)

	ctx, cancel := context.WithTimeout(context.Background(), LifeTime)
	defer cancel()

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(in)
		StartMessagePublisher(ctx, in, Rate)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		StartMessageConsumer(ctx, in, os.Stdout)
	}()

	wg.Wait()
}
