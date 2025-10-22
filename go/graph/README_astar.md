# A* (A-Star) Pathfinding Algorithm in Go

A comprehensive implementation of the A* pathfinding algorithm, widely used in AI, game development, robotics, and navigation systems.

## Overview

A* (pronounced "A-star") is a best-first search algorithm that finds the optimal path between nodes in a graph. It was first published in 1968 by Peter Hart, Nils Nilsson, and Bertram Raphael. A* combines the benefits of Dijkstra's algorithm (guaranteed shortest path) with greedy best-first search (efficiency through heuristics).

## Algorithm Details

**Time Complexity**: O(b^d) where b is branching factor and d is depth
**Space Complexity**: O(b^d) for storing nodes in open and closed sets

The algorithm uses the evaluation function: **f(n) = g(n) + h(n)**
- **g(n)**: Actual cost from start to node n
- **h(n)**: Heuristic (estimated cost from node n to goal)
- **f(n)**: Estimated total cost of path through node n

## Features

- **Multiple Heuristics**: Manhattan, Euclidean, Chebyshev, and Diagonal distance
- **Grid-Based Pathfinding**: 2D grid navigation with obstacle support
- **Diagonal Movement**: Optional 8-directional movement with proper cost calculation
- **Priority Queue**: Efficient node selection using heap-based priority queue
- **Path Visualization**: ASCII grid visualization with path display
- **Comprehensive Analysis**: Detailed path statistics and performance metrics
- **Flexible Grid System**: Dynamic obstacle placement and removal
- **Extensive Testing**: 50+ test cases covering edge cases and complex scenarios

## Usage

### Basic Pathfinding

```go
package main

import (
    "fmt"
    "github.com/yourusername/DSA_Code/go/graph"
)

func main() {
    // Create a 10x10 grid
    grid := astar.NewGrid(10, 10)
    pathfinder := astar.NewAStar(grid)
    
    // Add some obstacles
    grid.AddObstacle(5, 1)
    grid.AddObstacle(5, 2)
    grid.AddObstacle(5, 3)
    grid.AddObstacle(5, 4)
    
    // Find path from top-left to bottom-right
    start := astar.Point{X: 0, Y: 0}
    goal := astar.Point{X: 9, Y: 9}
    
    path, err := pathfinder.FindPath(start, goal)
    if err != nil {
        fmt.Printf("No path found: %v\n", err)
        return
    }
    
    fmt.Printf("Path found with %d steps:\n", len(path))
    for i, point := range path {
        fmt.Printf("Step %d: %s\n", i, point.String())
    }
}
```

### Advanced Configuration

```go
// Enable diagonal movement
pathfinder.SetAllowDiagonal(true)

// Change heuristic function
pathfinder.SetHeuristic(astar.EuclideanDistance)

// Custom movement costs
pathfinder.StraightCost = 1.0
pathfinder.DiagonalCost = 1.4  // Slightly less than √2 for performance

// Get detailed results
result := pathfinder.FindPathWithDetails(start, goal)
fmt.Printf("Success: %t\n", result.Success)
fmt.Printf("Path cost: %.2f\n", result.PathCost)
fmt.Printf("Path length: %d\n", result.PathLength)
```

### Visualization

```go
// Visualize the grid with path
visualization := pathfinder.VisualizeGrid(path)
fmt.Println(visualization)
// Output:
// S·····
// ·█████
// ·····█
// ·····█
// ·····G
```

### Path Analysis

```go
stats := pathfinder.GetPathStatistics(path)
fmt.Printf("Total cost: %.2f\n", stats["cost"].(float64))
fmt.Printf("Straight moves: %d\n", stats["straight_moves"].(int))
fmt.Printf("Diagonal moves: %d\n", stats["diagonal_moves"].(int))
```

## Heuristic Functions

### Manhattan Distance (4-directional)
```go
pathfinder.SetHeuristic(astar.ManhattanDistance)
// Best for grid-based movement without diagonals
// h(n) = |x₁ - x₂| + |y₁ - y₂|
```

### Euclidean Distance (Any direction)
```go
pathfinder.SetHeuristic(astar.EuclideanDistance)
// Best for continuous movement
// h(n) = √((x₁ - x₂)² + (y₁ - y₂)²)
```

### Chebyshev Distance (8-directional, equal cost)
```go
pathfinder.SetHeuristic(astar.ChebyshevDistance)
// Best for 8-directional movement with equal diagonal cost
// h(n) = max(|x₁ - x₂|, |y₁ - y₂|)
```

### Diagonal Distance (8-directional, realistic cost)
```go
pathfinder.SetHeuristic(astar.DiagonalDistance)
// Best for 8-directional movement with realistic diagonal cost
// h(n) = max(dx, dy) + (√2 - 1) × min(dx, dy)
```

## Algorithm Workflow

1. **Initialize**: Add start node to open set
2. **Main Loop**: While open set is not empty:
   - Select node with lowest f(n) from open set
   - If goal reached, reconstruct and return path
   - Move current node from open to closed set
   - For each neighbor:
     - Skip if in closed set or not walkable
     - Calculate tentative g score
     - If new path to neighbor is shorter:
       - Update neighbor's parent and scores
       - Add to open set if not already present
3. **Result**: Return path or "no path found"

## Performance Characteristics

| Scenario | Time | Notes |
|----------|------|-------|
| Simple 20x20 grid | ~19μs | Direct path, no obstacles |
| Complex maze 50x50 | ~465μs | Multiple obstacles and detours |
| Diagonal movement | ~32μs | 8-directional pathfinding |

### Optimization Tips

1. **Choose appropriate heuristic**: Match heuristic to movement model
2. **Limit search space**: Use reasonable grid sizes
3. **Pre-process obstacles**: Static obstacles can be pre-computed
4. **Tie-breaking**: The implementation includes H-value tie-breaking for consistency

## Real-World Applications

### Game Development
- **NPCs pathfinding**: Character movement in games
- **Strategy games**: Unit movement planning
- **Racing games**: Optimal racing lines

### Robotics
- **Robot navigation**: Mobile robot path planning
- **Drone routing**: Autonomous flight planning
- **Warehouse automation**: Robotic picking systems

### Navigation Systems
- **GPS routing**: Road network navigation
- **Indoor navigation**: Building floor plans
- **Emergency evacuation**: Optimal exit planning

## Comparison with Other Algorithms

| Algorithm | Optimality | Speed | Memory |
|-----------|------------|-------|---------|
| A* | ✅ Optimal | 🟡 Fast | 🟡 Moderate |
| Dijkstra | ✅ Optimal | 🔴 Slow | 🟡 Moderate |
| Greedy Best-First | ❌ Not optimal | 🟢 Very fast | 🟢 Low |
| BFS | ✅ Optimal (unweighted) | 🔴 Slow | 🔴 High |
| DFS | ❌ Not optimal | 🟢 Fast | 🟢 Low |

## Educational Value

This implementation demonstrates:
- **Heuristic search**: Balancing optimality and efficiency
- **Priority queues**: Using heap data structures effectively
- **Graph algorithms**: Practical pathfinding techniques
- **Algorithm optimization**: Trade-offs in search strategies
- **Real-world problem solving**: From theory to implementation

## Advanced Features

### Dynamic Obstacles
```go
// Add obstacles during runtime
grid.AddObstacle(x, y)
grid.RemoveObstacle(x, y)

// Re-run pathfinding with updated grid
newPath, _ := pathfinder.FindPath(start, goal)
```

### Path Validation
```go
// Check if path is still valid after obstacles change
for _, point := range path {
    if !grid.IsWalkable(point.X, point.Y) {
        // Path is blocked, need to recalculate
        newPath, _ := pathfinder.FindPath(currentPos, goal)
        break
    }
}
```

## Testing

Run the comprehensive test suite:
```bash
go test -v astar.go astar_test.go
```

Run performance benchmarks:
```bash
go test -bench=. astar.go astar_test.go
```

## Limitations & Considerations

- **Grid-based**: Designed for discrete grid navigation
- **2D only**: Current implementation is 2D-focused
- **Memory usage**: Can be memory-intensive for large grids
- **Heuristic admissibility**: Heuristic must not overestimate for optimality

## Historical Context

A* was developed at Stanford Research Institute and has become one of the most widely used pathfinding algorithms. Its combination of optimality guarantees and practical efficiency makes it the gold standard for many pathfinding applications.

## References

- [A Formal Basis for the Heuristic Determination of Minimum Cost Paths](https://www.cs.auckland.ac.nz/courses/compsci709s2c/resources/Mike.d/astarNilsson.pdf) - Original 1968 paper
- [Amit's A* Pages](http://theory.stanford.edu/~amitp/GameProgramming/) - Comprehensive A* tutorial
- [A* Pathfinding for Beginners](http://www.policyalmanac.org/games/aStarTutorial.htm) - Practical implementation guide

## Author

Created for Hacktoberfest 2025 - Contributing to open source algorithm education and practical AI implementation.
