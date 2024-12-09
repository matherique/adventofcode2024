package main

import (
	"testing"
)

func init() {
	// logger := slog.New(slog.)NewTextHandler(io.Discard, nil))
	// slog.SetDefault(logger)
}

func TestDay5Part1(t *testing.T) {
	input := `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`
	result := part1(parseInput(input))

	if result != 143 {
		t.Fatalf("expected %d got %d", 143, result)
	}
}

func TestDay5Part2(t *testing.T) {
	input := `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`
	result := part2(parseInput(input))

	if result != 123 {
		t.Fatalf("expected %d got %d", 123, result)
	}
}
