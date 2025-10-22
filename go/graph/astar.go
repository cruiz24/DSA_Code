// time complexity: O(b^d) where b is branching factor and d is depth
// space complexity: O(b^d) for storing nodes in open and closed sets

package astar

import (
	"container/heap"
	"fmt"
	"math"
)

// Point represents a coordinate in 2D space
type Point struct {
	X, Y int
}

// String returns string representation of a point
func (p Point) String() string {
	return fmt.Sprintf("(%d,%d)", p.X, p.Y)
}

// Node represents a node in the A* search
type Node struct {
	Position Point   // Current position
	Parent   *Node   // Parent node for path reconstruction
	G        float64 // Cost from start to current node
	H        float64 // Heuristic cost from current node to goal
	F        float64 // Total cost (G + H)
}

// String returns string representation of a node
func (n *Node) String() string {
	return fmt.Sprintf("Node{pos:%s, g:%.1f, h:%.1f, f:%.1f}", 
		n.Position.String(), n.G, n.H, n.F)
}

// PriorityQueue implements a priority queue for A* nodes
type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// Lower F score has higher priority
	if pq[i].F == pq[j].F {
		return pq[i].H < pq[j].H // Tie-breaker: prefer lower H
	}
	return pq[i].F < pq[j].F
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Node))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

// HeuristicFunc defines the signature for heuristic functions
type HeuristicFunc func(a, b Point) float64

// ManhattanDistance calculates Manhattan distance between two points
func ManhattanDistance(a, b Point) float64 {
	return math.Abs(float64(a.X-b.X)) + math.Abs(float64(a.Y-b.Y))
}

// EuclideanDistance calculates Euclidean distance between two points
func EuclideanDistance(a, b Point) float64 {
	dx := float64(a.X - b.X)
	dy := float64(a.Y - b.Y)
	return math.Sqrt(dx*dx + dy*dy)
}

// ChebyshevDistance calculates Chebyshev distance (maximum of X and Y distances)
func ChebyshevDistance(a, b Point) float64 {
	dx := math.Abs(float64(a.X - b.X))
	dy := math.Abs(float64(a.Y - b.Y))
	return math.Max(dx, dy)
}

// DiagonalDistance calculates diagonal distance (octile distance)
func DiagonalDistance(a, b Point) float64 {
	dx := math.Abs(float64(a.X - b.X))
	dy := math.Abs(float64(a.Y - b.Y))
	return math.Max(dx, dy) + (math.Sqrt2-1)*math.Min(dx, dy)
}

// Grid represents a 2D grid for pathfinding
type Grid struct {
	Width     int       // Grid width
	Height    int       // Grid height
	Obstacles map[Point]bool // Set of obstacle positions
}

// NewGrid creates a new grid with specified dimensions
func NewGrid(width, height int) *Grid {
	return &Grid{
		Width:     width,
		Height:    height,
		Obstacles: make(map[Point]bool),
	}
}

// AddObstacle adds an obstacle at the specified position
func (g *Grid) AddObstacle(x, y int) {
	if g.IsValid(x, y) {
		g.Obstacles[Point{x, y}] = true
	}
}

// RemoveObstacle removes an obstacle at the specified position
func (g *Grid) RemoveObstacle(x, y int) {
	delete(g.Obstacles, Point{x, y})
}

// IsValid checks if a position is within grid bounds
func (g *Grid) IsValid(x, y int) bool {
	return x >= 0 && x < g.Width && y >= 0 && y < g.Height
}

// IsWalkable checks if a position is walkable (not an obstacle)
func (g *Grid) IsWalkable(x, y int) bool {
	if !g.IsValid(x, y) {
		return false
	}
	return !g.Obstacles[Point{x, y}]
}

// GetNeighbors returns valid neighboring positions
func (g *Grid) GetNeighbors(pos Point, allowDiagonal bool) []Point {
	var neighbors []Point
	
	// 4-directional movement (up, down, left, right)
	directions := []Point{
		{0, -1}, {0, 1}, {-1, 0}, {1, 0},
	}
	
	// Add diagonal directions if allowed
	if allowDiagonal {
		diagonals := []Point{
			{-1, -1}, {-1, 1}, {1, -1}, {1, 1},
		}
		directions = append(directions, diagonals...)
	}
	
	for _, dir := range directions {
		newX := pos.X + dir.X
		newY := pos.Y + dir.Y
		
		if g.IsWalkable(newX, newY) {
			// For diagonal movement, check if both adjacent cells are walkable
			if allowDiagonal && dir.X != 0 && dir.Y != 0 {
				if !g.IsWalkable(pos.X+dir.X, pos.Y) || !g.IsWalkable(pos.X, pos.Y+dir.Y) {
					continue // Skip this diagonal if adjacent cells are blocked
				}
			}
			neighbors = append(neighbors, Point{newX, newY})
		}
	}
	
	return neighbors
}

// AStar represents the A* pathfinding algorithm
type AStar struct {
	Grid           *Grid         // The grid to search on
	Heuristic      HeuristicFunc // Heuristic function to use
	AllowDiagonal  bool          // Whether diagonal movement is allowed
	DiagonalCost   float64       // Cost of diagonal movement (typically √2 ≈ 1.414)
	StraightCost   float64       // Cost of straight movement (typically 1.0)
}

// NewAStar creates a new A* pathfinder
func NewAStar(grid *Grid) *AStar {
	return &AStar{
		Grid:          grid,
		Heuristic:     ManhattanDistance,
		AllowDiagonal: false,
		DiagonalCost:  math.Sqrt2,
		StraightCost:  1.0,
	}
}

// SetHeuristic sets the heuristic function
func (a *AStar) SetHeuristic(heuristic HeuristicFunc) {
	a.Heuristic = heuristic
}

// SetAllowDiagonal enables or disables diagonal movement
func (a *AStar) SetAllowDiagonal(allow bool) {
	a.AllowDiagonal = allow
	// Use appropriate heuristic for diagonal movement
	if allow {
		a.Heuristic = DiagonalDistance
	} else {
		a.Heuristic = ManhattanDistance
	}
}

// CalculateMovementCost calculates the cost of moving from one point to another
func (a *AStar) CalculateMovementCost(from, to Point) float64 {
	dx := math.Abs(float64(from.X - to.X))
	dy := math.Abs(float64(from.Y - to.Y))
	
	// Diagonal movement
	if dx == 1 && dy == 1 {
		return a.DiagonalCost
	}
	// Straight movement
	return a.StraightCost
}

// FindPath finds the shortest path from start to goal using A* algorithm
func (a *AStar) FindPath(start, goal Point) ([]Point, error) {
	// Validate start and goal positions
	if !a.Grid.IsWalkable(start.X, start.Y) {
		return nil, fmt.Errorf("start position %s is not walkable", start.String())
	}
	if !a.Grid.IsWalkable(goal.X, goal.Y) {
		return nil, fmt.Errorf("goal position %s is not walkable", goal.String())
	}
	
	// If start equals goal, return path with just the start point
	if start == goal {
		return []Point{start}, nil
	}
	
	// Initialize open and closed sets
	openSet := &PriorityQueue{}
	heap.Init(openSet)
	closedSet := make(map[Point]bool)
	allNodes := make(map[Point]*Node)
	
	// Create start node
	startNode := &Node{
		Position: start,
		Parent:   nil,
		G:        0,
		H:        a.Heuristic(start, goal),
	}
	startNode.F = startNode.G + startNode.H
	
	heap.Push(openSet, startNode)
	allNodes[start] = startNode
	
	// A* main loop
	for openSet.Len() > 0 {
		// Get node with lowest F score
		current := heap.Pop(openSet).(*Node)
		
		// Check if we reached the goal
		if current.Position == goal {
			return a.reconstructPath(current), nil
		}
		
		// Move current node from open to closed set
		closedSet[current.Position] = true
		
		// Examine neighbors
		neighbors := a.Grid.GetNeighbors(current.Position, a.AllowDiagonal)
		for _, neighborPos := range neighbors {
			// Skip if neighbor is in closed set
			if closedSet[neighborPos] {
				continue
			}
			
			// Calculate tentative G score
			movementCost := a.CalculateMovementCost(current.Position, neighborPos)
			tentativeG := current.G + movementCost
			
			// Check if this neighbor is already in allNodes
			neighbor, exists := allNodes[neighborPos]
			if !exists {
				// Create new neighbor node
				neighbor = &Node{
					Position: neighborPos,
					Parent:   current,
					G:        tentativeG,
					H:        a.Heuristic(neighborPos, goal),
				}
				neighbor.F = neighbor.G + neighbor.H
				allNodes[neighborPos] = neighbor
				heap.Push(openSet, neighbor)
			} else if tentativeG < neighbor.G {
				// Found a better path to this neighbor
				neighbor.Parent = current
				neighbor.G = tentativeG
				neighbor.F = neighbor.G + neighbor.H
				
				// If neighbor is not in open set, add it
				if !a.isInOpenSet(openSet, neighbor) {
					heap.Push(openSet, neighbor)
				}
			}
		}
	}
	
	// No path found
	return nil, fmt.Errorf("no path found from %s to %s", start.String(), goal.String())
}

// reconstructPath reconstructs the path from goal to start by following parent pointers
func (a *AStar) reconstructPath(goalNode *Node) []Point {
	var path []Point
	current := goalNode
	
	for current != nil {
		path = append([]Point{current.Position}, path...)
		current = current.Parent
	}
	
	return path
}

// isInOpenSet checks if a node is in the open set
func (a *AStar) isInOpenSet(openSet *PriorityQueue, node *Node) bool {
	for _, n := range *openSet {
		if n.Position == node.Position {
			return true
		}
	}
	return false
}

// PathfindingResult contains the result of pathfinding operation
type PathfindingResult struct {
	Path          []Point // The found path (empty if no path)
	PathLength    int     // Length of the path
	PathCost      float64 // Total cost of the path
	NodesExplored int     // Number of nodes explored during search
	Success       bool    // Whether a path was found
	Error         error   // Error if any occurred
}

// FindPathWithDetails finds a path and returns detailed results
func (a *AStar) FindPathWithDetails(start, goal Point) PathfindingResult {
	path, err := a.FindPath(start, goal)
	
	result := PathfindingResult{
		Path:    path,
		Success: err == nil,
		Error:   err,
	}
	
	if result.Success {
		result.PathLength = len(path)
		result.PathCost = a.CalculatePathCost(path)
	}
	
	return result
}

// CalculatePathCost calculates the total cost of a path
func (a *AStar) CalculatePathCost(path []Point) float64 {
	if len(path) <= 1 {
		return 0
	}
	
	totalCost := 0.0
	for i := 1; i < len(path); i++ {
		cost := a.CalculateMovementCost(path[i-1], path[i])
		totalCost += cost
	}
	
	return totalCost
}

// VisualizeGrid creates a string representation of the grid with path
func (a *AStar) VisualizeGrid(path []Point) string {
	var result string
	pathSet := make(map[Point]bool)
	
	// Convert path to set for O(1) lookup
	for _, point := range path {
		pathSet[point] = true
	}
	
	for y := 0; y < a.Grid.Height; y++ {
		for x := 0; x < a.Grid.Width; x++ {
			point := Point{x, y}
			if a.Grid.Obstacles[point] {
				result += "█" // Obstacle
			} else if pathSet[point] {
				if len(path) > 0 && point == path[0] {
					result += "S" // Start
				} else if len(path) > 0 && point == path[len(path)-1] {
					result += "G" // Goal
				} else {
					result += "·" // Path
				}
			} else {
				result += " " // Empty space
			}
		}
		result += "\n"
	}
	
	return result
}

// GetPathStatistics returns detailed statistics about a path
func (a *AStar) GetPathStatistics(path []Point) map[string]interface{} {
	if len(path) == 0 {
		return map[string]interface{}{
			"length": 0,
			"cost":   0.0,
		}
	}
	
	straightMoves := 0
	diagonalMoves := 0
	totalCost := 0.0
	
	for i := 1; i < len(path); i++ {
		cost := a.CalculateMovementCost(path[i-1], path[i])
		totalCost += cost
		
		if cost == a.StraightCost {
			straightMoves++
		} else if cost == a.DiagonalCost {
			diagonalMoves++
		}
	}
	
	return map[string]interface{}{
		"length":         len(path),
		"cost":           totalCost,
		"straight_moves": straightMoves,
		"diagonal_moves": diagonalMoves,
		"start":          path[0],
		"goal":           path[len(path)-1],
	}
}
