package main

import (
	"slices"
	"testing"
)

func TestContains(t *testing.T) {
	test := []struct {
		slice1 []int
		slice2 []int
		want   []int
		result bool
	}{
		{[]int{65, 3, 58, 678, 64}, []int{64, 2, 3, 43}, []int{64, 3}, true},
		{[]int{65, 3, 58, 678, 64}, []int{64, 678, 58, 3, 65}, []int{64, 678, 58, 3, 65}, true},
		{[]int{65, 3, 58, 678, 64}, []int{66, 2, 4, 43}, []int{}, false},
		{[]int{65, 3, 58, 678, 64}, []int{}, []int{}, false},
		{[]int{}, []int{65, 3, 58, 678, 64}, []int{}, false},
		{[]int{}, []int{}, []int{}, false}}
	for i, r := range test {
		ok, result := contains(r.slice1, r.slice2)
		if len(result) == 0 && ok == r.result {
			t.Logf("Test %d: OK, expected %v, result %v, ok expected %v, result %v: OK", i+1, r.want, result, r.result, ok)
		} else {
			flag := true
			if len(r.want) == len(result) {
				for _, value := range r.want {
					if !slices.Contains(result, value) {
						flag = false
						continue
					}
				}
			} else {
				flag = false
			}
			if flag && ok == r.result {
				t.Logf("Test %d: OK, expected %v, result %v, ok expected %v, result %v: OK", i+1, r.want, result, r.result, ok)
			} else {
				t.Errorf("Test %d: expected %v, result %v, ok expected %v, result %v: OK", i+1, r.want, result, r.result, ok)
			}
		}
	}
}
