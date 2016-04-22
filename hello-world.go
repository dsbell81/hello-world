package main

import (
	"fmt"
	"github.com/dsbell81/hello-world/calc"
)

func main() {
	fmt.Printf("Hello, world.\n")

	var x, y int = 10, 5
	fmt.Println("x is:", x)
	fmt.Println("y is:", y)
	fmt.Println(calc.Add(x, y))
	fmt.Println(calc.Subtract(x, y))
}
