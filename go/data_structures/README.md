# Binary Heap Implementation in Go

A complete implementation of binary heap data structure supporting both min-heap and max-heap operations.

## Features

- **Min Heap & Max Heap**: Support for both heap types
- **Heap Operations**: Insert, extract, peek, build heap
- **Heap Sort**: Complete heap sort implementation
- **K-th Element**: Find k-th largest element efficiently
- **Validation**: Heap property validation
- **Comprehensive Testing**: Full test coverage with benchmarks

## Time Complexity

| Operation | Time Complexity |
|-----------|----------------|
| Insert    | O(log n)       |
| Extract   | O(log n)       |
| Peek      | O(1)           |
| Build Heap| O(n)           |
| Heap Sort | O(n log n)     |

## Usage

```go
package main

import (
    "fmt"
    "data_structures"
)

func main() {
    // Create max heap
    maxHeap := data_structures.NewMaxHeap()
    
    // Insert elements
    maxHeap.Insert(10)
    maxHeap.Insert(5)
    maxHeap.Insert(15)
    
    // Extract max element
    max, err := maxHeap.Extract()
    fmt.Printf("Max: %d\n", max) // Output: Max: 15
    
    // Create min heap and build from array
    minHeap := data_structures.NewMinHeap()
    arr := []int{4, 1, 7, 3, 8, 5}
    minHeap.BuildHeap(arr)
    
    // Extract elements (sorted order)
    for !minHeap.IsEmpty() {
        val, _ := minHeap.Extract()
        fmt.Printf("%d ", val) // Output: 1 3 4 5 7 8
    }
    
    // Heap sort
    heap := data_structures.NewMaxHeap()
    unsorted := []int{64, 34, 25, 12, 22, 11, 90}
    sorted := heap.HeapSort(unsorted)
    fmt.Printf("Sorted: %v\n", sorted)
    
    // Find k-th largest element
    kth, _ := heap.GetKthLargest([]int{3, 2, 1, 5, 6, 4}, 2)
    fmt.Printf("2nd largest: %d\n", kth) // Output: 5
}
```

## Applications

- **Priority Queues**: Efficient priority queue implementation
- **Heap Sort**: In-place sorting algorithm
- **Selection Algorithms**: Finding k-th smallest/largest elements
- **Graph Algorithms**: Dijkstra's algorithm, Prim's MST
- **Memory Management**: Heap allocation algorithms

## Implementation Details

### Heap Property
- **Max Heap**: Parent ≥ Children
- **Min Heap**: Parent ≤ Children

### Array Representation
- Root at index 0
- Left child of node i: 2*i + 1
- Right child of node i: 2*i + 2
- Parent of node i: (i-1)/2

### Key Methods

- `Insert(item)`: Add element maintaining heap property
- `Extract()`: Remove and return root element
- `Peek()`: Return root without removal
- `BuildHeap(array)`: Create heap from existing array
- `HeapSort(array)`: Sort array using heap sort
- `GetKthLargest(array, k)`: Find k-th largest element

## Testing

Run the comprehensive test suite:

```bash
go test -v
```

Run benchmarks:

```bash
go test -bench=.
```

## Advanced Features

- **Heap Validation**: Verify heap property maintenance
- **Statistics**: Calculate various heap statistics
- **Memory Efficient**: In-place operations where possible
- **Thread Safety**: Can be made thread-safe with mutex wrapper

## Educational Value

This implementation demonstrates:
- Complete binary tree properties
- Heap algorithms and optimizations
- Array-based tree representation
- Priority queue applications
- Algorithm analysis and complexity
