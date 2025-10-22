package graph

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

func TestNewGraph(t *testing.T) {
	g := NewGraph(5)
	if g.Vertices != 5 {
		t.Errorf("Expected 5 vertices, got %d", g.Vertices)
	}
	if len(g.Edges) != 5 {
		t.Errorf("Expected 5 edge lists, got %d", len(g.Edges))
	}
}

func TestAddEdge(t *testing.T) {
	g := NewGraph(3)
	g.AddEdge(0, 1, 5.0)
	g.AddEdge(1, 2, 3.0)
	
	if len(g.Edges[0]) != 1 {
		t.Errorf("Expected 1 edge from vertex 0, got %d", len(g.Edges[0]))
	}
	
	edge := g.Edges[0][0]
	if edge.To != 1 || edge.Weight != 5.0 {
		t.Errorf("Expected edge to vertex 1 with weight 5.0, got to %d with weight %.1f", edge.To, edge.Weight)
	}
	
	// Test invalid edges
	g.AddEdge(-1, 1, 1.0) // Should be ignored
	g.AddEdge(0, 5, 1.0)  // Should be ignored
	
	if len(g.Edges[0]) != 1 {
		t.Error("Invalid edges should be ignored")
	}
}

func TestAddUndirectedEdge(t *testing.T) {
	g := NewGraph(3)
	g.AddUndirectedEdge(0, 1, 5.0)
	
	if len(g.Edges[0]) != 1 || len(g.Edges[1]) != 1 {
		t.Error("Undirected edge should create edges in both directions")
	}
	
	if g.Edges[0][0].To != 1 || g.Edges[1][0].To != 0 {
		t.Error("Undirected edge endpoints incorrect")
	}
}

func TestDijkstraSimple(t *testing.T) {
	// Create simple graph: 0 → 1 → 2
	//                     ↓    ↑
	//                     3 ←←←
	g := NewGraph(4)
	g.AddEdge(0, 1, 1.0)
	g.AddEdge(1, 2, 2.0)
	g.AddEdge(0, 3, 4.0)
	g.AddEdge(3, 2, 1.0)
	
	result, err := g.Dijkstra(0)
	if err != nil {
		t.Fatalf("Dijkstra failed: %v", err)
	}
	
	expectedDistances := []float64{0, 1, 3, 4}
	for i, expected := range expectedDistances {
		if math.Abs(result.Distances[i]-expected) > 1e-9 {
			t.Errorf("Distance to vertex %d: expected %.1f, got %.1f", i, expected, result.Distances[i])
		}
	}
}

func TestDijkstraInvalidSource(t *testing.T) {
	g := NewGraph(3)
	
	_, err := g.Dijkstra(-1)
	if err == nil {
		t.Error("Expected error for negative source vertex")
	}
	
	_, err = g.Dijkstra(5)
	if err == nil {
		t.Error("Expected error for source vertex out of bounds")
	}
}

func TestGetShortestPath(t *testing.T) {
	g := NewGraph(4)
	g.AddEdge(0, 1, 1.0)
	g.AddEdge(1, 2, 2.0)
	g.AddEdge(0, 3, 4.0)
	g.AddEdge(3, 2, 1.0)
	
	result, _ := g.Dijkstra(0)
	
	// Path from 0 to 2 should be 0 → 1 → 2
	path, distance, err := result.GetShortestPath(2)
	if err != nil {
		t.Fatalf("GetShortestPath failed: %v", err)
	}
	
	expectedPath := []int{0, 1, 2}
	if !reflect.DeepEqual(path, expectedPath) {
		t.Errorf("Expected path %v, got %v", expectedPath, path)
	}
	
	if math.Abs(distance-3.0) > 1e-9 {
		t.Errorf("Expected distance 3.0, got %.1f", distance)
	}
}

func TestGetShortestPathNoPath(t *testing.T) {
	g := NewGraph(3)
	g.AddEdge(0, 1, 1.0)
	// Vertex 2 is unreachable
	
	result, _ := g.Dijkstra(0)
	
	_, _, err := result.GetShortestPath(2)
	if err == nil {
		t.Error("Expected error for unreachable vertex")
	}
}

func TestGetShortestPathInvalidTarget(t *testing.T) {
	g := NewGraph(3)
	result, _ := g.Dijkstra(0)
	
	_, _, err := result.GetShortestPath(-1)
	if err == nil {
		t.Error("Expected error for invalid target vertex")
	}
	
	_, _, err = result.GetShortestPath(5)
	if err == nil {
		t.Error("Expected error for target vertex out of bounds")
	}
}

func TestGetAllPaths(t *testing.T) {
	g := NewGraph(4)
	g.AddEdge(0, 1, 1.0)
	g.AddEdge(1, 2, 2.0)
	g.AddEdge(0, 3, 4.0)
	
	result, _ := g.Dijkstra(0)
	paths := result.GetAllPaths()
	
	if len(paths) != 4 { // Including path to itself
		t.Errorf("Expected 4 paths, got %d", len(paths))
	}
	
	// Check path to vertex 2
	expectedPath := []int{0, 1, 2}
	if !reflect.DeepEqual(paths[2], expectedPath) {
		t.Errorf("Expected path to vertex 2: %v, got %v", expectedPath, paths[2])
	}
}

func TestDijkstraAllPairs(t *testing.T) {
	g := NewGraph(3)
	g.AddUndirectedEdge(0, 1, 1.0)
	g.AddUndirectedEdge(1, 2, 2.0)
	g.AddUndirectedEdge(0, 2, 4.0)
	
	distances, err := g.DijkstraAllPairs()
	if err != nil {
		t.Fatalf("DijkstraAllPairs failed: %v", err)
	}
	
	// Distance from 0 to 2 should be 3 (via vertex 1)
	if math.Abs(distances[0][2]-3.0) > 1e-9 {
		t.Errorf("Expected distance from 0 to 2: 3.0, got %.1f", distances[0][2])
	}
	
	// Distance from 2 to 0 should also be 3 (undirected graph)
	if math.Abs(distances[2][0]-3.0) > 1e-9 {
		t.Errorf("Expected distance from 2 to 0: 3.0, got %.1f", distances[2][0])
	}
}

func TestFindShortestPath(t *testing.T) {
	g := NewGraph(3)
	g.AddEdge(0, 1, 1.0)
	g.AddEdge(1, 2, 2.0)
	
	path, distance, err := g.FindShortestPath(0, 2)
	if err != nil {
		t.Fatalf("FindShortestPath failed: %v", err)
	}
	
	expectedPath := []int{0, 1, 2}
	if !reflect.DeepEqual(path, expectedPath) {
		t.Errorf("Expected path %v, got %v", expectedPath, path)
	}
	
	if math.Abs(distance-3.0) > 1e-9 {
		t.Errorf("Expected distance 3.0, got %.1f", distance)
	}
}

func TestHasNegativeWeight(t *testing.T) {
	g := NewGraph(3)
	g.AddEdge(0, 1, 1.0)
	g.AddEdge(1, 2, 2.0)
	
	if g.HasNegativeWeight() {
		t.Error("Graph should not have negative weights")
	}
	
	g.AddEdge(2, 0, -1.0)
	
	if !g.HasNegativeWeight() {
		t.Error("Graph should have negative weights")
	}
}

func TestIsConnected(t *testing.T) {
	// Connected graph
	g1 := NewGraph(3)
	g1.AddUndirectedEdge(0, 1, 1.0)
	g1.AddUndirectedEdge(1, 2, 1.0)
	
	if !g1.IsConnected() {
		t.Error("Graph should be connected")
	}
	
	// Disconnected graph
	g2 := NewGraph(3)
	g2.AddEdge(0, 1, 1.0)
	// Vertex 2 is isolated
	
	if g2.IsConnected() {
		t.Error("Graph should not be connected")
	}
	
	// Single vertex
	g3 := NewGraph(1)
	if !g3.IsConnected() {
		t.Error("Single vertex graph should be connected")
	}
	
	// Empty graph
	g4 := NewGraph(0)
	if !g4.IsConnected() {
		t.Error("Empty graph should be connected")
	}
}

func TestGetDensity(t *testing.T) {
	g := NewGraph(3)
	
	// Initially no edges
	if g.GetDensity() != 0 {
		t.Errorf("Expected density 0 for empty graph, got %.3f", g.GetDensity())
	}
	
	// Add all possible directed edges
	g.AddEdge(0, 1, 1.0)
	g.AddEdge(0, 2, 1.0)
	g.AddEdge(1, 0, 1.0)
	g.AddEdge(1, 2, 1.0)
	g.AddEdge(2, 0, 1.0)
	g.AddEdge(2, 1, 1.0)
	
	expectedDensity := 1.0 // 6 edges out of 6 possible (3 * 2)
	if math.Abs(g.GetDensity()-expectedDensity) > 1e-9 {
		t.Errorf("Expected density %.1f, got %.3f", expectedDensity, g.GetDensity())
	}
	
	// Test single vertex (edge case)
	g1 := NewGraph(1)
	if g1.GetDensity() != 0 {
		t.Errorf("Single vertex graph should have density 0, got %.3f", g1.GetDensity())
	}
}

func TestDijkstraWithPath(t *testing.T) {
	g := NewGraph(3)
	g.AddEdge(0, 1, 1.0)
	g.AddEdge(1, 2, 2.0)
	
	distance, path, err := g.DijkstraWithPath(0, 2)
	if err != nil {
		t.Fatalf("DijkstraWithPath failed: %v", err)
	}
	
	expectedDistance := 3.0
	expectedPath := []int{0, 1, 2}
	
	if math.Abs(distance-expectedDistance) > 1e-9 {
		t.Errorf("Expected distance %.1f, got %.1f", expectedDistance, distance)
	}
	
	if !reflect.DeepEqual(path, expectedPath) {
		t.Errorf("Expected path %v, got %v", expectedPath, path)
	}
}

func TestKShortestPaths(t *testing.T) {
	g := NewGraph(3)
	g.AddEdge(0, 1, 1.0)
	g.AddEdge(1, 2, 2.0)
	
	paths, distances, err := g.KShortestPaths(0, 2, 1)
	if err != nil {
		t.Fatalf("KShortestPaths failed: %v", err)
	}
	
	if len(paths) != 1 || len(distances) != 1 {
		t.Errorf("Expected 1 path and 1 distance, got %d paths and %d distances", len(paths), len(distances))
	}
	
	expectedPath := []int{0, 1, 2}
	expectedDistance := 3.0
	
	if !reflect.DeepEqual(paths[0], expectedPath) {
		t.Errorf("Expected path %v, got %v", expectedPath, paths[0])
	}
	
	if math.Abs(distances[0]-expectedDistance) > 1e-9 {
		t.Errorf("Expected distance %.1f, got %.1f", expectedDistance, distances[0])
	}
	
	// Test invalid k
	_, _, err = g.KShortestPaths(0, 2, 0)
	if err == nil {
		t.Error("Expected error for k = 0")
	}
}

// Benchmark tests
func BenchmarkDijkstra(b *testing.B) {
	sizes := []int{100, 500, 1000}
	
	for _, size := range sizes {
		g := NewGraph(size)
		
		// Create a dense graph
		for i := 0; i < size; i++ {
			for j := 0; j < size; j++ {
				if i != j {
					g.AddEdge(i, j, float64(i+j+1))
				}
			}
		}
		
		b.Run(fmt.Sprintf("Size_%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				g.Dijkstra(0)
			}
		})
	}
}

func BenchmarkGetShortestPath(b *testing.B) {
	g := NewGraph(1000)
	
	// Create a linear graph
	for i := 0; i < 999; i++ {
		g.AddEdge(i, i+1, 1.0)
	}
	
	result, _ := g.Dijkstra(0)
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result.GetShortestPath(999)
	}
}

func BenchmarkDijkstraAllPairs(b *testing.B) {
	g := NewGraph(50)
	
	// Create a complete graph
	for i := 0; i < 50; i++ {
		for j := 0; j < 50; j++ {
			if i != j {
				g.AddEdge(i, j, float64(i+j+1))
			}
		}
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		g.DijkstraAllPairs()
	}
}
