package main

import (
	"fmt"

	"golang.org/x/tour/tree"
	"practice.com/exercise"
)

func main() {
	t1 := tree.New(1)
	t2 := tree.New(1)
	fmt.Println(exercise.Same(t1, t2))
}

/*
OUTPUT
list of palindromes
if 0 and 1 are same then check all other indices are same or not, if diffrent then sent back all charcters on previous indices. this applies for all indices

01234
2

bab/aba
*/
