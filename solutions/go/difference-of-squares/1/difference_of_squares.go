package differenceofsquares

func SquareOfSum(n int) int {
	sum := (n * (n + 1)) / 2
	return sum * sum
}

func SumOfSquares(n int) int {
	return (n * (n + 1) * (2*n + 1)) / 6
}

func Difference(n int) int {
	if n < 1 {
		return 0
	}

	return SquareOfSum(n) - SumOfSquares(n)
}
