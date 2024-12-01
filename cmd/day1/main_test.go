package main

import (
	"testing"
)

func init() {
	logger := slog.New(slog.NewT=extHandler(io.Discard, nil))
	slog.SetDefault(logger)
}

func TestDay1Part1(t *testing.T) {
	left := []int{3, 4, 2, 1, 3, 3}
	right := []int{4, 3, 5, 3, 9, 3}

	result := part1(right, left)

	if result != 11 {
		t.Fatalf("expected %d got %d", 11, result)
	}
}

func TestDay1Part2(t *testing.T) {
	left := []int{3, 4, 2, 1, 3, 3}
	right := []int{4, 3, 5, 3, 9, 3}

	result := part2(right, left)

	if result != 31 {
		t.Fatalf("expected %d got %d", 31, result)
	}
}
