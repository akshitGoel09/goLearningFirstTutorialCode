package main

import (
	"fmt"
)

func main() {
	loopsExample()
}

func loopsExample() {
	loopTillCondition()
	infiniteLoop()
	loopOverCollections()
}

func loopOverCollections() {
	fmt.Println("loop over collections")

	slice := []int{1, 2, 3}

	//Basic way
	for i := 0; i < len(slice); i++ {
		println(slice[i])
	}

	fmt.Println("Better way")

	//Better way
	for i, v := range slice {
		println(i, v)
	}

	fmt.Println("Similar way for map")

	wellKnownPorts := map[string]int{ "http" : 80, "https" : 443}

	for k, v := range wellKnownPorts{
		println(k, v)
	}
}

func infiniteLoop() {
	fmt.Println("infinite loop example")
	var i int
	for {
		if i == 5 {
			break
		}
		i++
		println(i)
	}
}

func loopTillCondition() {
	fmt.Println("loop till condition example")
	for i := 0; i < 5; i++ {
		println(i)
		if i == 3 {
			break
			//continue
		}
		println("continuing...")
	}
}
