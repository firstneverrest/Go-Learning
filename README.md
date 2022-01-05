# Go Learning

Go is an open source programming language that is fast, statically typed and compiled language. [Visit official website here](https://golang.org/). Go can be used for web development (server-side), developing network-based programs, cross-platform application, etc.

## Installation

1. [Download Go here](https://golang.org/doc/install).
2. Check whether Go is already installed in your computer

```
go version
```

3. Install Go extension in vscode
4. When Go extension window pop up, click Install All to make vscode ready for Go.
5. Write some code

```go
// main.go
package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
```

6. Run Go code

```
go run main.go
```

7. If you want to save the program as an executable program, type as follow:

```
go build main.go
```

8. Run .exe file

```
./main.exe
```

## Basic Syntax

- Package declaration - need to declare the package name of the file
- Import packages - use system packages and other packages in the file
- Functions - need to have main()
- Single-line comments (//) and Multiple-line comments (/\* \*/)
- Case-sensitive
- Double quote ("") only

## Variable

```go
package main

// go standard library
import "fmt"

var globalVariable string = "global"

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

    // declare multiple variables
    var a, b, c, d int = 1, 3, 5, 7
}
```

- string
- int, int8, int16, int32, int64, uint
- float
- bool

[More information about Go variables](https://pkg.go.dev/builtin#pkg-types).

## Constants

```go
package main

import "fmt"

const studentNumber uint = 40

const (
	age int = 2
	name = "Carlos"
)

func main() {
	fmt.Println(studentNumber)
	fmt.Println(age, name)
}
```

## Output Functions

```go
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

	// Sprintf
	var message = fmt.Sprintf("my name is %v \n", name)
	fmt.Println("The saved string is:", message)
}
```

- Print() - print without a new line, no whitespace is added between arguments.
- Println() - print with new line, add whitespace between arguments
- Printf() - format string based on the given formatting verb and then print.
  - `%v` - print value of the arguments
  - `%T` - print type of the arguments
  - `%q` - print value of the arguments with double quote wrapping
- Sprintf() - save formatted strings in variable

## Arrays

Arrays are used to store multiple values of the same type in a single variable.

```go
package main

import "fmt"

func main() {
	// define length
	var animals = [3]string{"dog", "cat", "elephant"}
	fmt.Println(animals, len(animals))

	// undefined length
	var foods = [...]string{"tom yum", "som tum"}
	foods[1] = "omelette"
	fmt.Println(foods)

	prices := [4]int{10, 30, 45, 60}
	fmt.Println(prices)
}
```

## Slices

Slice is more flexible than arrays, the length of the slice can grow and shrink.

- `len()` - returns the length of the slice
- `cap()` - returns the capacity of the slice (number of elements the slice can grow or shrink to)

```go
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
	mySlice3 := prices[:2]
	fmt.Printf("mySlice1 = %v\n", mySlice1)
	fmt.Printf("mySlice2 = %v\n", mySlice2)
	fmt.Printf("mySlice3 = %v\n", mySlice3)
}
```

```go
// trick: convert array to slice
grades := [4]string{"A", "B", "C", "D"}
getSlice(grades[:])
```

## Condition

### If-else

```go
package main
import ("fmt")

func main() {
  time := 22
  if time < 10 {
    fmt.Println("Good morning.")
  } else if time < 20 {
    fmt.Println("Good day.")
  } else {
    fmt.Println("Good evening.")
  }
}
```

### Switch Case

```go
package main
import ("fmt")

func main() {
  day := 4

  switch day {
  case 1:
    fmt.Println("Monday")
  case 2:
    fmt.Println("Tuesday")
  case 3:
    fmt.Println("Wednesday")
  case 4:
    fmt.Println("Thursday")
  case 5:
    fmt.Println("Friday")
  case 6:
    fmt.Println("Saturday")
  case 7:
    fmt.Println("Sunday")
  default:
    fmt.Println("Not a weekday")
  }
}
```

## Loops

```go
// nested loop
package main
import ("fmt")

func main() {
  adj := [2]string{"big", "tasty"}
  fruits := [3]string{"apple", "orange", "banana"}
  for i:=0; i < len(adj); i++ {
    for j:=0; j < len(fruits); j++ {
      fmt.Println(adj[i],fruits[j])
    }
  }
}
```

```go
// Infinite loop
package main
import ("fmt")

func main() {
  i := 0
  for {
	  fmt.Println(i)
	  i++
  }
}
```

```go
package main
import ("fmt")

func main() {
  fruits := [3]string{"apple", "orange", "banana"}
  for index, val := range fruits {
     fmt.Printf("%v\t%v\n", index, val)
  }

  // don't use index
  fruits := [3]string{"apple", "orange", "banana"}
  for _, val := range fruits {
     fmt.Printf("%v\t%v\n", val)
  }

  // use only index
   fruits := [3]string{"apple", "orange", "banana"}
   for index := range fruits {
     fmt.Printf("%v\t%v\n", index)
   }
}
```

## Function

Create function like `func main() {}`

```go
// function.go
package main

func getAnimals() []string {
	var animals = []string{"elephant", "weaver"}
	return animals
}
```

```go
// if arguments' type are the same, you can omit declaring type.
package main

func getAnimal(name, area string, age int) string {
	var animal = name + area;
	return animal
}

// shorthand
var add = func(a int, b int) int { return a + b}
```

```go
// main.go
package main

import "fmt"

func main() {
	fmt.Println(getAnimals())
}

```

When you would like to separate function into multiple files, you could not refer to function name and use it directly. It requires run all .go files that include function like below.

```
go run *.go
```

You can make your function be public by starting function name with uppercase.

## Struct

A struct is used to create a collection of members of different data types, into a single variable.

- Struct name and each member start with uppercase - public
- Struct name and each member start with lowercase - private

```go
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

	mira := Person{
		name: "Mira",
		age: 25,
		job: "Frontend Developer",
		salary: 30000,
	}

	fmt.Println("Name: ", carlos.name)
	fmt.Println("Age: ", carlos.age)
	fmt.Println("Job: ", carlos.job)
	fmt.Println("Salary: ", carlos.salary)

	fmt.Println("Name: ", mint.name)
	fmt.Println("Age: ", mint.age)
	fmt.Println("Job: ", mint.job)
	fmt.Println("Salary: ", mint.salary)

	fmt.Println("Name: ", mira.name)
}
```

## Method

Struct can has method inside it.

```go
type Person struct {
  name string
  age int
  job string
  salary int
}

func (p Person) getPersonInfo() {
  fmt.Printf("name: %s, age: %d, job: %s, salary: %d\n", p.name, p.age, p.job, p.salary)
}

func main() {
  mira.getPersonInfo()
}
```

## Map

A Map is used to store data values in key:value pairs.

```go
package main

import "fmt"

func main() {
  var score = map[string]string{"A": "Excellent", "B": "Good", "C": "Normal", "D": "Bad"}
  food := map[string]int{"Tom Yum": 50, "Omelette": 30, "Hotdog": 30,}

  fmt.Printf("score\t%v\n", score)
  fmt.Printf("food\t%v\n", food)

  // change a value
  food["Tom Yum"] = 60

  // insert a new value
  food["Brownie"] = 30

  // delete a value
  delete(food, "Omelette")

  // check whether the value exists in the map or not
  value, isExisted := food["Tuna Steak"]
  // isExisted is false if "Tuna Steak" doesn't exist
  // In case of "Tuna Steak" equals to "" or empty string
  // isExisted will return true because it exists

  // not allocate memory for map (map is nil)
  var m map[string]string

  // allocate memory for map
  var m map[string]string = map[string]string{}

  // or
  m := map[string]string{}

  // or
  m := make(map[string]string)
}

```

## Pointer

Pointer store variable address where it is stored in RAM.

```go
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
```

## Interface

Interfaces are named collections of method signatures.

```go
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
```

## Error

Error in Go is interface which means that if error is not equal to nil, it has an error.

```go
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
```

## Create Go Project

1. Create project with below command

```
go mod init <PROJECT_NAME>
```

1. Go created file called `go.mod` for you. This file store dependencies and meta data like package.json in JavaScript framework. Conventional module name is `github.com/<USERNAME>/<PACKAGE_PATH>` to evade duplicate package name.
2. when you create multiple packages and would like to use function in that separate package, You need to import with `import <ROOT_DIRECTORY>/<PACKAGE_NAME>`. `<ROOT_DIRECTORY>` depends on module name that defined in go.mod file.

## Go Resources

- [Go code guideline](https://github.com/uber-go/guide/blob/master/style.md#introduction)
- [Go codebangkok](https://github.com/codebangkok/golang)
- [Awesome Go](https://github.com/avelino/awesome-go)
- [Go tutorial from Net Ninja](https://www.youtube.com/playlist?list=PL4cUxeGkcC9gC88BEo9czgyS72A3doDeM)
