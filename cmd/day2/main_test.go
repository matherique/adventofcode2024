package main

import (
	"testing"
)

func init() {
	// logger := slog.New(slog.)NewTextHandler(io.Discard, nil))
	// slog.SetDefault(logger)
}

func TestDay2Part1(t *testing.T) {
	levels := [][]int{
		{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9},
	}

	result := part1(levels)

	if result != 2 {
		t.Fatalf("expected %d got %d", 2, result)
	}
}

func TestDay2Part2(t *testing.T) {
	levels := [][]int{
		{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9},
	}

	result := part2(levels)

	if result != 4 {
		t.Fatalf("expected %d got %d", 4, result)
	}
}
