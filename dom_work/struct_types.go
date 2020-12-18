package main

import "fmt"

type example struct {
	flag    bool
	counter int16
	pi      float32
}
type bill struct {
	flag    bool
	counter int16
	pi      float32
}
type nancy struct {
	flag    bool
	counter int16
	pi      float32
}

func StructTypes() {
	// construct a value of type example
	var e1 example

	fmt.Printf("%+v\n", e1)

	e2 := example{flag: true, counter: 10, pi: 3.14159265358979323846}
	fmt.Printf("%+v\n", e2)

	var e3 struct {
		flag    bool
		counter int
		pi      float64
	}

	e3.flag = true
	e3.counter = 12345
	e3.pi = 3.14159265358979323846

	fmt.Printf("e3:%+v\n", e3)

	var b bill
	var n nancy
	b = bill(n)
	fmt.Println(b, n)
	fmt.Printf("%p\n", &b)
	fmt.Printf("%p\n", &n)

	fruits := [4]string{"Apple", "Banana", "Cherry", "Strawberry"}

	for _, fruit := range fruits {
		fmt.Println(fruit)
	}
}
