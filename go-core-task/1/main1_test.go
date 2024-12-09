package main

import (
	"fmt"
	"testing"
)

func TestStringBuilder(t *testing.T) {
	test := []struct {
		values []interface{}
		want   string
	}{
		{[]interface{}{42, 052, 0x2A, 3.14, "Golang", true, complex64(1 + 2i)}, "4242423.14Golangtrue(1+2i)"},
		{[]interface{}{}, ""},
		{[]interface{}{float32(3.14)}, ""},
		{[]interface{}{-25, "", -25}, "-25-25"},
		{[]interface{}{int32(25), "", int64(25), false, uint(11)}, "false"},
	}
	for i, r := range test {
		str := stringBuilder(r.values...)
		if str != r.want {
			t.Errorf("Test %d: FAIL, expected %s, result %s", i+1, r.want, str)
		} else {
			t.Logf("Test %d: OK, expected %s, result %s: OK", i+1, r.want, str)
		}
	}
}

func TestRuneConv(t *testing.T) {
	test := []struct {
		str  string
		want []rune
	}{
		{"4242423.14Golangtrue(1+2i)", []rune{'4', '2', '4', '2', '4', '2', '3', '.', '1', '4', 'G', 'o', 'l', 'a', 'n', 'g', 't', 'r', 'u', 'e', '(', '1', '+', '2', 'i', ')'}},
		{"", []rune{}},
		{"false", []rune{'f', 'a', 'l', 's', 'e'}},
	}
	for i, r := range test {
		runeTmp := runeConv(r.str)
		if len(runeTmp) != len(r.want) {
			t.Errorf("Test %d: FAIL, expected %v, result %v", i+1, r.want, runeTmp)
			continue
		}
		flag := true
		for i := 0; i < len(runeTmp) && flag; i++ {
			if runeTmp[i] != r.want[i] {
				t.Errorf("Test %d: FAIL, expected %v, result %v", i+1, r.want, runeTmp)
				flag = false
			}
		}
		if flag {
			t.Logf("Test %d: OK, expected %v, result %v: OK", i+1, r.want, runeTmp)
		}
	}
}

func TestHash(t *testing.T) {
	test := []struct {
		value []rune
		want  string
	}{
		{[]rune{'4', '2', '4', '2', '4', '2', '3', '.', '1', '4', 'G', 'o', 'l', 'a', 'n', 'g', 't', 'r', 'u', 'e', '(', '1', '+', '2', 'i', ')'}, "53f2f60ac6c41389d3ed3d84d88d8c2860bf8981c677be18243a6f35a6b6a1b3"},
		{[]rune{}, "66802df107aace17871a5b610ff9eb11706e13477bb24e93966ca80671c0fac6"},
	}
	for i, r := range test {
		str := fmt.Sprintf("%x", hash(r.value))
		if str != r.want {
			t.Errorf("Test %d: FAIL, expected %s, result %s", i+1, r.want, str)
		} else {
			t.Logf("Test %d: OK, expected %s, result %s: OK", i+1, r.want, str)
		}
	}
}
