package main

import "fmt"

func main() {
	animal := "panda"
	fmt.Println(animal)

	var addr *string = &animal
	fmt.Printf("%T\n", addr)

	// change value
	*addr = "tiger"
	fmt.Printf("%s\n", *addr)
}