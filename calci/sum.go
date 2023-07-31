//Pacakge calci is used for math operations
package calci

// Sum is used to calculate addition of passeed list of numbers
func Sum(args ...int) int {
	sum := 0
	for _, v := range args {
		sum += v
	}
	return sum
}
