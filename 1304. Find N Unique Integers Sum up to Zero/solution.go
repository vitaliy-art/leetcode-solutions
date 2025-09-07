func sumZero(n int) []int {
	if n == 1 {
		return []int{0}
	}

	res := make([]int, n)
	sum := 0

	for idx := range n - 1 {
		cur := idx + 1
		res[idx] = cur
		sum += cur
	}

	res[n-1] = -sum
	return res
}
