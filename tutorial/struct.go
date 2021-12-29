package main

import "fmt"

type Person struct {
  name string
  age int
  job string
  salary int
}

func main() {
	var carlos Person
	var mint Person

	carlos.name = "Carlos"
	carlos.age = 21
	carlos.job = "Frontend Developer"
	carlos.salary = 25000

	mint.name = "Mint"
	mint.age = 23
	mint.job = "Business Analyst"
	mint.salary = 27000

	fmt.Println("Name: ", carlos.name)
	fmt.Println("Age: ", carlos.age)
	fmt.Println("Job: ", carlos.job)
	fmt.Println("Salary: ", carlos.salary)

	fmt.Println("Name: ", mint.name)
	fmt.Println("Age: ", mint.age)
	fmt.Println("Job: ", mint.job)
	fmt.Println("Salary: ", mint.salary)
}