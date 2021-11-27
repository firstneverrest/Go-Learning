package main

import "fmt"

func main() {
	var name string = "Carlos"
	var age uint = 21
	// no white space & new line
	fmt.Print("Hello")
	fmt.Print("World")

	// need to add own space
	fmt.Print(name, " ", age) 
	fmt.Print("\n")

	// add white space & new line
	fmt.Println(name, age)

	// Printf
	fmt.Printf("I'm %v and I'm %v years old", name, age)
	fmt.Print("\n")
	fmt.Printf("name's type: %T age's type: %T", name, age)

	// Sprinf
	var message = fmt.Sprintf("my name is %v \n", name)
	fmt.Println("The saved string is:", message)
}