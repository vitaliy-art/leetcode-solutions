func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	maxLen := len(nums1) + len(nums2)
	isOdd := maxLen%2 == 0
	minLen := maxLen/2 + 1
	minIdx := minLen - 1

	nums3, rest, endIdx := insertSorted(nums1, nums2, 0)
	for endIdx < minIdx && len(rest) > 0 {
		nums3, rest, endIdx = insertSorted(nums3, rest, endIdx)
	}

	if !isOdd {
		return float64(nums3[minLen-1])
	}

	return float64(nums3[minLen-1]+nums3[minLen-2]) / float64(2)
}

func insertSorted(s1, s2 []int, startIdx int) (s3, rest []int, endIdx int) {
	if len(s1) == 0 {
		return s2, nil, 0
	}

	if len(s2) == 0 {
		return s1, nil, 0
	}

	if s1[0] >= s2[len(s2)-1] {
		return append(s2, s1...), nil, 0
	}

	if s2[0] >= s1[len(s1)-1] {
		return append(s1, s2...), nil, 0
	}

	left, right := s1, s2
	if s1[0] > s2[0] {
		left, right = s2, s1
	}

	insertIdx := searchMinIdx(left, right[0], startIdx)
	s3 = make([]int, insertIdx)
	copy(s3, left[:insertIdx])
	s3 = append(s3, right...)
	return s3, left[insertIdx:], insertIdx
}

func searchMinIdx(s []int, value, startIdx int) int {
	if len(s) == 1 {
		return 0
	}
	leftIdx, rightIdx := startIdx, len(s)-1
	for leftIdx < rightIdx {
		rightIdx /= 2
		if s[rightIdx] < value && s[rightIdx+1] > value {
			return rightIdx + 1
		}
		if s[rightIdx] < value {
			leftIdx = rightIdx + 1
		}
	}
	for ; leftIdx < len(s); leftIdx++ {
		if s[leftIdx] > value {
			return leftIdx
		}
	}

	return leftIdx
}
