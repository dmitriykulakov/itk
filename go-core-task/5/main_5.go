package main

import "fmt"

func contains(slice1 []int, slice2 []int) (bool, []int) {
	result := make([]int, 0, len(slice1))
	mapa := make(map[int]bool, len(slice1))
	for _, number := range slice1 {
		mapa[number] = false
	}
	for _, number := range slice2 {
		if _, ok := mapa[number]; ok {
			result = append(result, number)
		}
	}
	return len(result) > 0, result
}

func main() {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}
	c := []int{63, 2, 4, 43}
	fmt.Println(contains(a, b))
	fmt.Println(contains(a, c))
}
