package main

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestRandom1(t *testing.T) {
	ch := randomGenerator(time.Second * 2)
	counter := 0
	for elem := range ch {
		fmt.Println(elem)
		counter++
		if counter > 3 {
			time.Sleep(time.Second * 5)
		}
	}
	_, close := <-ch
	if !close {
		t.Logf("Test: OK\n")
	} else {
		t.Errorf("Test FAIL\n")
	}
}

func TestRandom2(t *testing.T) {
	counter := 0
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ch := randomGenerator2(ctx)
	for elem := range ch {
		fmt.Println(elem)
		counter++
		if counter > 3 {
			cancel()
		}
	}
	_, close := <-ch
	if !close {
		t.Logf("Test: OK\n")
	} else {
		t.Errorf("Test FAIL\n")
	}
}
