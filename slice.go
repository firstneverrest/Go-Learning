package main

import "fmt"

func main() {
	var animals = []string{"dog"}
	fmt.Println(animals, len(animals), cap(animals))

	var foods = []string{"tom yum", "som tum"}
	foods[1] = "omelette"
	foods = append(foods, "hotdog")
	fmt.Println(foods)

	// create slice from array
	prices := [4]int{10, 30, 45, 60}
	mySlice1 := prices[0:2]
	mySlice2 := prices[2:]
	mySlice3 := prices[:3]
	fmt.Printf("mySlice1 = %v\n", mySlice1)
	fmt.Printf("mySlice2 = %v\n", mySlice2)
	fmt.Printf("mySlice3 = %v\n", mySlice3)
}