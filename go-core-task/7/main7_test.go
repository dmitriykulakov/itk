package main

import (
	"reflect"
	"slices"
	"sync"
	"testing"
)

func TestMultiplexChannels(t *testing.T) {
	test := []struct {
		channels []chan interface{}
		slices   [][]interface{}
		want     []interface{}
	}{
		{[]chan interface{}{make(chan interface{}), make(chan interface{})},
			[][]interface{}{{1, 2, 3}, {1, 2, 3}},
			[]interface{}{1, 1, 2, 2, 3, 3}},
		{[]chan interface{}{make(chan interface{}), make(chan interface{}), make(chan interface{})},
			[][]interface{}{{1, "hello", true}, {1.12, -2, [3]int{1, 2, 3}}, {77, 88}},
			[]interface{}{1.12, -2, 1, "hello", [3]int{1, 2, 3}, true, 77, 88}},
		{[]chan interface{}{}, [][]interface{}{}, []interface{}{}}}

	for i, r := range test {
		slice := make([]interface{}, 0)
		var wg sync.WaitGroup
		result := multiplexChannels(r.channels...)
		wg.Add(1)
		go func() {
			for value := range result {
				slice = append(slice, value)
			}
			wg.Done()
		}()

		for j, ch := range r.channels {
			go func() {
				for _, value := range r.slices[j] {
					ch <- value
				}
				close(ch)
			}()
		}
		wg.Wait()

		flag := true
		if len(r.want) == len(slice) {
			for _, value := range r.want {
				if !slices.ContainsFunc(slice, func(n interface{}) bool {
					if reflect.TypeOf(value) == reflect.TypeOf(n) {
						if value == n {
							return true
						}
					}
					return false
				}) {
					flag = false
					continue
				}
			}
		} else {
			flag = false
		}
		if flag {
			t.Logf("Test %d: OK, expected %v, result %v", i+1, r.want, slice)
		} else {
			t.Errorf("Test %d FAIL: expected %v, result %v", i+1, r.want, slice)
		}
	}
}
