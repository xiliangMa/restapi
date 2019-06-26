package main

import "testing"

func TestA(t *testing.T){
	a := "aabbccdde"
	cache := make(map[string]int)

	for _ , value := range a {
		println(string(value))
		if _, ok := cache[string(value)]; ok {
			cache[string(value)] = cache[string(value)] + 1
		} else {
			cache[string(value)] = 1
		}

	}
	t.Log("========", cache)
}