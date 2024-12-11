package main

import (
	"fmt"
	"time"
)

type customWaitGroup struct {
	counter int
}

func (wg *customWaitGroup) Add(n int) {
	wg.counter += n
	if n < 0 {
		panic("customWaitGroup: negative WaitGroup counter")
	}
}

func (wg *customWaitGroup) Done() {
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
	fmt.Println(wg)
	wg.Done()
	fmt.Println(wg)
	go func() {
		wg.Done()
		wg.Done()
		wg.Done()
		wg.Done()
	}()
	wg.Wait()
	fmt.Println(wg)
	wg.Done()
}
