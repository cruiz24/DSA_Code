# Dijkstra's Shortest Path Algorithm in Go

A comprehensive implementation of Dijkstra's algorithm for finding shortest paths in weighted graphs, with support for single-source and all-pairs shortest paths.

## Features

- **Single-Source Shortest Path**: Find distances to all vertices from source
- **Path Reconstruction**: Get actual shortest paths, not just distances  
- **All-Pairs Shortest Path**: Shortest paths between all vertex pairs
- **Graph Analysis**: Connectivity, density, negative weight detection
- **Priority Queue**: Efficient min-heap implementation for optimization
- **Comprehensive API**: Multiple convenience functions for common use cases
- **Full Testing**: Extensive test suite with benchmarks

## Time Complexity

| Operation | Time Complexity |
|-----------|-----------------|
| Single-Source | O((V + E) log V) |
| All-Pairs | O(V³) or O(V(V + E) log V) |
| Path Reconstruction | O(V) |
| Graph Analysis | O(V + E) |

## Space Complexity: O(V²) for all-pairs, O(V) for single-source

## Usage

```go
package main

import (
    "fmt"
    "graph"
)

func main() {
    // Create a weighted graph
    g := graph.NewGraph(5)
    
    // Add weighted edges
    g.AddEdge(0, 1, 4.0)
    g.AddEdge(0, 2, 2.0)
    g.AddEdge(1, 2, 1.0)
    g.AddEdge(1, 3, 5.0)
    g.AddEdge(2, 3, 8.0)
    g.AddEdge(2, 4, 10.0)
    g.AddEdge(3, 4, 2.0)
    
    // Find shortest paths from vertex 0
    result, err := g.Dijkstra(0)
    if err != nil {
        panic(err)
    }
    
    // Print all distances
    result.PrintDistances()
    
    // Get shortest path to specific vertex
    path, distance, err := result.GetShortestPath(4)
    if err == nil {
        fmt.Printf("Path to vertex 4: %v (distance: %.1f)\n", path, distance)
    }
    
    // Convenience function for single path
    path2, distance2, err := g.FindShortestPath(0, 4)
    if err == nil {
        fmt.Printf("Direct query: %v (%.1f)\n", path2, distance2)
    }
    
    // Get all shortest paths from source
    allPaths := result.GetAllPaths()
    for target, path := range allPaths {
        fmt.Printf("Path to %d: %v\n", target, path)
    }
    
    // All-pairs shortest paths
    distances, err := g.DijkstraAllPairs()
    if err == nil {
        fmt.Printf("Distance from 0 to 3: %.1f\n", distances[0][3])
    }
}
```

## Graph Construction

```go
// Create undirected graph
g := graph.NewGraph(4)
g.AddUndirectedEdge(0, 1, 5.0)
g.AddUndirectedEdge(1, 2, 3.0)
g.AddUndirectedEdge(2, 3, 2.0)

// Analyze graph properties
fmt.Printf("Connected: %t\n", g.IsConnected())
fmt.Printf("Density: %.3f\n", g.GetDensity())
fmt.Printf("Has negative weights: %t\n", g.HasNegativeWeight())

// Display graph structure  
g.PrintGraph()
```

## Algorithm Details

### Dijkstra's Process
1. Initialize distances to infinity (except source = 0)
2. Add all vertices to priority queue
3. Extract vertex with minimum distance
4. Update distances to all neighbors
5. Repeat until queue is empty

### Priority Queue Implementation
- Min-heap for efficient extraction of minimum distance vertex
- Supports decrease-key operation for distance updates
- O(log V) insertion and extraction operations

### Path Reconstruction
- Maintains parent pointers during algorithm execution
- Reconstructs path by following parent chain backwards
- Handles unreachable vertices gracefully

## Applications

- **Network Routing**: Internet packet routing protocols
- **GPS Navigation**: Shortest route finding in road networks  
- **Social Networks**: Shortest connection paths between users
- **Game AI**: Pathfinding for NPCs and characters
- **Flight Planning**: Cheapest flight route optimization
- **Supply Chain**: Optimal delivery route planning

## Comparison with Other Algorithms

| Algorithm | Graph Type | Time Complexity | Space | Negative Weights |
|-----------|------------|----------------|-------|------------------|
| Dijkstra | Weighted | O((V+E) log V) | O(V) | No |
| Bellman-Ford | Weighted | O(VE) | O(V) | Yes |
| Floyd-Warshall | All-pairs | O(V³) | O(V²) | Yes |
| A* | Weighted | O(E) | O(V) | No (with heuristic) |

## Advanced Features

### Graph Analysis
```go
// Check connectivity
connected := g.IsConnected()

// Calculate graph density
density := g.GetDensity()

// Detect negative weights (Dijkstra limitation)
hasNegative := g.HasNegativeWeight()
```

### Statistical Analysis
```go
// All-pairs shortest paths
distances, _ := g.DijkstraAllPairs()

// Analyze graph diameter, average path length, etc.
```

### Performance Optimization
- **Priority Queue**: Efficient heap-based implementation
- **Early Termination**: Stop when target is reached (single target)
- **Memory Layout**: Cache-friendly data structures
- **Batch Operations**: All-pairs computation optimization

## Edge Cases Handled

- **Unreachable Vertices**: Infinite distance reporting
- **Self-loops**: Proper handling of zero-weight self-edges
- **Disconnected Graphs**: Graceful handling of graph components  
- **Single Vertex**: Edge case for graphs with one vertex
- **No Edges**: Proper handling of vertex-only graphs

## Testing

Comprehensive test suite:

```bash
go test -v                    # All functionality tests
go test -bench=.             # Performance benchmarks  
go test -race                # Race condition detection
```

Test categories:
- Basic shortest path finding
- Path reconstruction accuracy
- Edge cases and error handling
- Graph property analysis
- Performance benchmarking

## Limitations

- **Negative Weights**: Cannot handle negative edge weights
- **Negative Cycles**: Cannot detect negative cycles
- **Memory Usage**: O(V²) for dense graphs
- **Dynamic Updates**: Not efficient for frequently changing graphs

## When to Use Dijkstra

**Good for:**
- Non-negative weighted graphs
- Single-source shortest paths
- Road networks, flight routes
- Network protocols
- Game pathfinding

**Consider alternatives for:**
- Negative weights → Bellman-Ford
- All-pairs → Floyd-Warshall  
- Unweighted graphs → BFS
- Heuristic available → A*

## Educational Value

This implementation demonstrates:
- Graph algorithms and data structures
- Priority queue applications  
- Greedy algorithm design
- Path reconstruction techniques
- Algorithm complexity analysis
- Real-world problem solving

## Extensions

Possible enhancements:
- Bidirectional Dijkstra
- Goal-directed search (A*)
- Parallel/concurrent implementation
- Dynamic graph updates
- Memory-efficient variants
