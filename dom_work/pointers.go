package main

import "fmt"

func Pointers() {
	count := 10

	fmt.Println("count:\tValue Of[", count, "]\tAddr of[", &count, "]")
	increment(count)
	fmt.Println("count:\tValue Of[", count, "]\tAddr of[", &count, "]")
}

func increment(inc int) {
	inc++
	fmt.Println("inc:\tValue Of[", inc, "]\tAddr of[", &inc, "]")
}
