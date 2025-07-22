package main

import (
	"fmt"
	"strings"
)

func Simulate(ants int, path []*Room) {
	positions := make([]int, ants)
	finished := 0
	for finished < ants {
		moves := []string{}
		for i := 0; i < ants; i++ {
			if positions[i] >= len(path) {
				continue
			}
			next := positions[i] + 1
			if next == len(path) {
				finished++
			}
			positions[i] = next
			moves = append(moves, fmt.Sprintf("L%d-%s", i+1, path[next-1].Name))
		}
		if len(moves) > 0 {
			fmt.Println(strings.Join(moves, " "))
		}
	}
}
