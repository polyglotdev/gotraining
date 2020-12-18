// Package main provides main process for pointers_v1
package main

import "fmt"

func PointersV1() {
	fmt.Println("\nPointers V1")
	myCount := 10
	fmt.Println("count:\tValue Of[", myCount, "]\t\t\tAddr of[", &myCount, "]")

	addTo(&myCount)
}

// Define a function that takes in an action like increment or decrement and amount to prefrom action on
func addTo(inc *int) {
	*inc++
	fmt.Println("inc:\tValue Of[", inc, "]\tAddr of[", *inc, "]")
}

// func subtractFrom(inc int) {
// 	inc--
// 	fmt.Println("count:\tValue Of[", inc, "]\tAddr of[", &inc, "]")

// }

// func multiply(inc int) {
// 	inc *= inc
// 	fmt.Println("count:\tValue Of[", inc, "]\tAddr of[", &inc, "]")
// }

// func divide(inc int) {
// 	inc /= inc
// 	fmt.Println("count:\tValue Of[", inc, "]\tAddr of[", &inc, "]")
// }
