package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Room struct {
	Name  string
	X, Y  int
	Links []*Room
}

type Rooms struct {
	All   map[string]*Room
	Start *Room
	End   *Room
}

func ParseFile(path string) (int, *Rooms, [][2]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, nil, nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	rooms := &Rooms{All: make(map[string]*Room)}
	var links [][2]string
	lineNum := 0
	ants := 0
	expectRoom := ""

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		lineNum++
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, "#") {
			if line == "##start" {
				expectRoom = "start"
			} else if line == "##end" {
				expectRoom = "end"
			}
			continue
		}
		// first non-comment line is number of ants
		if lineNum == 1 {
			ants, err = strconv.Atoi(line)
			if err != nil {
				return 0, nil, nil, fmt.Errorf("invalid ants")
			}
			continue
		}

		fields := strings.Fields(line)
		if len(fields) == 3 {
			x, err1 := strconv.Atoi(fields[1])
			y, err2 := strconv.Atoi(fields[2])
			if err1 != nil || err2 != nil {
				return 0, nil, nil, fmt.Errorf("invalid room coords")
			}
			room := &Room{Name: fields[0], X: x, Y: y}
			rooms.All[room.Name] = room
			if expectRoom == "start" {
				rooms.Start = room
				expectRoom = ""
			} else if expectRoom == "end" {
				rooms.End = room
				expectRoom = ""
			}
			continue
		}
		if strings.Contains(line, "-") {
			parts := strings.Split(line, "-")
			if len(parts) != 2 {
				return 0, nil, nil, fmt.Errorf("invalid link")
			}
			links = append(links, [2]string{parts[0], parts[1]})
			continue
		}
		return 0, nil, nil, fmt.Errorf("invalid line")
	}
	// create adjacency
	for _, l := range links {
		a := rooms.All[l[0]]
		b := rooms.All[l[1]]
		if a != nil && b != nil {
			a.Links = append(a.Links, b)
			b.Links = append(b.Links, a)
		}
	}
	return ants, rooms, links, nil
}
