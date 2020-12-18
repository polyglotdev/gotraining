package main

import "fmt"

// user represents a user in the system.
type user struct {
	name  string
	email string
}

// NewUser represents the entry point for PointersEscape
func PointersEscape() {
	u1 := createUserV1()
	u2 := createUserV2()

	fmt.Println("u1", &u1, "u2", u2)
}

// createUserV1 creates a user value and passess
// a copy back to the caller.
//go:noinline
func createUserV1() user {
	u := user{name: "Bill", email: "bill@gmail.com"}
	fmt.Println("V1:", u)
	return u
}

// createUserV1 creates a user value and passess
// a copy back to the caller.
//go:noinline
func createUserV2() *user {
	u := user{name: "Dom", email: "dom@gmail.com"}
	fmt.Println("V2:", &u)
	return &u
}
