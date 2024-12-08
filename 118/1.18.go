package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	value int
	mutex sync.RWMutex
}

func NewCounter() *Counter {
	return &Counter{
		value: 0,
		mutex: sync.RWMutex{},
	}
}

func (c *Counter) Inc() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.value++
}

func (c *Counter) Dec() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.value--
}

func (c *Counter) Value() int {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.value
}

func (c *Counter) Reset() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.value = 0
}

func main() {
	counter := NewCounter()
	wg := sync.WaitGroup{}

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Inc()
		}()
	}
	wg.Wait()
	fmt.Println(counter.Value())
}
