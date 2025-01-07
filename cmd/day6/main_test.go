package main

import (
	"testing"
)

func init() {
	// logger := slog.New(slog.)NewTextHandler(io.Discard, nil))
	// slog.SetDefault(logger)
}

const input string = `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

func TestDay6Part1(t *testing.T) {
	result := part1(parseInput(input))

	if result != 41 {
		t.Fatalf("expected %d got %d", 41, result)
	}
}

func TestDay6Part2(t *testing.T) {
	result := part2(parseInput(input))

	if result != 123 {
		t.Fatalf("expected %d got %d", 123, result)
	}
}
