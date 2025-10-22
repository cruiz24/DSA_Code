package astar

import (
	"container/heap"
	"math"
	"strings"
	"testing"
)

func TestPoint(t *testing.T) {
	p := Point{X: 5, Y: 10}
	expected := "(5,10)"
	if p.String() != expected {
		t.Errorf("Expected %s, got %s", expected, p.String())
	}
}

func TestNode(t *testing.T) {
	node := &Node{
		Position: Point{X: 1, Y: 2},
		G:        5.0,
		H:        3.0,
		F:        8.0,
	}
	
	expected := "Node{pos:(1,2), g:5.0, h:3.0, f:8.0}"
	if node.String() != expected {
		t.Errorf("Expected %s, got %s", expected, node.String())
	}
}

func TestPriorityQueue(t *testing.T) {
	pq := &PriorityQueue{}
	heap.Init(pq)
	
	// Test empty queue
	if pq.Len() != 0 {
		t.Error("Expected empty queue length to be 0")
	}
	
	// Add nodes with different F scores
	node1 := &Node{F: 10.0, H: 5.0}
	node2 := &Node{F: 5.0, H: 3.0}
	node3 := &Node{F: 15.0, H: 8.0}
	node4 := &Node{F: 5.0, H: 2.0} // Same F as node2, but lower H
	
	heap.Push(pq, node1)
	heap.Push(pq, node2)
	heap.Push(pq, node3)
	heap.Push(pq, node4)
	
	if pq.Len() != 4 {
		t.Errorf("Expected queue length 4, got %d", pq.Len())
	}
	
	// Should pop in order of F score (lowest first), with H as tiebreaker
	popped1 := heap.Pop(pq).(*Node)
	if popped1.F != 5.0 || popped1.H != 2.0 {
		t.Errorf("Expected first node to have F=5.0, H=2.0, got F=%.1f, H=%.1f", popped1.F, popped1.H)
	}
	
	popped2 := heap.Pop(pq).(*Node)
	if popped2.F != 5.0 || popped2.H != 3.0 {
		t.Errorf("Expected second node to have F=5.0, H=3.0, got F=%.1f, H=%.1f", popped2.F, popped2.H)
	}
}

func TestHeuristicFunctions(t *testing.T) {
	a := Point{0, 0}
	b := Point{3, 4}
	
	// Test Manhattan distance: |3-0| + |4-0| = 7
	manhattan := ManhattanDistance(a, b)
	if manhattan != 7.0 {
		t.Errorf("Expected Manhattan distance 7.0, got %.1f", manhattan)
	}
	
	// Test Euclidean distance: sqrt(3^2 + 4^2) = 5
	euclidean := EuclideanDistance(a, b)
	if math.Abs(euclidean-5.0) > 0.001 {
		t.Errorf("Expected Euclidean distance 5.0, got %.3f", euclidean)
	}
	
	// Test Chebyshev distance: max(3, 4) = 4
	chebyshev := ChebyshevDistance(a, b)
	if chebyshev != 4.0 {
		t.Errorf("Expected Chebyshev distance 4.0, got %.1f", chebyshev)
	}
	
	// Test Diagonal distance
	diagonal := DiagonalDistance(a, b)
	expected := 4.0 + (math.Sqrt2-1)*3.0 // max(3,4) + (√2-1)*min(3,4)
	if math.Abs(diagonal-expected) > 0.001 {
		t.Errorf("Expected diagonal distance %.3f, got %.3f", expected, diagonal)
	}
}

func TestGrid(t *testing.T) {
	grid := NewGrid(5, 5)
	
	// Test grid creation
	if grid.Width != 5 || grid.Height != 5 {
		t.Errorf("Expected grid size 5x5, got %dx%d", grid.Width, grid.Height)
	}
	
	// Test valid positions
	if !grid.IsValid(0, 0) || !grid.IsValid(4, 4) {
		t.Error("Corner positions should be valid")
	}
	
	if grid.IsValid(-1, 0) || grid.IsValid(5, 0) || grid.IsValid(0, -1) || grid.IsValid(0, 5) {
		t.Error("Out-of-bounds positions should be invalid")
	}
	
	// Test walkable positions
	if !grid.IsWalkable(2, 2) {
		t.Error("Empty position should be walkable")
	}
	
	// Add obstacle and test
	grid.AddObstacle(2, 2)
	if grid.IsWalkable(2, 2) {
		t.Error("Obstacle position should not be walkable")
	}
	
	// Remove obstacle and test
	grid.RemoveObstacle(2, 2)
	if !grid.IsWalkable(2, 2) {
		t.Error("Position should be walkable after removing obstacle")
	}
}

func TestGridNeighbors(t *testing.T) {
	grid := NewGrid(5, 5)
	
	// Test neighbors without diagonal movement
	neighbors := grid.GetNeighbors(Point{2, 2}, false)
	expectedCount := 4 // up, down, left, right
	if len(neighbors) != expectedCount {
		t.Errorf("Expected %d neighbors, got %d", expectedCount, len(neighbors))
	}
	
	// Test neighbors with diagonal movement
	neighborsWithDiagonal := grid.GetNeighbors(Point{2, 2}, true)
	expectedCountWithDiagonal := 8 // all 8 directions
	if len(neighborsWithDiagonal) != expectedCountWithDiagonal {
		t.Errorf("Expected %d neighbors with diagonal, got %d", expectedCountWithDiagonal, len(neighborsWithDiagonal))
	}
	
	// Test edge case: corner position
	cornerNeighbors := grid.GetNeighbors(Point{0, 0}, false)
	if len(cornerNeighbors) != 2 { // only right and down
		t.Errorf("Expected 2 neighbors for corner, got %d", len(cornerNeighbors))
	}
	
	// Test with obstacles
	grid.AddObstacle(1, 2) // left of center
	grid.AddObstacle(3, 2) // right of center
	blockedNeighbors := grid.GetNeighbors(Point{2, 2}, false)
	if len(blockedNeighbors) != 2 { // only up and down
		t.Errorf("Expected 2 unblocked neighbors, got %d", len(blockedNeighbors))
	}
}

func TestAStar(t *testing.T) {
	grid := NewGrid(5, 5)
	astar := NewAStar(grid)
	
	// Test basic properties
	if astar.Grid != grid {
		t.Error("A* should reference the provided grid")
	}
	
	if astar.AllowDiagonal {
		t.Error("Diagonal movement should be disabled by default")
	}
	
	if astar.StraightCost != 1.0 {
		t.Errorf("Expected straight cost 1.0, got %.1f", astar.StraightCost)
	}
}

func TestAStarMovementCost(t *testing.T) {
	grid := NewGrid(5, 5)
	astar := NewAStar(grid)
	
	// Test straight movement cost
	straightCost := astar.CalculateMovementCost(Point{1, 1}, Point{2, 1})
	if straightCost != 1.0 {
		t.Errorf("Expected straight movement cost 1.0, got %.1f", straightCost)
	}
	
	// Test diagonal movement cost
	diagonalCost := astar.CalculateMovementCost(Point{1, 1}, Point{2, 2})
	if math.Abs(diagonalCost-math.Sqrt2) > 0.001 {
		t.Errorf("Expected diagonal movement cost √2, got %.3f", diagonalCost)
	}
}

func TestAStarSimplePath(t *testing.T) {
	grid := NewGrid(5, 5)
	astar := NewAStar(grid)
	
	start := Point{0, 0}
	goal := Point{4, 0}
	
	path, err := astar.FindPath(start, goal)
	if err != nil {
		t.Fatalf("Expected to find path, got error: %v", err)
	}
	
	expectedLength := 5 // 0,0 -> 1,0 -> 2,0 -> 3,0 -> 4,0
	if len(path) != expectedLength {
		t.Errorf("Expected path length %d, got %d", expectedLength, len(path))
	}
	
	if path[0] != start {
		t.Errorf("Path should start at %s, got %s", start.String(), path[0].String())
	}
	
	if path[len(path)-1] != goal {
		t.Errorf("Path should end at %s, got %s", goal.String(), path[len(path)-1].String())
	}
}

func TestAStarWithObstacles(t *testing.T) {
	grid := NewGrid(5, 5)
	astar := NewAStar(grid)
	
	// Create a wall
	grid.AddObstacle(2, 0)
	grid.AddObstacle(2, 1)
	grid.AddObstacle(2, 2)
	grid.AddObstacle(2, 3)
	// Leave 2,4 open for path
	
	start := Point{0, 2}
	goal := Point{4, 2}
	
	path, err := astar.FindPath(start, goal)
	if err != nil {
		t.Fatalf("Expected to find path around obstacle, got error: %v", err)
	}
	
	// Path should go around the wall
	if len(path) < 5 { // Should be longer than straight path due to obstacle
		t.Errorf("Path should be longer due to obstacle, got length %d", len(path))
	}
	
	// Verify path doesn't go through obstacles
	for _, point := range path {
		if grid.Obstacles[point] {
			t.Errorf("Path should not go through obstacle at %s", point.String())
		}
	}
}

func TestAStarNoPath(t *testing.T) {
	grid := NewGrid(5, 5)
	astar := NewAStar(grid)
	
	// Create complete wall blocking the path
	for y := 0; y < 5; y++ {
		grid.AddObstacle(2, y)
	}
	
	start := Point{0, 2}
	goal := Point{4, 2}
	
	path, err := astar.FindPath(start, goal)
	if err == nil {
		t.Error("Expected no path error, but path was found")
	}
	
	if path != nil {
		t.Error("Expected nil path when no path exists")
	}
}

func TestAStarDiagonalMovement(t *testing.T) {
	grid := NewGrid(5, 5)
	astar := NewAStar(grid)
	astar.SetAllowDiagonal(true)
	
	start := Point{0, 0}
	goal := Point{4, 4}
	
	path, err := astar.FindPath(start, goal)
	if err != nil {
		t.Fatalf("Expected to find diagonal path, got error: %v", err)
	}
	
	// With diagonal movement, shortest path should be 5 steps
	expectedLength := 5
	if len(path) != expectedLength {
		t.Errorf("Expected diagonal path length %d, got %d", expectedLength, len(path))
	}
}

func TestAStarSameStartGoal(t *testing.T) {
	grid := NewGrid(5, 5)
	astar := NewAStar(grid)
	
	start := Point{2, 2}
	goal := Point{2, 2}
	
	path, err := astar.FindPath(start, goal)
	if err != nil {
		t.Fatalf("Expected to find path for same start/goal, got error: %v", err)
	}
	
	if len(path) != 1 || path[0] != start {
		t.Errorf("Expected path with single point %s, got %v", start.String(), path)
	}
}

func TestAStarInvalidPositions(t *testing.T) {
	grid := NewGrid(5, 5)
	astar := NewAStar(grid)
	
	// Test invalid start position
	grid.AddObstacle(0, 0)
	_, err := astar.FindPath(Point{0, 0}, Point{4, 4})
	if err == nil {
		t.Error("Expected error for invalid start position")
	}
	
	// Test invalid goal position
	grid.RemoveObstacle(0, 0)
	grid.AddObstacle(4, 4)
	_, err = astar.FindPath(Point{0, 0}, Point{4, 4})
	if err == nil {
		t.Error("Expected error for invalid goal position")
	}
}

func TestAStarHeuristicChange(t *testing.T) {
	grid := NewGrid(5, 5)
	astar := NewAStar(grid)
	
	// Change heuristic to Euclidean
	astar.SetHeuristic(EuclideanDistance)
	
	start := Point{0, 0}
	goal := Point{4, 0}
	
	path, err := astar.FindPath(start, goal)
	if err != nil {
		t.Fatalf("Expected to find path with Euclidean heuristic, got error: %v", err)
	}
	
	if len(path) == 0 {
		t.Error("Expected non-empty path")
	}
}

func TestPathfindingResult(t *testing.T) {
	grid := NewGrid(5, 5)
	astar := NewAStar(grid)
	
	start := Point{0, 0}
	goal := Point{2, 0}
	
	result := astar.FindPathWithDetails(start, goal)
	
	if !result.Success {
		t.Error("Expected successful pathfinding")
	}
	
	if result.Error != nil {
		t.Errorf("Expected no error, got %v", result.Error)
	}
	
	if result.PathLength != len(result.Path) {
		t.Errorf("Path length mismatch: reported %d, actual %d", result.PathLength, len(result.Path))
	}
	
	if result.PathCost <= 0 {
		t.Errorf("Expected positive path cost, got %.1f", result.PathCost)
	}
}

func TestCalculatePathCost(t *testing.T) {
	grid := NewGrid(5, 5)
	astar := NewAStar(grid)
	
	// Test empty path
	emptyCost := astar.CalculatePathCost([]Point{})
	if emptyCost != 0 {
		t.Errorf("Expected 0 cost for empty path, got %.1f", emptyCost)
	}
	
	// Test single point path
	singleCost := astar.CalculatePathCost([]Point{{0, 0}})
	if singleCost != 0 {
		t.Errorf("Expected 0 cost for single point path, got %.1f", singleCost)
	}
	
	// Test straight path
	straightPath := []Point{{0, 0}, {1, 0}, {2, 0}}
	straightCost := astar.CalculatePathCost(straightPath)
	if straightCost != 2.0 {
		t.Errorf("Expected cost 2.0 for straight path, got %.1f", straightCost)
	}
	
	// Test diagonal path
	astar.SetAllowDiagonal(true)
	diagonalPath := []Point{{0, 0}, {1, 1}, {2, 2}}
	diagonalCost := astar.CalculatePathCost(diagonalPath)
	expectedDiagonalCost := 2 * math.Sqrt2
	if math.Abs(diagonalCost-expectedDiagonalCost) > 0.001 {
		t.Errorf("Expected diagonal cost %.3f, got %.3f", expectedDiagonalCost, diagonalCost)
	}
}

func TestVisualizeGrid(t *testing.T) {
	grid := NewGrid(3, 3)
	astar := NewAStar(grid)
	
	// Add obstacle
	grid.AddObstacle(1, 1)
	
	path := []Point{{0, 0}, {0, 1}, {0, 2}, {1, 2}, {2, 2}}
	
	visualization := astar.VisualizeGrid(path)
	
	// Check that visualization contains expected characters
	if !strings.Contains(visualization, "S") { // Start
		t.Error("Visualization should contain start marker 'S'")
	}
	
	if !strings.Contains(visualization, "G") { // Goal
		t.Error("Visualization should contain goal marker 'G'")
	}
	
	if !strings.Contains(visualization, "█") { // Obstacle
		t.Error("Visualization should contain obstacle marker '█'")
	}
	
	if !strings.Contains(visualization, "·") { // Path
		t.Error("Visualization should contain path marker '·'")
	}
}

func TestGetPathStatistics(t *testing.T) {
	grid := NewGrid(5, 5)
	astar := NewAStar(grid)
	
	// Test empty path
	emptyStats := astar.GetPathStatistics([]Point{})
	if emptyStats["length"] != 0 || emptyStats["cost"] != 0.0 {
		t.Error("Empty path should have 0 length and cost")
	}
	
	// Test straight path
	straightPath := []Point{{0, 0}, {1, 0}, {2, 0}}
	straightStats := astar.GetPathStatistics(straightPath)
	
	if straightStats["length"] != 3 {
		t.Errorf("Expected length 3, got %v", straightStats["length"])
	}
	
	if straightStats["cost"] != 2.0 {
		t.Errorf("Expected cost 2.0, got %v", straightStats["cost"])
	}
	
	if straightStats["straight_moves"] != 2 {
		t.Errorf("Expected 2 straight moves, got %v", straightStats["straight_moves"])
	}
	
	if straightStats["diagonal_moves"] != 0 {
		t.Errorf("Expected 0 diagonal moves, got %v", straightStats["diagonal_moves"])
	}
	
	// Test with diagonal movement
	astar.SetAllowDiagonal(true)
	diagonalPath := []Point{{0, 0}, {1, 1}, {2, 1}}
	diagonalStats := astar.GetPathStatistics(diagonalPath)
	
	if diagonalStats["diagonal_moves"] != 1 {
		t.Errorf("Expected 1 diagonal move, got %v", diagonalStats["diagonal_moves"])
	}
	
	if diagonalStats["straight_moves"] != 1 {
		t.Errorf("Expected 1 straight move, got %v", diagonalStats["straight_moves"])
	}
}

func TestComplexMaze(t *testing.T) {
	// Create a more complex maze for thorough testing
	grid := NewGrid(10, 10)
	astar := NewAStar(grid)
	
	// Create a maze pattern
	obstacles := []Point{
		{1, 0}, {1, 1}, {1, 2}, {1, 3}, {1, 4}, {1, 5}, {1, 6}, {1, 7},
		{3, 2}, {3, 3}, {3, 4}, {3, 5}, {3, 6}, {3, 7}, {3, 8}, {3, 9},
		{5, 0}, {5, 1}, {5, 2}, {5, 3}, {5, 4}, {5, 5}, {5, 6},
		{7, 2}, {7, 3}, {7, 4}, {7, 5}, {7, 6}, {7, 7}, {7, 8},
	}
	
	for _, obs := range obstacles {
		grid.AddObstacle(obs.X, obs.Y)
	}
	
	start := Point{0, 0}
	goal := Point{9, 9}
	
	path, err := astar.FindPath(start, goal)
	if err != nil {
		t.Fatalf("Expected to find path through complex maze, got error: %v", err)
	}
	
	// Verify path doesn't go through obstacles
	for _, point := range path {
		if grid.Obstacles[point] {
			t.Errorf("Path should not go through obstacle at %s", point.String())
		}
	}
	
	// Verify path is connected (each step is adjacent)
	for i := 1; i < len(path); i++ {
		prev := path[i-1]
		curr := path[i]
		
		dx := math.Abs(float64(curr.X - prev.X))
		dy := math.Abs(float64(curr.Y - prev.Y))
		
		// Should be adjacent (difference of at most 1 in each direction)
		if dx > 1 || dy > 1 {
			t.Errorf("Path not connected: jump from %s to %s", prev.String(), curr.String())
		}
		
		// Should not be the same point
		if dx == 0 && dy == 0 {
			t.Errorf("Path contains duplicate point: %s", curr.String())
		}
	}
}

// Benchmark tests
func BenchmarkAStarSimplePath(b *testing.B) {
	grid := NewGrid(20, 20)
	astar := NewAStar(grid)
	start := Point{0, 0}
	goal := Point{19, 19}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		astar.FindPath(start, goal)
	}
}

func BenchmarkAStarComplexMaze(b *testing.B) {
	grid := NewGrid(50, 50)
	astar := NewAStar(grid)
	
	// Add some obstacles
	for i := 0; i < 50; i += 3 {
		for j := 0; j < 25; j++ {
			grid.AddObstacle(i, j)
		}
	}
	
	start := Point{0, 49}
	goal := Point{49, 0}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		astar.FindPath(start, goal)
	}
}

func BenchmarkAStarDiagonal(b *testing.B) {
	grid := NewGrid(30, 30)
	astar := NewAStar(grid)
	astar.SetAllowDiagonal(true)
	
	start := Point{0, 0}
	goal := Point{29, 29}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		astar.FindPath(start, goal)
	}
}
