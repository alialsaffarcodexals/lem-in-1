package main

func BFS(start, end *Room) []*Room {
	if start == nil || end == nil {
		return nil
	}
	type node struct {
		r    *Room
		prev *node
	}
	queue := []*node{{r: start}}
	visited := map[*Room]bool{start: true}
	var endNode *node
	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]
		if n.r == end {
			endNode = n
			break
		}
		for _, nb := range n.r.Links {
			if !visited[nb] {
				visited[nb] = true
				queue = append(queue, &node{r: nb, prev: n})
			}
		}
	}
	if endNode == nil {
		return nil
	}
	// reconstruct
	var path []*Room
	for n := endNode; n != nil; n = n.prev {
		path = append([]*Room{n.r}, path...)
	}
	return path
}
