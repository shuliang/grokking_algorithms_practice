package main

import (
	"fmt"
)

var graph = map[string][]string{
	"you":    []string{"alice", "bob", "claire"},
	"bob":    []string{"anuj", "peggy"},
	"alice":  []string{"you", "peggy", "alice"},
	"claire": []string{"jonny", "thom", "bob"},
	"anuj":   []string{},
	"peggy":  []string{"you"},
	"thom":   []string{},
	"jonny":  []string{"peggy", "peggy"},
	"jim":    []string{"thom"},
}

func personIsSeller(name string) bool {
	return len(name) > 0 && name[len(name)-1] == 'm'
}

// bfs breadth-first search `seller` in graph, begin at `start`
// return `seller` name if found, otherwise ""
func bfs(graph map[string][]string, start string) string {
	deque := []string{start}
	visited := map[string]bool{start: true}
	for len(deque) > 0 {
		head := deque[0]
		if list, ok := graph[head]; ok {
			for _, person := range list {
				if _, exist := visited[person]; !exist {
					deque = append(deque, person)
					visited[person] = true
				}
			}
		}
		if len(deque) > 1 {
			deque = deque[1:]
		} else {
			return ""
		}

		if personIsSeller(deque[0]) {
			return deque[0]
		}
	}
	return ""
}

func main() {
	r := bfs(graph, "you")
	// r = bfs(graph, "jonny")
	if r != "" {
		fmt.Printf("found seller: %v\n", r)
	} else {
		fmt.Println("There is no seller.")
	}
}
