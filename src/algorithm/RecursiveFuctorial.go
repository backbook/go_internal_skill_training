package algorithm

func Recursive(n int) int {
	if n <= 2 {
		return n
	}
	return n * Recursive(n-1)
}
