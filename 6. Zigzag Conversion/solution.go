func convert(s string, numRows int) string {
	runes := []rune(s)
	runesLen := len(runes)
	if runesLen == 1 || numRows == 1 {
		return s
	}

	subRows := make([][]rune, numRows)
	dblNumRows := numRows*2 - 2
	backNumRows := numRows - 2

main_loop:
	for i := 0; i < runesLen; i += dblNumRows {
		for r := range numRows {
			idx := i + r
			if idx >= runesLen {
				break main_loop
			}
			subRows[r] = append(subRows[r], runes[idx])
		}

		for r := range backNumRows {
			idx := i + numRows + r
			if idx >= runesLen {
				break main_loop
			}
			subRows[backNumRows-r] = append(subRows[numRows-r-2], runes[idx])
		}
	}

	var sb strings.Builder
	sb.Grow(len(s))

	for i := range subRows {
		sb.WriteString(string(subRows[i]))
	}

	return sb.String()
}
