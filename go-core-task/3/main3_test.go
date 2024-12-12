package main

import (
	"testing"
)

func TestExists(t *testing.T) {
	var mapa StringIntMap
	mapa.elem = map[string]int{"hello": 100, "key": 1000}
	test := []struct {
		key    string
		result bool
	}{
		{"hello", true},
		{"key", true},
		{"wrong", false},
	}
	for i, r := range test {
		if mapa.Exists(r.key) == r.result {
			t.Logf("Test %d: OK", i+1)
		} else {
			t.Errorf("Test %d: FAIL", i+1)
		}
	}
}

func TestAdd(t *testing.T) {
	var mapa StringIntMap
	mapa.elem = make(map[string]int)
	test := []struct {
		key   string
		value int
		len   int
		want  map[string]int
	}{
		{"hello", 100, 1, map[string]int{"hello": 100}},
		{"key", 100, 2, map[string]int{"key": 100}},
		{"hello", 100, 2, map[string]int{"hello": 100}},
	}
	for i, r := range test {
		mapa.Add(r.key, r.value)
		if len(mapa.elem) != r.len {
			t.Errorf("Test %d: FAIL, the mapa size expected %v, result %v", i+1, r.len, len(mapa.elem))
		} else if mapa.Exists(r.key) {
			t.Logf("Test %d: OK, the key %s exists", i+1, r.key)
		} else {
			t.Errorf("Test %d: FAIL, the key %s doesn't exist", i+1, r.key)
		}
	}
}

func TestRemove(t *testing.T) {
	var mapa StringIntMap
	mapa.elem = map[string]int{"hello": 100, "key": 1000}
	test := []struct {
		key  string
		len  int
		want map[string]int
	}{
		{"hello", 1, map[string]int{"hello": 100}},
		{"key", 0, map[string]int{"key": 100}},
		{"hello", 0, map[string]int{"hello": 100}},
	}
	for i, r := range test {
		mapa.Remove(r.key)
		if len(mapa.elem) != r.len {
			t.Errorf("Test %d: FAIL, the mapa size expected %v, result %v", i+1, r.len, len(mapa.elem))
		} else if !mapa.Exists(r.key) {
			t.Logf("Test %d: OK, the key %s doesn't exists", i+1, r.key)
		} else {
			t.Errorf("Test %d: FAIL, the key %s exists", i+1, r.key)
		}
	}
}

func TestGet(t *testing.T) {
	var mapa StringIntMap
	mapa.elem = map[string]int{"hello": 100, "key": 1000}
	test := []struct {
		key    string
		value  int
		result bool
	}{
		{"hello", 100, true},
		{"key", 1000, true},
		{"wrong", 0, false},
	}
	for i, r := range test {
		value, result := mapa.Get(r.key)
		if result == r.result && value == r.value {
			t.Logf("Test %d: OK", i+1)
		} else {
			t.Errorf("Test %d: FAIL", i+1)
		}
	}
}

func TestCopy(t *testing.T) {
	var mapa StringIntMap
	mapa.elem = make(map[string]int)
	var mapa1 StringIntMap
	mapa1.elem = map[string]int{"hello": 100, "key": 1000}
	test := []struct {
		mapa StringIntMap
		len  int
	}{
		{mapa, 0},
		{mapa1, 2},
	}
	for i, r := range test {
		mapaNew := r.mapa.Copy()
		if len(mapaNew) == r.len {
			t.Logf("Test %d: OK", i+1)
		} else {
			t.Errorf("Test %d: FAIL", i+1)
		}
		r.mapa.Add("test", 100)
		if len(mapaNew) != len(r.mapa.elem) {
			t.Logf("Test %d: OK", i+1)
		} else {
			t.Errorf("Test %d: FAIL", i+1)
		}
	}
}
