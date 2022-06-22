package main

import (
	"fmt"

	"github.com/hannydevelop/golang_submodules/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
