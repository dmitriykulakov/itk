package main

import (
	"fmt"
	"sync"
)

type Value struct {
	SD int
}

func main() {
	var wg sync.WaitGroup
	one := make(chan interface{})
	two := make(chan interface{})
	slice := make([]interface{}, 0)
	result := multiplexChannels(one, two)
	tmp := Value{SD: 15}
	wg.Add(1)
	go func() {
		for value := range result {
			slice = append(slice, value)
		}
		wg.Done()
	}()
	one <- 5
	one <- 5.44
	one <- "dsdadddd"
	two <- "dasd"
	two <- tmp
	one <- true
	close(one)
	close(two)
	wg.Wait()
	fmt.Println(slice)
}

func multiplexChannels(input ...chan interface{}) <-chan interface{} {
	var wg sync.WaitGroup
	output := make(chan interface{})
	for _, inputCh := range input {
		wg.Add(1)
		go func() {
			for elem := range inputCh {
				output <- elem
			}
			wg.Done()
		}()
	}
	go func() {
		wg.Wait()
		close(output)
	}()
	return output
}
