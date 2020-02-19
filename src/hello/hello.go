package main

import (
	"fmt"
)

/*
  Comment out the methods on which type you want to test
*/
func main() {
	// simplePrint()
	// primitiveDataTypes()
	pointers()
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

func pointers() {
	var firstName *string = new(string)

	*firstName = "Akshit"

	fmt.Println(firstName)
	fmt.Println(*firstName)

	lastName := "Goel"

	fmt.Println(lastName)

	ptr := &lastName
	fmt.Println(ptr, *ptr)

	lastName = "Aggarwal"
	fmt.Println(ptr, *ptr)
}
