package main

import (
	"errors"
	"fmt"
)

func main() {

	port := 3000
	err := startWebServer(port, 2)
	fmt.Println(err)

	port2, err := startWebServerMultipleValues(port, 2)
	fmt.Println(port2, err)

	//For multiple return but if we don't want to use a particular variable
	_, err3 := startWebServerMultipleValues(port, 2)
	fmt.Println(err3)
}

func startWebServer(port, numberOfRetries int) error {
	fmt.Println("Starting server...")

	//important stuff
	fmt.Println("server started on port", port)
	fmt.Println("Number of retries", numberOfRetries)
	return errors.New("Something went wrong")
}

func startWebServerMultipleValues(port, numberOfRetries int) (int, error) {
	fmt.Println("Starting server...")

	//important stuff
	fmt.Println("server started on port", port)
	fmt.Println("Number of retries", numberOfRetries)
	return port, nil
}
