package differenceofsquares

func SquareOfSum(n int) int {
	sum := n * (n + 1) / 2
	return sum * sum
}

func SumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}


// func SquareOfSum(n int) int {
// 	result := 0
// 	for i := 1; i<=n; i++ {
// 		result += i
// 	}
// 	return result * result
// }
//
// func SumOfSquares(n int) int {
// 	result := 0
// 	for i := 1; i<=n; i++ {
// 		result += i*i
// 	}
// 	return result
// }
//
// func Difference(n int) int {
// 	return SquareOfSum(n) - SumOfSquares(n)
// }
