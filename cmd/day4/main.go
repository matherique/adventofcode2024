package main

import (
	"log/slog"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/matherique/adventofcode2024/pkg/assert"
)

const filename = "inputs/day4.txt"

func main() {
	fp, err := filepath.Abs(filename)
	assert.NotNil(err, "invalid file path", "filename", filename)
	f, err := os.ReadFile(fp)
	assert.NotNil(err, "failed to read the file", "filename", filename)

	input := parseInput(string(f))
	part1(input)
	part2(input)
}

func part1(input [][]string) int {
	var count int
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if input[i][j] != "X" {
				continue
			}

			// XMAS
			// 00 01 02 03
			// 10 11 12 13
			// 20 21 22 23
			// 30 31 32 33

			// i-1, j
			if i-3 >= 0 && input[i-1][j] == "M" && input[i-2][j] == "A" && input[i-3][j] == "S" {
				count++
			}

			// i+1, j
			if i+3 < len(input) && input[i+1][j] == "M" && input[i+2][j] == "A" && input[i+3][j] == "S" {
				count++
			}

			// i, j-1
			if j-3 >= 0 && input[i][j-1] == "M" && input[i][j-2] == "A" && input[i][j-3] == "S" {
				count++
			}

			// i, j+1
			if j+3 < len(input[i]) && input[i][j+1] == "M" && input[i][j+2] == "A" && input[i][j+3] == "S" {
				count++
			}

			// i+1, j+1
			if (j+3 < len(input[j]) && i+3 < len(input)) && input[i+1][j+1] == "M" && input[i+2][j+2] == "A" && input[i+3][j+3] == "S" {
				count++
			}

			// i-1, j-1
			if (j-3 >= 0 && i-3 >= 0) && input[i-1][j-1] == "M" && input[i-2][j-2] == "A" && input[i-3][j-3] == "S" {
				count++
			}

			// i-1, j+1
			if (j+3 < len(input[i]) && i-3 >= 0) && input[i-1][j+1] == "M" && input[i-2][j+2] == "A" && input[i-3][j+3] == "S" {
				count++
			}

			// i+1, j-1
			if (j-3 >= 0 && i+3 < len(input)) && input[i+1][j-1] == "M" && input[i+2][j-2] == "A" && input[i+3][j-3] == "S" {
				count++
			}
		}
	}

	slog.Info("the result of day 4 part 1 is", "count", count)
	return count
}

func word(input [][]string, i, j int) func(io, jo int) string {
	return func(io, jo int) string {
		return input[i+io][j+jo]
	}
}

func part2(input [][]string) int {
	var count int
	for i := 1; i < len(input)-1; i++ {
		for j := 1; j < len(input[i])-1; j++ {
			if input[i][j] != "A" {
				continue
			}
			// 00 01 02 03
			// 10 11 12 13
			// 20 21 22 23
			// 30 31 32 33
			//
			// M . S
			// . A .
			// M . S
			// M S S M

			// S . S
			// . A .
			// M . M
			// S S M M

			// S . M
			// . A .
			// S . M
			// S M M S

			// M . M
			// . A .
			// S . S
			// M M S S
			p1 := []string{"M", "S", "S", "M"}
			p2 := []string{"S", "S", "M", "M"}
			p3 := []string{"S", "M", "M", "S"}
			p4 := []string{"M", "M", "S", "S"}

			t := []string{input[i-1][j-1], input[i-1][j+1], input[i+1][j+1], input[i+1][j-1]}

			if reflect.DeepEqual(p1, t) ||
				reflect.DeepEqual(p2, t) ||
				reflect.DeepEqual(p3, t) ||
				reflect.DeepEqual(p4, t) {
				count++
			}

		}
	}

	slog.Info("the result of day 4 part 2 is", "count", count)
	return count

}

func parseInput(input string) [][]string {
	ws := make([][]string, 0)
	for i, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			continue
		}
		ws = append(ws, make([]string, len(line)))

		for j, l := range strings.Split(line, "") {
			ws[i][j] = l
		}
	}

	return ws
}
