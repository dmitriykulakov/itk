package main

import (
	"fmt"
	"math"
)

func conveyor(slice []uint8) chan float64 {
	first := make(chan uint8)
	second := make(chan float64)

	go func() {
		for _, elem := range slice {
			first <- elem
		}
		close(first)
	}()
	go func() {
		for elem := range first {
			second <- math.Pow(float64(elem), 3.0)
		}
		close(second)
	}()
	return second
}

func main() {
	second := conveyor([]uint8{0, 1, 2, 3, 4, 5})
	for elem := range second {
		fmt.Println(elem)
	}
}
