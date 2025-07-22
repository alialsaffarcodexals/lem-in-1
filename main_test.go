package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestParseAndBFS(t *testing.T) {
	ants, rooms, _, err := ParseFile("testdata/example.txt")
	if err != nil {
		t.Fatalf("parse failed: %v", err)
	}
	if ants != 3 {
		t.Fatalf("expected 3 ants")
	}
	if rooms.Start == nil || rooms.End == nil {
		t.Fatalf("missing start or end")
	}
	path := BFS(rooms.Start, rooms.End)
	if len(path) == 0 {
		t.Fatalf("no path")
	}
}

func TestSimulateSimple(t *testing.T) {
	start := &Room{Name: "start"}
	mid := &Room{Name: "mid"}
	end := &Room{Name: "end"}
	path := []*Room{start, mid, end}

	var buf bytes.Buffer
	orig := outWriter
	outWriter = &buf
	defer func() { outWriter = orig }()

	Simulate(2, path)

	lines := strings.Split(strings.TrimSpace(buf.String()), "\n")
	expected := []string{
		"L1-mid",
		"L1-end L2-mid",
		"L2-end",
	}
	if len(lines) != len(expected) {
		t.Fatalf("expected %d lines, got %d", len(expected), len(lines))
	}
	for i, line := range lines {
		if line != expected[i] {
			t.Fatalf("line %d: expected %q got %q", i+1, expected[i], line)
		}
	}
}
