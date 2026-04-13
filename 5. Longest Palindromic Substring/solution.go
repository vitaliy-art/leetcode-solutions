func longestPalindrome(s string) string {
	var res []rune
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		for j := i + len(res); j < len(runes); j++ {
			subRunes := runes[i : j+1]
			if len(subRunes) > len(res) && isPalindrome(subRunes) {
				res = subRunes
			}
		}
	}
	return string(res)
}

func isPalindrome(runes []rune) bool {
	runesLen := len(runes)
	lastIdx := runesLen - 1
	midIdx := runesLen / 2
	for i := 0; i <= midIdx; i++ {
		if runes[i] != runes[lastIdx-i] {
			return false
		}
	}
	return true
}
