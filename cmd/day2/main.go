package main

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/matherique/adventofcode2024/pkg/assert"
	"github.com/matherique/adventofcode2024/pkg/utils"
)

const filename = "inputs/day2.txt"

func main() {
	fp, err := filepath.Abs(filename)
	assert.NotNil(err, "invalid file path", "filename", filename)
	f, err := os.ReadFile(fp)
	assert.NotNil(err, "failed to read the file", "filename", filename)

	lines := strings.Split(string(f), "\n")

	levels := make([][]int, len(lines)-1)

	for i, line := range lines[:len(lines)-1] {
		itens := strings.Split(line, " ")
		levels[i] = make([]int, 0, len(itens))
		fmt.Println(itens)
		for _, item := range itens {
			v, err := strconv.Atoi(item)
			assert.NotNil(err, "invalid number", "number", item)
			levels[i] = append(levels[i], v)
		}
	}

	part1(levels)
	part2(levels)
}

func part1(levels [][]int) int {
	var safe int
	for _, level := range levels {
		if isSafe(level) {
			safe += 1
		}
	}

	slog.Info("the result of day 2 part 1", "safe", safe)

	return safe
}

func part2(levels [][]int) int {
	possibilities := make([][][]int, len(levels))

	for i, level := range levels {
		possibilities[i] = make([][]int, 0)

		for j := 0; j < len(level); j++ {
			cp := make([]int, len(level))
			copy(cp, level)
			nlevel := append(cp[:j], cp[j+1:]...)
			possibilities[i] = append(possibilities[i], nlevel)
		}

	}

	safe := 0
	for _, plevels := range possibilities {
		for _, level := range plevels {
			if isSafe(level) {
				safe += 1
				break
			}
		}
	}

	slog.Info("the result of day 2 part 2", "safe", safe)

	return safe
}

func isSafe(level []int) bool {
	if !utils.IsSorted(level) {
		return false
	}

	safe := true
	for i := 0; i < len(level)-1; i++ {
		diff := utils.Abs(level[i] - level[i+1])
		if (diff > 3) || (diff == 0) {
			safe = false
			break
		}
	}

	return safe
}
