package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
)

func reverse(x int) int {
	abs := int(math.Abs(float64(x)))
	runes := []rune(fmt.Sprint(abs))
	slices.Reverse(runes)

	reversed, err := strconv.ParseInt(string(runes), 10, 32)
	if err != nil {
		return 0
	}

	if x < 0 {
		return -int(reversed)
	}

	return int(reversed)
}
