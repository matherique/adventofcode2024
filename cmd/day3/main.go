package main

import (
	"log/slog"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/matherique/adventofcode2024/pkg/assert"
)

const filename = "inputs/day3.txt"

func main() {
	fp, err := filepath.Abs(filename)
	assert.NotNil(err, "invalid file path", "filename", filename)
	f, err := os.ReadFile(fp)
	assert.NotNil(err, "failed to read the file", "filename", filename)

	part1(string(f))
}

func part1(memory string) int {
	var total int
	for i := 0; i < len(memory)-3; i++ {
		x := memory[i : i+4]
		if x == "mul(" {
			for j := 0; j <= 8; j++ {
				if i+4+j > len(memory) {
					break
				}

				p := memory[i : i+4+j]

				if x, y, ok := isValidMult(p); ok {
					total += x * y
					break
				}
			}
		}
	}

	slog.Info("the result of day 3 part 1", "total", total)

	return total
}

// "mul(1,2)"
func isValidMult(txt string) (int, int, bool) {
	if len(txt) < 8 {
		return 0, 0, false
	}

	if txt[:4] == "mul(" && txt[len(txt)-1:] == ")" {
		npart := txt[4 : len(txt)-1]
		slog.Info("number parts", "npart", npart)

		parts := strings.Split(npart, ",")

		if len(parts) != 2 {
			return 0, 0, false
		}

		x, err := strconv.Atoi(parts[0])
		if err != nil {
			return 0, 0, false
		}

		y, err := strconv.Atoi(parts[1])
		if err != nil {
			return 0, 0, false
		}

		return x, y, true
	}

	return 0, 0, false
}

func part2(memory string) int {
	return 0
}
