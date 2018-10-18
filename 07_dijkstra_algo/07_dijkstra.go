package main

import (
	"fmt"
	"math"
)

type graph map[string]map[string]int

const endNode = "NONE"

var g = graph{
	"start": map[string]int{"a": 6, "b": 2},
	"a":     map[string]int{"fin": 1},
	"b":     map[string]int{"a": 3, "fin": 5},
	"fin":   map[string]int{},
}

/*
dijkstra return the shortest path of `graph` if found, otherwise return nil.
The sequnence of `path` slice represents the shortest route, if path is
`{"start", "b", "a", "fin"}`, the shortest path is `start => b => a => fin`.
*/
func dijkstra(graph graph, start string) (path []string) {
	visited := map[string]bool{}
	parents := calcInitialParents(graph, start)
	costs := calcInitialCosts(graph, start)
	if parents == nil || costs == nil {
		return nil
	}
	node := findLowestCostNode(costs, visited)
	for node != endNode {
		cost := costs[node]
		neighbours := graph[node]
		for n, d := range neighbours {
			newCost := cost + d
			if costs[n] > newCost {
				costs[n] = newCost
				parents[n] = node
			}
		}
		visited[node] = true
		node = findLowestCostNode(costs, visited)
	}
	path = calcPathInParents(parents, start)
	return path
}

// calculate inital costs of graph with start
func calcInitialCosts(g graph, start string) map[string]int {
	if _, ok := g[start]; !ok {
		return nil
	}
	costs := g[start]
	costs[start] = 0
	for k := range g {
		if _, ok := costs[k]; !ok && k != start {
			costs[k] = math.MaxInt32
		}
	}
	return costs
}

// calculate start point and "fin" point's parents of graph at begining
func calcInitialParents(g graph, start string) map[string]string {
	if _, ok := g[start]; !ok {
		return nil
	}
	parents := map[string]string{}
	for k := range g[start] {
		parents[k] = start
	}
	if _, ok := parents["fin"]; !ok {
		parents["fin"] = endNode
	}
	return parents
}

func calcPathInParents(parents map[string]string, start string) []string {
	if parents["fin"] == endNode {
		return nil
	}
	currentPoint := "fin"
	path := []string{}
	for currentPoint != start {
		for k, v := range parents {
			if k == currentPoint {
				path = append(path, currentPoint)
				currentPoint = v
			}
		}
	}
	if len(path) == 1 {
		return []string{start, path[0]}
	}
	reversed := []string{start}
	for i := len(path) - 1; i >= 0; i-- {
		reversed = append(reversed, path[i])
	}
	return reversed
}

func findLowestCostNode(costs map[string]int, visited map[string]bool) string {
	node := endNode
	low := math.MaxInt32
	for k, v := range costs {
		if _, ok := visited[k]; !ok && v < low {
			low = v
			node = k
		}
	}
	return node
}

func main() {
	if path := dijkstra(g, "start"); path == nil {
		fmt.Println("no path.")
	} else {
		fmt.Println(path)
	}
}
