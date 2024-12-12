package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

// вариант с остановкой по таймауту, если значения из канала не принимаются определенное время
func randomGenerator(t time.Duration) chan int {
	result := make(chan int)
	go func() {
		for {
			select {
			case result <- rand.Int():
				time.Sleep(time.Second)
			case <-time.After(t):
				close(result)
				return
			}
		}
	}()
	return result
}

// вариант с остановкой по контексту
func randomGenerator2(ctx context.Context) chan int {
	result := make(chan int)
	go func() {
		for {
			select {
			case <-ctx.Done():
				close(result)
				return
			default:
				result <- rand.Int()
				time.Sleep(time.Second * 1)
			}
		}
	}()
	return result
}

func main() {
	ch := randomGenerator(time.Second * 3)
	counter := 0
	for elem := range ch {
		fmt.Println(elem)
		counter++
		if counter > 3 {
			time.Sleep(time.Second * 5)
		}
	}

	fmt.Println("channel is closed")
	counter = 0
	ctx, cancel := context.WithCancel(context.Background())
	ch2 := randomGenerator2(ctx)
	for elem := range ch2 {
		fmt.Println(elem)
		counter++
		if counter > 5 {
			cancel()
		}
	}

	fmt.Println("channel is closed")
}
