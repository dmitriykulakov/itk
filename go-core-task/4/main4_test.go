package main

import (
	"slices"
	"testing"
)

func TestExists(t *testing.T) {
	test := []struct {
		slice1 []string
		slice2 []string
		want   []string
	}{
		{[]string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}, []string{"banana", "date", "fig"}, []string{"apple", "cherry", "43", "lead", "gno1"}},
		{[]string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}, []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}, []string{}},
		{[]string{}, []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}, []string{}},
		{[]string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}, []string{}, []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}},
	}
	for i, r := range test {
		result := getCommonSlice(r.slice1, r.slice2)
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
		if flag {
			t.Logf("Test %d: OK, expected %v, result %v: OK", i+1, r.want, result)
		} else {
			t.Errorf("Test %d: FAIL, expected %v, result %v", i+1, r.want, result)
		}
	}
}
