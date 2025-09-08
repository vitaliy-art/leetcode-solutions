func getNoZeroIntegers(n int) []int {
	isNoZero := func(num int) bool {
		numStr := fmt.Sprint(num)
		if strings.Contains(numStr, "0") {
			return false
		}
		return true
	}

	for i := 1; i < n; i++ {
		right := n - i
		if isNoZero(i) && isNoZero(right) {
			return []int{i, right}
		}
	}

	return nil
}
