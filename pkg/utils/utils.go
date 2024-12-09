package utils

import (
	"sort"
	"strconv"

	"github.com/matherique/adventofcode2024/pkg/assert"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func IsSorted(arr []int) bool {
	return sort.IntsAreSorted(arr) || sort.IsSorted(sort.Reverse(sort.IntSlice(arr)))
}

func ParseNumber(s string) int {
	v, err := strconv.Atoi(s)
	assert.NotNil(err, "invalid number", "number", s)

	return v
}
