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
	// pointers()
	constants()
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

func constants() {
	const pi = 3.1415
	fmt.Println(pi)
	// pi = 2.14 Leads to error as it is constant

	const c = 3

	fmt.Println(c + 3)

	//A bunch of code

	fmt.Println(3 + 1.2)

	//const with specific type
	const d int = 3

	fmt.Println(d + 3)

	//We need to type cast it to fload32 and then add 1.2 now earlier we could do it without it
	//Check typecasting is a little different from java
	fmt.Println(float32(d) + 1.2)

}
