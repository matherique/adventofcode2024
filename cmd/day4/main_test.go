package main

import (
	"testing"
)

func init() {
	// logger := slog.New(slog.)NewTextHandler(io.Discard, nil))
	// slog.SetDefault(logger)
}

func TestDay4Part1(t *testing.T) {
	input := `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`
	result := part1(parseInput(input))

	if result != 18 {
		t.Fatalf("expected %d got %d", 18, result)
	}
}

func TestDay4Part2(t *testing.T) {
	input := `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`
	result := part2(parseInput(input))

	if result != 9 {
		t.Fatalf("expected %d got %d", 9, result)
	}
}
