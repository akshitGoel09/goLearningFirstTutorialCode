package main

import "fmt"

func main() {
	// panicExample()
	// ifStatementExample()
	switchStatementExample()
}

func switchStatementExample() {
	type HTTPRequest struct {
		Method string
	}

	r1 := HTTPRequest{
		Method: "GET",
	}

	switch r1.Method {
	case "GET":
		fmt.Println("GET request")
		fallthrough
	case "POST":
		fmt.Println("POST request")
	case "DELETE":
		fmt.Println("DELETE request")
	case "PUT":
		fmt.Println("PUT request")
	default:
		fmt.Println("Unhandled method")
	}
}

func ifStatementExample() {
	type User struct {
		ID        int
		firstName string
		lastName  string
	}

	u1 := User{
		ID:        1,
		firstName: "Akshit",
		lastName:  "Goel",
	}

	u2 := User{
		ID:        2,
		firstName: "Kehav",
		lastName:  "Singhal",
	}

	if u1.ID == u2.ID {
		fmt.Println("Same user")
	} else if u1.firstName == u2.firstName {
		fmt.Println("Similar user")
	} else {
		fmt.Println("Different users")
	}

}

func panicExample() {
	fmt.Println("starting web server")

	panic("Something bad just happened")

	fmt.Println("web server started")
}
