package main

import "fmt"

func main() {
  var score = map[string]string{"A": "Excellent", "B": "Good", "C": "Normal", "D": "Bad"}
  food := map[string]int{"Tom Yum": 50, "Omelette": 30, "Hotdog": 30,}

  fmt.Printf("score\t%v\n", score)
  fmt.Printf("food\t%v\n", food)
}
