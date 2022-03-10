package main

import (
	"fmt"

	"github.com/rwxrob/term"
)

func main() {
	fmt.Print("Enter sample text: ")
	p := term.Read()
	fmt.Println("You entered: ", p)
	fmt.Print("Enter sample hidden text: ")
	p = term.ReadHidden()
	fmt.Println()
	fmt.Println("You entered while hidden: ", p)
}
