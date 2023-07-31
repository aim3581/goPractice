// Pacakge calci impliments math operation fuctions
// This package supports variety of math operations started from addition of multiple numbers...
//
package calci

// Sum is used to calculate addition of passeed list of numbers
func Sum(args ...int) int {
	sum := 0
	for _, v := range args {
		sum += v
	}
	return sum
}
