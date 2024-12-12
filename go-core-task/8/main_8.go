package main

import (
	"fmt"
	"sync"
	"time"
)

type customWaitGroup struct {
	counter int
	mu      sync.Mutex
}

func (wg *customWaitGroup) Add(n int) {
	wg.mu.Lock()
	defer wg.mu.Unlock()
	wg.counter += n
	if n < 0 {
		panic("customWaitGroup: negative WaitGroup counter")
	}
}

func (wg *customWaitGroup) Done() {
	wg.mu.Lock()
	defer wg.mu.Unlock()
	wg.counter--
	if wg.counter < 0 {
		panic("customWaitGroup: negative WaitGroup counter")
	}
}

func (wg *customWaitGroup) Wait() {
	for wg.counter > 0 {
		time.Sleep(time.Microsecond)
	}
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in", r)
		}
	}()
	var wg customWaitGroup
	wg.Add(5)
	fmt.Println(wg.counter)
	wg.Done()
	fmt.Println(wg.counter)
	go func() {
		wg.Done()
		wg.Done()
		wg.Done()
		wg.Done()
	}()
	wg.Wait()
	fmt.Println(wg.counter)
	wg.Done()
}
