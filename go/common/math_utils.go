package common

// GCD - Greatest Common Denominator calculated from Euclidian Algorithm.
func GCD(a, b int) int {
	if b == 0 {
		return a
	}

	return GCD(b, (a % b))
}

// LCM - Least Common Multiple can be calculated by taking a product of two numbers
// and dividing them by their GCD.
// Please note that this approach only works for two numbers.
func LCM(a, b int) int {
	return (a * b) / GCD(a, b)
}
