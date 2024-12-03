package main

import (
	"testing"
)

func init() {
	// logger := slog.New(slog.)NewTextHandler(io.Discard, nil))
	// slog.SetDefault(logger)
}

func TestDay3Part1(t *testing.T) {
	memory := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	result := part1(memory)

	if result != 161 {
		t.Fatalf("expected %d got %d", 161, result)
	}
}

func TestDay3Part2(t *testing.T) {
	memory := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
	result := part2(memory)

	if result != 48 {
		t.Fatalf("expected %d got %d", 48, result)
	}
}
