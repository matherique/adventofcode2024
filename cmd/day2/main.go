package main

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"sort"
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
		levels[i] = make([]int, len(itens))
		fmt.Println(itens)
		for j, item := range itens {
			v, err := strconv.Atoi(item)
			assert.NotNil(err, "invalid number", "number", item)
			levels[i][j] = v
		}
	}

	part1(levels)
}

func part1(levels [][]int) int {
	var safe int
	for _, level := range levels {
		if sort.IntsAreSorted(level) || sort.IsSorted(sort.Reverse(sort.IntSlice(level))) {
			isSafe := true
			slog.Info("is sorted", "level", level)
			for i := 0; i < len(level)-1; i++ {
				diff := utils.Abs(level[i] - level[i+1])
				slog.Info("diff", "diff", diff, fmt.Sprintf("%d", i), level[i], fmt.Sprintf("%d", i+1), level[i+1])
				if diff > 3 || diff == 0 {
					slog.Info("not safe", "level", level)
					isSafe = false
					break
				}
			}

			if isSafe {
				slog.Info("is safe", "level", level)
				safe += 1
			}
		}
	}

	slog.Info("the result of day 2 part 1", "safe", safe)

	return safe
}

func part2() int {
	return 0
}
