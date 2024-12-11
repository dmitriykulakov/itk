package main

import (
	"fmt"
	"testing"
)

func TestWaitGroup(t *testing.T) {
	test := []struct {
		n int
	}{
		{0},
		{5}}
	for i, r := range test {
		var wg customWaitGroup
		wg.Add(r.n)
		go func() {
			for i := 0; i < r.n; i++ {
				wg.Done()
			}
		}()
		wg.Wait()
		if wg.counter == 0 {
			t.Logf("Test %d: OK", i+1)
		} else {
			t.Errorf("Test %d: )Fail", i+1)
		}
	}
}

func TestAdd(t *testing.T) {
	test := []struct {
		n int
	}{
		{0},
		{5}}
	for i, r := range test {
		var wg customWaitGroup
		wg.Add(r.n)
		if wg.counter == r.n {
			t.Logf("Test %d: OK", i+1)
		} else {
			t.Errorf("Test %d: )Fail", i+1)
		}
	}
}

func TestDone(t *testing.T) {
	test := []struct {
		n int
	}{
		{5}}
	for i, r := range test {
		var wg customWaitGroup
		wg.Add(r.n)
		for j := 0; j < r.n; j++ {
			if wg.counter == r.n-i {
				t.Logf("Test %d: OK", i+1)
			} else {
				t.Errorf("Test %d: )Fail", i+1)
			}
		}
	}
}

func TestDonePanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in", r)
			t.Logf("Test: OK")
		}
	}()
	var wg customWaitGroup
	wg.Done()
}

func TestAddPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in", r)
			t.Logf("Test: OK")
		}
	}()
	var wg customWaitGroup
	wg.Add(-1)
}
