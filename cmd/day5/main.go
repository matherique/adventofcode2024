package main

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/matherique/adventofcode2024/pkg/assert"
	"github.com/matherique/adventofcode2024/pkg/utils"
)

const filename = "inputs/day5.txt"

func main() {
	fp, err := filepath.Abs(filename)
	assert.NotNil(err, "invalid file path", "filename", filename)
	f, err := os.ReadFile(fp)
	assert.NotNil(err, "failed to read the file", "filename", filename)

	rules, pages := parseInput(string(f))

	part1(rules, pages)
	part2(rules, pages)
}

func part1(rules map[string]bool, pages [][]string) int {
	var sum int

	for _, line := range pages {
		if isValidLine(rules, line) {
			m := len(line) / 2
			sum += utils.ParseNumber(line[m])
		}
	}

	slog.Info("the result of day 5 part 1", "sum", sum)

	return sum
}

func part2(rules map[string]bool, pages [][]string) int {
	var sum int

	invalids := make([][]string, 0)
	for _, line := range pages {
		if len(line) == 1 {
			continue
		}

		if !isValidLine(rules, line) {
			invalids = append(invalids, line)
		}
	}

	for _, invalid := range invalids {
		sort.Slice(invalid, func(i, j int) bool {
			return !rules[fmt.Sprintf("%s|%s", invalid[i], invalid[j])]
		})

		sum += utils.ParseNumber(invalid[len(invalid)/2])
	}

	slog.Info("the result of day 5 part 2", "sum", sum)

	return sum
}

func parseInput(input string) (map[string]bool, [][]string) {
	p := strings.Split(input, "\n\n")
	assert.True(len(p) == 2, "invalid input")
	rulesList := strings.Split(p[0], "\n")

	pageList := strings.Split(p[1], "\n")

	pages := make([][]string, len(pageList))
	for i, p := range pageList {
		pages[i] = make([]string, 0)

		for _, page := range strings.Split(p, ",") {
			pages[i] = append(pages[i], page)
		}
	}

	rules := make(map[string]bool)
	for _, rule := range rulesList {
		rules[rule] = true
	}

	return rules, pages
}

func isValidLine(rules map[string]bool, line []string) bool {
	valid := false
	for i := 0; i < len(line); i++ {
		for j := 0; j < len(line); j++ {
			if i == j {
				continue
			}

			k1 := line[i] + "|" + line[j]
			k2 := line[j] + "|" + line[i]

			if _, ok := rules[k1]; ok {
				valid = i < j
			}

			if _, ok := rules[k2]; ok {
				valid = j < i
			}

			if !valid {
				break
			}

		}
		if !valid {
			break
		}
	}

	return valid
}
