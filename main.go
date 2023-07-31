package main

import (
	"fmt"

	"practice.com/calci"
)

func main() {
	fmt.Println("Hello World!")
	result := calci.Add(12, 10)
	fmt.Println("Sum = ", result)
}
