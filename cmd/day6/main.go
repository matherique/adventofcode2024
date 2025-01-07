package main

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/matherique/adventofcode2024/pkg/assert"
)

const filename = "inputs/day6.txt"

func main() {
	fp, err := filepath.Abs(filename)
	assert.NotNil(err, "invalid file path", "filename", filename)
	f, err := os.ReadFile(fp)
	assert.NotNil(err, "failed to read the file", "filename", filename)

	board, px, py := parseInput(string(f))

	part1(board, px, py)
	part2(board, px, py)

}

const (
	UP int = iota
	RIGHT
	DOWN
	LEFT
)

var DIR = []int{UP, RIGHT, DOWN, LEFT}

func part1(board [][]string, px int, py int) int {
	result := 1

	current := UP
	sizeY := len(board[0])
	sizeX := len(board)
	visited := make(map[string]bool)

	for {
		if (px-1 < 0 || px+1 == sizeX) || (py-1 < 0 || py+1 == sizeY) {
			slog.Info("finihsed")
			break
		}

		if current == UP {

			if board[px-1][py] == "#" {
				current = RIGHT
				continue
			}
			// fmt.Println("GO UP ")
			px--
		} else if current == RIGHT {
			if board[px][py+1] == "#" {
				current = DOWN
				continue
			}

			// fmt.Println("GO RIGHT ")
			py++

		} else if current == DOWN {
			if board[px+1][py] == "#" {
				current = LEFT
				continue
			}
			// fmt.Println("GO DOW ")
			px++
		} else if current == LEFT {
			if board[px][py-1] == "#" {
				current = UP
				continue
			}
			// fmt.Println("GO LEFT ")
			py--
		}

		k := fmt.Sprintf("%03d-03%d", px, py)
		if _, ok := visited[k]; !ok {
			result++
			visited[k] = true
		}
	}

	slog.Info("the result of day 5 part 1", "result", result)

	return result
}

func part2(board [][]string, px int, py int) int {
	var result int

	slog.Info("the result of day 5 part 2", "result", result)
	return result
}

func parseInput(input string) ([][]string, int, int) {
	lines := strings.Split(input, "\n")
	assert.True(len(lines) > 0, "invalid input", "len", len(lines))

	var x, y int
	board := make([][]string, 0, len(lines[0]))

	for i, line := range lines {
		board = append(board, make([]string, 0))
		for j, pos := range strings.Split(line, "") {
			if pos == "^" {
				x = i
				y = j
			}

			board[i] = append(board[i], pos)
		}
	}

	return board, x, y
}
