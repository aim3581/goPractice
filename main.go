package main

import (
	"fmt"

	"practice.com/calci"
)

func main() {
	fmt.Println("Hello World!")
	result := calci.Sum(12, 10)
	fmt.Println("Sum = ", result)
}
