package main

import (
	"fmt"
	"slices"
	"testing"
)

func TestGetRandomSlice(t *testing.T) {
	for i := 0; i < 5; i++ {
		s1 := getRandomSlice()
		s2 := getRandomSlice()
		fmt.Println(s1)
		fmt.Println(s2)
		if slices.Equal(s1, s2) {
			t.Errorf("Test %d: FAIL, the slices are equal", i+1)
		} else if len(s1) != 10 || len(s2) != 10 {
			t.Errorf("Test %d: FAIL, the slice has a wrong size", i+1)
		} else {
			t.Logf("Test %d: OK, slices have the len = 10 and they are not equal", i)
		}
	}
}

func TestSliceExample(t *testing.T) {
	test := []struct {
		slice []int
		want  []int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, []int{2, 4, 6, 8, 10}},
		{[]int{10, 2, 30, 4, 50, 6, 70, 8, 90, 10}, []int{10, 2, 30, 4, 50, 6, 70, 8, 90, 10}},
		{[]int{1, 5, 7, 9}, []int{}},
		{[]int{}, []int{}},
	}
	for i, r := range test {
		result := sliceExample(r.slice)
		if slices.Equal(r.want, result) {
			t.Logf("Test %d: OK, expected %v, result %v: OK", i+1, r.want, result)
		} else {
			t.Errorf("Test %d: FAIL, expected %v, result %v", i+1, r.want, result)
		}
	}
}

func TestAddElement(t *testing.T) {
	test := []struct {
		slice []int
		value int
		want  []int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 11, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}},
		{[]int{10, 2, 30, 4, 50, 6, 70, 8, 90, 10}, 0, []int{10, 2, 30, 4, 50, 6, 70, 8, 90, 10, 0}},
		{[]int{1, 5, 7, 9}, 5, []int{1, 5, 7, 9, 5}},
		{[]int{}, 1, []int{1}},
	}
	for i, r := range test {
		result := addElements(r.slice, r.value)
		if slices.Equal(r.want, result) {
			t.Logf("Test %d: OK, expected %v, result %v: OK", i+1, r.want, result)
		} else {
			t.Errorf("Test %d: FAIL, expected %v, result %v", i+1, r.want, result)
		}
	}
}

func TestCopySlice(t *testing.T) {
	test := []struct {
		slice []int
		want  []int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{[]int{10, 2, 30, 4, 50, 6, 70, 8, 90, 10}, []int{10, 2, 30, 4, 50, 6, 70, 8, 90, 10}},
		{[]int{1, 5, 7, 9}, []int{1, 5, 7, 9}},
		{[]int{}, []int{}},
	}
	for i, r := range test {
		result := copySlice(r.slice)
		if slices.Equal(r.want, result) {
			t.Logf("Test %d: OK, expected %v, result %v: OK", i+1, r.want, result)
		} else {
			t.Errorf("Test %d: FAIL, expected %v, result %v", i+1, r.want, result)
		}
		if len(r.slice) > 0 {
			r.slice[0] = -1
			if !slices.Equal(r.slice, result) {
				t.Logf("Test %d: OK, there are different slices", i+1)
			} else {
				t.Errorf("Test %d: FAIL, there aren't different slices", i+1)
			}
		}

	}
}

func TestRemoveElement(t *testing.T) {
	test := []struct {
		slice []int
		index uint
		want  []int
		err   error
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 0, []int{2, 3, 4, 5, 6, 7, 8, 9, 10}, nil},
		{[]int{10, 2, 30, 4, 50, 6, 70, 8, 90, 10}, 9, []int{10, 2, 30, 4, 50, 6, 70, 8, 90}, nil},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5, []int{1, 2, 3, 4, 5, 7, 8, 9, 10}, nil},
		{[]int{1}, 0, []int{}, nil},
		{[]int{}, 0, []int{}, fmt.Errorf("index %d is out of range", 0)},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10, []int{}, fmt.Errorf("index %d is out of range", 10)},
	}
	for i, r := range test {
		result, err := removeElement(r.slice, r.index)
		if r.err != nil {
			if slices.Equal(r.want, result) && err.Error() == r.err.Error() {
				t.Logf("Test %d: OK, slice expected %v, result %v, err expected %s, result %s: OK", i+1, r.want, result, r.err, err)
			} else {
				t.Errorf("Test %d: FAIL, slice expected %v, result %v, err expected %s, result %s", i+1, r.want, result, r.err, err)
			}
		} else if slices.Equal(r.want, result) {
			t.Logf("Test %d: OK, slice expected %v, result %v, err expected %p, result %p: OK", i+1, r.want, result, r.err, err)
		} else {
			t.Errorf("Test %d: FAIL, slice expected %v, result %v, err expected %p, result %p", i+1, r.want, result, r.err, err)
		}
	}

}
