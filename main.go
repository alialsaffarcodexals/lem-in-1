package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var outWriter io.Writer = os.Stdout

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(outWriter, "usage: go run . <file>")
		return
	}

	ants, rooms, _, err := ParseFile(os.Args[1])
	if err != nil {
		fmt.Fprintln(outWriter, "ERROR: invalid data format")
		return
	}

	start, end := rooms.Start, rooms.End
	if start == nil || end == nil {
		fmt.Fprintln(outWriter, "ERROR: invalid data format")
		return
	}

	path := BFS(start, end)
	if len(path) == 0 {
		fmt.Fprintln(outWriter, "ERROR: no path")
		return
	}

	// Print input file content
	file, _ := os.Open(os.Args[1])
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Fprintln(outWriter, scanner.Text())
	}

	Simulate(ants, path)
}
