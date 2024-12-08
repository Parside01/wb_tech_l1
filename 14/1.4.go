package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/brianvoe/gofakeit/v7"
	"io"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func Worker(ctx context.Context, writer io.Writer, in <-chan string) {
	for {
		select {
		case <-ctx.Done():
			return
		case job := <-in:
			if _, err := fmt.Fprintf(writer, "%s", job); err != nil {
				panic(err)
			}
		}
	}
}

func StartWorkers(ctx context.Context, writer io.Writer, workersCount int, in <-chan string) {
	wg := sync.WaitGroup{}
	for i := 0; i < workersCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			Worker(ctx, writer, in)
		}()
	}
	wg.Wait()
}

func startMessageGenerator(ctx context.Context, in chan string, ticker *time.Ticker) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			msg := gofakeit.Sentence(16)
			in <- msg
		}
	}
}

func RunDataProcessor(workersCount int, rate time.Duration, writer io.Writer) {
	msgs := make(chan string)
	ctx, cancel := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		StartWorkers(ctx, writer, workersCount, msgs)
	}()

	ticker := time.NewTicker(rate)

	wg.Add(1)
	go func() {
		defer wg.Done()
		startMessageGenerator(ctx, msgs, ticker)
	}()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	<-signals
	cancel()
	wg.Wait()
	close(msgs)
}

var (
	workersCount int
	rate         time.Duration
)

func main() {
	flag.IntVar(&workersCount, "workers", 10, "Number of workers")
	flag.DurationVar(&rate, "rate", time.Second*5, "Rate of workers")
	flag.Parse()
	RunDataProcessor(workersCount, rate, os.Stdout)
}
