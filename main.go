package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: go run . <file>")
		return
	}

	ants, rooms, _, err := ParseFile(os.Args[1])
	if err != nil {
		fmt.Println("ERROR: invalid data format")
		return
	}

	start, end := rooms.Start, rooms.End
	if start == nil || end == nil {
		fmt.Println("ERROR: invalid data format")
		return
	}

	path := BFS(start, end)
	if len(path) == 0 {
		fmt.Println("ERROR: no path")
		return
	}

	// Print input file content
	file, _ := os.Open(os.Args[1])
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	Simulate(ants, path)
}
