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
go run .\main.go
```

7. If you want to save the program as an executable program, type as follow:

```
go build .\main.go
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
