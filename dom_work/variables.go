package main

import "fmt"

func Variables() {
	var a int
	var b string
	var c float64
	var d bool

	fmt.Printf("var a int \t %T [%v]\n", a, a)
	fmt.Printf("var b string \t %T [%v]\n", b, b)
	fmt.Printf("var c float64 \t %T [%v]\n", c, c)
	fmt.Printf("var d bool \t %T [%v]\n", d, d)

	aa := 10
	bb := "hello there"
	cc := 3.14159
	dd := false

	fmt.Println("\n")

	fmt.Printf("var aa int \t %T [%v]\n", aa, aa)
	fmt.Printf("var bb string \t %T [%v]\n", bb, bb)
	fmt.Printf("var cc float64 \t %T [%v]\n", cc, cc)
	fmt.Printf("var dd bool \t %T [%v]\n", dd, dd)

	// specify type and perform a conversion
	aaa := int32(10)
	fmt.Printf("aaa := int32(10) %T [%v]\n", aaa, aaa)
}
