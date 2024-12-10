package main

import (
	"fmt"
	"math/rand"
)

func getRandomSlice() []int {
	result := make([]int, 10)
	for i := range result {
		result[i] = rand.Int()
	}
	return result
}

func sliceExample(input []int) []int {
	result := make([]int, 0, len(input))
	for _, value := range input {
		if value%2 == 0 {
			result = append(result, value)
		}
	}
	return result
}

func addElements(input []int, value int) []int {
	result := make([]int, len(input), len(input)+1)
	copy(result, input)
	result = append(result, value)
	return result
}

func copySlice(input []int) []int {
	result := make([]int, len(input))
	copy(result, input)
	return result

}

func removeElement(input []int, index uint) ([]int, error) {
	if len(input) <= int(index) {
		return []int{}, fmt.Errorf("index %d is out of range", index)
	}
	return append(input[:index], input[(index+1):]...), nil
}

func main() {
	originalSlice := getRandomSlice()
	fmt.Println(originalSlice)
	sliceEven := sliceExample(originalSlice)
	fmt.Println(sliceEven)
	addSlice := addElements(sliceEven, 100)
	fmt.Println(addSlice)
	newSlice := copySlice(addSlice)
	if len(addSlice) > 0 {
		addSlice[0] = 101
		fmt.Println(addSlice)
		fmt.Println(newSlice)
	}
	afterRemove, err := removeElement(newSlice, 2)
	fmt.Println(afterRemove, err)
	afterRemove2, err := removeElement(afterRemove, 100)
	fmt.Println(afterRemove)
	fmt.Println(afterRemove2, err)
}
