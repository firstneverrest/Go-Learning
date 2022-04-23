package main

import (
	"errors"
	"fmt"
)

func main() {
	result, err := divide(100.0, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Result:", result)
	}
}

func divide(x, y float32) (float32, error) {
	var result float32

	if y == 0 {
		return result, errors.New("divided by zero")
	}

	result = x / y
	return result, nil
}
