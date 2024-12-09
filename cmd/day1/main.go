package main

import (
	"bytes"
	"log/slog"
	"os"
	"path/filepath"
	"sort"
	"strconv"

	"github.com/matherique/adventofcode2024/pkg/assert"
	"github.com/matherique/adventofcode2024/pkg/utils"
)

const filename = "inputs/day1.txt"

func main() {
	fp, err := filepath.Abs(filename)
	assert.NotNil(err, "invalid file path", "filename", filename)
	f, err := os.ReadFile(fp)
	assert.NotNil(err, "failed to read the file", "filename", filename)

	lines := bytes.Split(f, []byte{10})

	rights, lefts := make([]int, len(lines)-1), make([]int, len(lines)-1)
	for i, line := range lines[:len(lines)-1] {
		sides := bytes.Split(line, []byte{32, 32, 32})
		assert.True(len(sides) == 2, "invalid line", "len", len(sides))

		left, err := strconv.Atoi(string(sides[0]))
		assert.NotNil(err, "invalid number", "number", string(sides[0]))
		lefts[i] = left

		right, err := strconv.Atoi(string(sides[1]))
		assert.NotNil(err, "invalid number", "number", string(sides[1]))
		rights[i] = right
	}
	part1(rights, lefts)
	part2(rights, lefts)

}

func part1(rights, lefts []int) int {
	sort.Ints(rights)
	sort.Ints(lefts)

	var sum int
	for i := 0; i < len(rights); i += 1 {
		diff := utils.Abs(rights[i] - lefts[i])
		sum += diff
	}

	slog.Info("the result of part 1", "sum", sum)
	return sum
}

func part2(rights, lefts []int) int {
	count := make(map[int]int)
	for i := 0; i < len(rights); i += 1 {
		for j := 0; j < len(lefts); j += 1 {
			if rights[j] == lefts[i] {
				count[lefts[i]] += 1
			}
		}
	}

	var sum int
	for k, v := range count {
		sum += k * v
	}

	slog.Info("the result of part 2", "sum", sum)
	return sum
}
