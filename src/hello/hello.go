package main

import (
	"fmt"
)

func main() {
	simplePrint()
	primitiveDataTypes()
}

func simplePrint() {
	fmt.Printf("hello, world\n")
}

func primitiveDataTypes() {
	var i int
	i = 42
	fmt.Println(i)

	var f float32 = 3.14
	fmt.Println(f)

	firstName := "Akshit"
	fmt.Println(firstName)

	b := true
	fmt.Println(b)

	c := complex(3, 4)
	fmt.Println(c)

	r, im := real(c), imag(c)
	fmt.Println(r, im)
}
