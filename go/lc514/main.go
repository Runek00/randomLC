package main

import "fmt"

var cache = make(map[string]int)

func findRotateSteps(ring string, key string) int {
	if len(key) == 0 {
		return 0
	}
	steps, ok := cache[ring + ":" + key]
	if ok {
		return steps
	}
	if len(key) == 1 && ring[0] == key[0] {
		return 1
	}
	out := len(ring) * len(key)

	for i := -(len(ring) / 2); i <= len(ring)/2; i++ {
		var k int
		if i < 0 {
			k = len(ring) + i
		} else {
			k = i
		}
		if ring[k] == key[0] {
			subring := findRotateSteps(ring[k:]+ring[0:k], key[1:])
			j := i
			if j < 0 {
				j = -j
			}
			if j+subring < out {
				out = j + subring
			}
		}
	}
	cache[ring + ":" + key] = out + 1
	return out + 1
}

func main() {
	fmt.Println(findRotateSteps("godding", "gd"))
	fmt.Println(findRotateSteps("godding", "godding"))
}
