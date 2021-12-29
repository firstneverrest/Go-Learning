package main

// go standard library
import "fmt"

func main() {
	// declare variable with var
	var programmer1 string = "Oliver"

	// type is inferred
	var programmer2 = "John"

	// type is inferred
	programmer3 := "Amelia"

	// declare without initial value
	var programmer4 string

	fmt.Println("Hello", programmer1)
	fmt.Println("Hello", programmer2)
	fmt.Println("Hello", programmer3)
	fmt.Println("Hello", programmer4)

	var myAge int = 21
    var myWeight float32 = 55.7
    var isStudent bool = true

	fmt.Println(myAge)
	fmt.Println(myWeight)
	fmt.Println(isStudent)

}