func lengthOfLongestSubstring(s string) int {
	set := map[rune]struct{}{}
	longest := 0
	current := 0
	subIdx := -1
	runes := []rune(s)

	for idx := range runes {
		r := runes[idx]
		if _, ok := set[r]; ok {
			if current > longest {
				longest = current
			}

			for {
				subIdx++
				delete(set, runes[subIdx])
				current--
				if runes[subIdx] == r {
					break
				}
			}
		}

		current++
		set[r] = struct{}{}
	}

	if current > longest {
		longest = current
	}

	return longest
}
