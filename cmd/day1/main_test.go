package main

import (
	"bytes"
	"io"
	"log/slog"
	"testing"
)

func init() {
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	slog.SetDefault(logger)
}

func TestDay1Part1(t *testing.T) {

	lines := []byte(`3   4
4   3
2   5
1   3
3   9
3   3`)

	result := part1(bytes.Split(lines, []byte{10}))

	if result != 11 {
		t.Fatalf("expected %d got %d", 11, result)
	}
}
