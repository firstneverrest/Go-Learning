package main

import "fmt"

type Animal interface {
	Eat(food string)
}

type Panda struct {
	Name string
}

func (p Panda) Eat(food string) {
	fmt.Printf("%s eats %s\n", p.Name, food)
}

func (p Panda) Sleep() {
	fmt.Printf("%s sleeps\n", p.Name)
}

func CheckAnimal(a Animal) {
	a.Eat("bamboo")
}

func main() {
	panda := Panda{Name: "panda"}

	CheckAnimal(panda)
}