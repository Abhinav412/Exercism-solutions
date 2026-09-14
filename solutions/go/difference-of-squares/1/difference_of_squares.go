package differenceofsquares

func SquareOfSum(n int) int {
    result := n*(n+1)/2
    return result * result
}

func SumOfSquares(n int) int {
	result := n*(n+1)*(2*n+1)/6
    return result
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
