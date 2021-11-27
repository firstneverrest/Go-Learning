package main

import "fmt"

func main() {
	var animals = [3]string{"dog", "cat", "elephant"}
	fmt.Println(animals, len(animals))

	var foods = [...]string{"tom yum", "som tum"}
	foods[1] = "omelette"
	fmt.Println(foods)

	prices := [4]int{10, 30, 45, 60}
	fmt.Println(prices)
}