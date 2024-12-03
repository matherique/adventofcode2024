package utils

import (
	"sort"
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
