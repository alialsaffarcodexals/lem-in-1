package main

import (
	"fmt"
	"strings"
)

func Simulate(ants int, path []*Room) {
	if len(path) < 2 {
		return
	}

	// positions holds the index in the path for each ant.
	positions := make([]int, ants)
	finished := 0
	for finished < ants {
		moves := []string{}
		// build a list of ants at each position
		antsAt := make([][]int, len(path))
		for i, p := range positions {
			if p < len(path) {
				antsAt[p] = append(antsAt[p], i)
			}
		}
		edgesUsed := make(map[int]bool) // path index -> used

		// move ants from the end of the path backwards
		for pos := len(path) - 2; pos >= 0; pos-- {
			if len(antsAt[pos]) == 0 {
				continue
			}
			// at start several ants may wait, but only one can move per turn
			antIdx := antsAt[pos][0]
			next := pos + 1
			if next < len(path)-1 && len(antsAt[next]) > 0 {
				continue
			}
			if edgesUsed[pos] {
				continue
			}
			edgesUsed[pos] = true
			positions[antIdx] = next
			antsAt[pos] = antsAt[pos][1:]
			antsAt[next] = append(antsAt[next], antIdx)
			moves = append(moves, fmt.Sprintf("L%d-%s", antIdx+1, path[next].Name))
			if next == len(path)-1 {
				finished++
			}
		}
		if len(moves) > 0 {
			fmt.Fprintln(outWriter, strings.Join(moves, " "))
		}
	}
}
