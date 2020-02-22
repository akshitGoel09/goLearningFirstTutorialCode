package main

import "fmt"

func main() {
	// array()
	slice()
}

func slice() {

	arr := [3]int{1, 2, 3}

	//Create a slice from arr from beginning of the arr to the end of the array
	//Here slice is just pointing to the arr
	//So any changes to the arr and slice will get reflected in the slice
	slice := arr[:]

	fmt.Println(arr, slice)

	arr[2] = 23

	fmt.Println(arr, slice)

	slice2 := []int{1, 2, 3}

	fmt.Println(slice2)

	slice2 = append(slice2, 4, 42, 27)

	fmt.Println(slice2)

	slice3 := slice2[1:]
	slice4 := slice2[:2]
	slice5 := slice2[1:2]

	fmt.Println(slice2, slice3, slice4, slice5)
}

func array() {
	fmt.Println("\n array examples")
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3

	fmt.Println(arr)
	fmt.Println(arr[2])

	arr2 := [3]int{1, 2, 3}

	fmt.Println(arr2)
}
