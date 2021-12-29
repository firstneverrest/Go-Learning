package main

import (
	"errors"
	"fmt"
)

func ReadFile(name string) (string, error) {
	var err error = errors.New("File not found")
	return "", err
}

func main() {
	data, err := ReadFile("test.txt")
	// if it has an error
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Read file successfully: ", data)
}