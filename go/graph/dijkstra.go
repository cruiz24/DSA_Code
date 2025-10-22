package graph

import (
	"container/heap"
	"fmt"
	"math"
)

// Edge represents a weighted edge in the graph
type Edge struct {
	To     int
	Weight float64
}

// Graph represents a weighted directed graph
type Graph struct {
	Vertices int
	Edges    [][]Edge
}

// NewGraph creates a new graph with the specified number of vertices
func NewGraph(vertices int) *Graph {
	return &Graph{
		Vertices: vertices,
		Edges:    make([][]Edge, vertices),
	}
}

// AddEdge adds a weighted edge to the graph
func (g *Graph) AddEdge(from, to int, weight float64) {
	if from < 0 || from >= g.Vertices || to < 0 || to >= g.Vertices {
		return
	}
	g.Edges[from] = append(g.Edges[from], Edge{To: to, Weight: weight})
}

// AddUndirectedEdge adds an undirected weighted edge
func (g *Graph) AddUndirectedEdge(u, v int, weight float64) {
	g.AddEdge(u, v, weight)
	g.AddEdge(v, u, weight)
}

// PriorityQueueItem represents an item in the priority queue
type PriorityQueueItem struct {
	Vertex   int
	Distance float64
	Index    int
}

// PriorityQueue implements a min-heap for Dijkstra's algorithm
type PriorityQueue []*PriorityQueueItem

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Distance < pq[j].Distance
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*PriorityQueueItem)
	item.Index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.Index = -1
	*pq = old[0 : n-1]
	return item
}

// DijkstraResult contains the results of Dijkstra's algorithm
type DijkstraResult struct {
	Distances []float64
	Previous  []int
	Source    int
}

// Dijkstra implements Dijkstra's shortest path algorithm
func (g *Graph) Dijkstra(source int) (*DijkstraResult, error) {
	if source < 0 || source >= g.Vertices {
		return nil, fmt.Errorf("invalid source vertex: %d", source)
	}

	// Initialize distances and previous vertices
	distances := make([]float64, g.Vertices)
	previous := make([]int, g.Vertices)
	visited := make([]bool, g.Vertices)

	for i := range distances {
		distances[i] = math.Inf(1)
		previous[i] = -1
	}
	distances[source] = 0

	// Initialize priority queue
	pq := &PriorityQueue{}
	heap.Init(pq)

	// Add all vertices to priority queue
	items := make([]*PriorityQueueItem, g.Vertices)
	for i := 0; i < g.Vertices; i++ {
		items[i] = &PriorityQueueItem{
			Vertex:   i,
			Distance: distances[i],
		}
		heap.Push(pq, items[i])
	}

	for pq.Len() > 0 {
		current := heap.Pop(pq).(*PriorityQueueItem)
		u := current.Vertex

		if visited[u] {
			continue
		}
		visited[u] = true

		// Process all neighbors
		for _, edge := range g.Edges[u] {
			v := edge.To
			if visited[v] {
				continue
			}

			newDistance := distances[u] + edge.Weight
			if newDistance < distances[v] {
				distances[v] = newDistance
				previous[v] = u

				// Add updated distance to priority queue
				heap.Push(pq, &PriorityQueueItem{
					Vertex:   v,
					Distance: newDistance,
				})
			}
		}
	}

	return &DijkstraResult{
		Distances: distances,
		Previous:  previous,
		Source:    source,
	}, nil
}

// GetShortestPath returns the shortest path from source to target
func (dr *DijkstraResult) GetShortestPath(target int) ([]int, float64, error) {
	if target < 0 || target >= len(dr.Distances) {
		return nil, 0, fmt.Errorf("invalid target vertex: %d", target)
	}

	if math.IsInf(dr.Distances[target], 1) {
		return nil, math.Inf(1), fmt.Errorf("no path from %d to %d", dr.Source, target)
	}

	// Reconstruct path
	path := make([]int, 0)
	current := target

	for current != -1 {
		path = append([]int{current}, path...)
		current = dr.Previous[current]
	}

	return path, dr.Distances[target], nil
}

// GetAllPaths returns shortest paths to all reachable vertices
func (dr *DijkstraResult) GetAllPaths() map[int][]int {
	paths := make(map[int][]int)

	for target := 0; target < len(dr.Distances); target++ {
		if !math.IsInf(dr.Distances[target], 1) {
			path, _, err := dr.GetShortestPath(target)
			if err == nil {
				paths[target] = path
			}
		}
	}

	return paths
}

// DijkstraAllPairs finds shortest paths between all pairs of vertices
func (g *Graph) DijkstraAllPairs() ([][]float64, error) {
	distances := make([][]float64, g.Vertices)
	for i := range distances {
		distances[i] = make([]float64, g.Vertices)
	}

	// Run Dijkstra from each vertex
	for source := 0; source < g.Vertices; source++ {
		result, err := g.Dijkstra(source)
		if err != nil {
			return nil, err
		}
		distances[source] = result.Distances
	}

	return distances, nil
}

// FindShortestPath is a convenience function to find shortest path between two vertices
func (g *Graph) FindShortestPath(source, target int) ([]int, float64, error) {
	result, err := g.Dijkstra(source)
	if err != nil {
		return nil, 0, err
	}

	return result.GetShortestPath(target)
}

// HasNegativeWeight checks if the graph has any negative weight edges
func (g *Graph) HasNegativeWeight() bool {
	for i := 0; i < g.Vertices; i++ {
		for _, edge := range g.Edges[i] {
			if edge.Weight < 0 {
				return true
			}
		}
	}
	return false
}

// IsConnected checks if the graph is connected (for undirected graphs)
func (g *Graph) IsConnected() bool {
	if g.Vertices <= 1 {
		return true
	}

	result, err := g.Dijkstra(0)
	if err != nil {
		return false
	}

	// Check if all vertices are reachable from vertex 0
	for i := 1; i < g.Vertices; i++ {
		if math.IsInf(result.Distances[i], 1) {
			return false
		}
	}

	return true
}

// GetDensity calculates the density of the graph
func (g *Graph) GetDensity() float64 {
	if g.Vertices <= 1 {
		return 0
	}

	edgeCount := 0
	for i := 0; i < g.Vertices; i++ {
		edgeCount += len(g.Edges[i])
	}

	maxEdges := g.Vertices * (g.Vertices - 1)
	return float64(edgeCount) / float64(maxEdges)
}

// PrintGraph prints the graph structure
func (g *Graph) PrintGraph() {
	fmt.Printf("Graph with %d vertices:\n", g.Vertices)
	for i := 0; i < g.Vertices; i++ {
		fmt.Printf("Vertex %d: ", i)
		for j, edge := range g.Edges[i] {
			if j > 0 {
				fmt.Print(", ")
			}
			fmt.Printf("→ %d (%.1f)", edge.To, edge.Weight)
		}
		fmt.Println()
	}
}

// PrintDistances prints the distance table from Dijkstra result
func (dr *DijkstraResult) PrintDistances() {
	fmt.Printf("Shortest distances from vertex %d:\n", dr.Source)
	for i, dist := range dr.Distances {
		if math.IsInf(dist, 1) {
			fmt.Printf("To vertex %d: ∞\n", i)
		} else {
			fmt.Printf("To vertex %d: %.1f\n", i, dist)
		}
	}
}

// DijkstraWithPath combines pathfinding and returns both distance and path
func (g *Graph) DijkstraWithPath(source, target int) (float64, []int, error) {
	result, err := g.Dijkstra(source)
	if err != nil {
		return 0, nil, err
	}

	path, distance, err := result.GetShortestPath(target)
	return distance, path, err
}

// KShortestPaths finds k shortest paths from source to target (simplified version)
func (g *Graph) KShortestPaths(source, target, k int) ([][]int, []float64, error) {
	if k <= 0 {
		return nil, nil, fmt.Errorf("k must be positive")
	}

	// For simplicity, this returns just the shortest path
	// A full k-shortest paths implementation would be much more complex
	distance, path, err := g.DijkstraWithPath(source, target)
	if err != nil {
		return nil, nil, err
	}

	paths := [][]int{path}
	distances := []float64{distance}

	return paths, distances, nil
}
