package main

import (
	"fmt"
	"strings"

	"practice.com/exercise"
)

func main() {
	s := "abbcccbbbcaaccbababcbcabca"
	s = exercise.LongestPalindrome(strings.ToLower(s))
	fmt.Println(s)
}

/*
OUTPUT
list of palindromes
if 0 and 1 are same then check all other indices are same or not, if diffrent then sent back all charcters on previous indices. this applies for all indices

01234
2

bab/aba
*/
