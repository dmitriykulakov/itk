package main

import (
	"fmt"
)

type StringIntMap struct {
	elem map[string]int
}

func (m StringIntMap) Add(key string, value int) {
	m.elem[key] = value
}

func (m StringIntMap) Remove(key string) {
	delete(m.elem, key)
}

func (m StringIntMap) Copy() map[string]int {
	result := make(map[string]int, len(m.elem))
	for key, value := range m.elem {
		result[key] = value
	}
	return result
}

func (m StringIntMap) Exists(key string) bool {
	_, result := m.elem[key]
	return result
}

func (m StringIntMap) Get(key string) (int, bool) {
	result, ok := m.elem[key]
	return result, ok
}

func main() {
	var test StringIntMap
	test.elem = make(map[string]int)
	test.Add("hello", 100)
	fmt.Println("after add", test)
	test.Add("hello1", 1000)
	fmt.Println("after add", test)
	test.Remove("hello")
	fmt.Println("after remove", test)
	result := test.Copy()
	fmt.Println("new after copy", result)
	test.elem["key"] = 10
	fmt.Println("old after copy and add", test)
	fmt.Println("new after copy and add", result)
	fmt.Println("exist hello", test.Exists("hello"))
	fmt.Println("exist key", test.Exists("key"))
	value, ok := test.Get("hello")
	fmt.Println("get hello", value, ok)
	value, ok = test.Get("key")
	fmt.Println("get key", value, ok)
}
