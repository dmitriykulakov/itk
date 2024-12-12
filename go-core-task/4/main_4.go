package main

import "fmt"

func getCommonSlice(slice1 []string, slice2 []string) []string {
	result := make([]string, 0, len(slice1))
	mapa := make(map[string]struct{}, len(slice1))
	for _, str := range slice1 {
		mapa[str] = struct{}{}
	}
	for _, str := range slice2 {
		delete(mapa, str)
	}
	for str := range mapa {
		result = append(result, str)
	}
	return result
}

func main() {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}
	fmt.Println(getCommonSlice(slice1, slice2))
}
