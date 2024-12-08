package main

import (
	"context"
	"fmt"
	"github.com/brianvoe/gofakeit/v7"
	"io"
	"runtime"
	"sync"
	"time"
)

func GoroutineStopUseChannel(w io.Writer, stopMessage string) {
	stop := make(chan struct{})
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func(stopChan <-chan struct{}) {
		defer wg.Done()
		for {
			select {
			case <-stopChan:
				if _, err := fmt.Fprintf(w, "%s", stopMessage); err != nil {
					panic(err)
				}
				return
			}
		}
	}(stop)
	stop <- struct{}{}
	wg.Wait()
}

func GoroutineStopUseCloseChannel(w io.Writer, messages []string, stopMessage string) {
	msgs := make(chan string)
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			select {
			case msg, ok := <-msgs:
				if !ok {
					if _, err := fmt.Fprintf(w, "%s", stopMessage); err != nil {
						panic(err)
					}
					return
				}
				if _, err := fmt.Fprintf(w, "%s", msg); err != nil {
					panic(err)
				}
			}
		}
	}()

	for _, msg := range messages {
		msgs <- msg
	}
	close(msgs)
	wg.Wait()
}

func GoroutineStopUseVariables(w io.Writer, stopMessage string) {
	stopped := make(map[string]bool)
	wg := sync.WaitGroup{}
	mutex := sync.RWMutex{}
	uuid := gofakeit.UUID()

	stopped[uuid] = false
	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			mutex.RLock()
			if stopped[uuid] {
				mutex.RUnlock()
				if _, err := fmt.Fprintf(w, "%s", stopMessage); err != nil {
					panic(err)
				}
				return
			}
		}
	}()

	stopped[uuid] = true
	wg.Wait()
}

// Ну просто context.Done() везде, так что случаи разбирать не надо.
func GoroutineStopUseContext(w io.Writer, stopMessage string) {
	wg := sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				if _, err := fmt.Fprintf(w, "%s", stopMessage); err != nil {
					panic(err)
				}
				return
			}
		}
	}()
	cancel()
	wg.Wait()
}

func GoroutineStopUseGoexit(w io.Writer, stopMessage string) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			if _, err := fmt.Fprintf(w, "%s", stopMessage); err != nil {
				panic(err)
			}
			runtime.Goexit()
		}
	}()
	wg.Wait()
}

func GoroutineStopUseCond(w io.Writer, stopMessage string) {
	wg := sync.WaitGroup{}
	mutex := sync.Mutex{}
	cond := sync.NewCond(&mutex)

	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			cond.L.Lock()
			defer cond.L.Unlock()

			cond.Wait()

			if _, err := fmt.Fprintf(w, "%s", stopMessage); err != nil {
				panic(err)
			}
			return
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(10 * time.Millisecond) // Ну чтобы точно вызвалось cond.Wait

		cond.L.Lock()
		defer cond.L.Unlock()

		cond.Broadcast()
	}()

	wg.Wait()
}
