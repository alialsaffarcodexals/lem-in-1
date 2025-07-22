package main

import "testing"

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
