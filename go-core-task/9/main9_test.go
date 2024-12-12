package main

import (
	"slices"
	"testing"
)

func TestConveyor(t *testing.T) {
	test := []struct {
		input []uint8
		want  []float64
	}{
		{[]uint8{0, 1, 2, 3, 4, 5}, []float64{0, 1, 8, 27, 64, 125}},
		{[]uint8{5, 4, 3, 2, 1, 0}, []float64{125, 64, 27, 8, 1, 0}},
		{[]uint8{}, []float64{}}}
	for i, r := range test {
		result := make([]float64, 0, len(r.want))
		second := conveyor(r.input)
		for elem := range second {
			result = append(result, elem)
		}
		if len(result) != len(r.want) {
			t.Errorf("Test %d FAIL: expected %v, result %v", i+1, r.want, result)
		} else {
			flag := true
			for _, value := range r.want {
				if !slices.Contains(result, value) {
					flag = false
					continue
				}
			}
			if flag {
				t.Logf("Test %d: OK, expected %v, result %v,: OK", i+1, r.want, result)
			} else {
				t.Errorf("Test %d FAIL: expected %v, result %v", i+1, r.want, result)
			}
		}
	}
}
