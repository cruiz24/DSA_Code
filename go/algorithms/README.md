# Quickselect Algorithm Implementation in Go

A comprehensive implementation of the Quickselect algorithm for finding the k-th smallest element in an array in expected O(n) time.

## Features

- **Quickselect Algorithm**: Find k-th smallest element efficiently
- **Iterative Version**: Non-recursive implementation available
- **K-th Largest**: Find k-th largest element
- **Median Finding**: Efficient median calculation
- **Top-K Elements**: Get k smallest elements
- **Custom Comparators**: Support for custom comparison functions
- **Statistical Analysis**: Calculate various statistics

## Time Complexity

| Operation | Average Case | Worst Case | Best Case |
|-----------|-------------|-----------|-----------|
| Quickselect | O(n)      | O(n²)     | O(n)      |
| Find Median | O(n)      | O(n²)     | O(n)      |
| Top-K      | O(n)      | O(n²)     | O(n)      |

## Space Complexity: O(log n) for recursion stack

## Usage

```go
package main

import (
    "fmt"
    "algorithms"
)

func main() {
    arr := []int{3, 6, 8, 10, 1, 2, 1}
    
    // Find 3rd smallest element
    kth, err := algorithms.QuickSelect(arr, 3)
    if err == nil {
        fmt.Printf("3rd smallest: %d\n", kth) // Output: 3
    }
    
    // Find 2nd largest element
    largest, err := algorithms.FindKthLargest(arr, 2)
    if err == nil {
        fmt.Printf("2nd largest: %d\n", largest) // Output: 8
    }
    
    // Find median
    median, err := algorithms.FindMedian(arr)
    if err == nil {
        fmt.Printf("Median: %.1f\n", median) // Output: 3.0
    }
    
    // Get top 3 smallest elements
    topK, err := algorithms.TopK(arr, 3)
    if err == nil {
        fmt.Printf("Top 3: %v\n", topK) // Output: [1, 1, 2]
    }
    
    // Calculate comprehensive statistics
    stats, err := algorithms.CalculateStatistics(arr)
    if err == nil {
        fmt.Printf("Min: %d, Max: %d\n", stats.Min, stats.Max)
        fmt.Printf("Median: %.1f\n", stats.Median)
        fmt.Printf("Q1: %d, Q3: %d, IQR: %d\n", stats.Q1, stats.Q3, stats.IQR)
    }
    
    // Use custom comparator for descending order
    reverseComparator := func(a, b int) bool { return a > b }
    result, _ := algorithms.QuickSelectWithCustomComparator(arr, 3, reverseComparator)
    fmt.Printf("3rd largest (custom): %d\n", result)
}
```

## Algorithm Details

### Quickselect Process
1. Choose a pivot element (randomized for better average performance)
2. Partition array around pivot
3. If pivot is at position k-1, return pivot
4. If k-1 < pivot position, search left partition
5. If k-1 > pivot position, search right partition

### Partition Strategy
- **Lomuto Partition**: Used for simplicity and clarity
- **Random Pivot**: Avoids worst-case O(n²) on sorted arrays
- **In-place Operation**: Minimal extra space usage

## Applications

- **Selection Problems**: Finding k-th order statistics
- **Median Finding**: Efficient median calculation for large datasets
- **Percentile Calculation**: Finding percentiles in data analysis
- **Database Operations**: SQL ORDER BY LIMIT optimization
- **Streaming Algorithms**: Approximate quantiles in data streams

## Comparison with Other Algorithms

| Algorithm | Average Time | Worst Time | Space | Stable |
|-----------|-------------|------------|-------|--------|
| Quickselect | O(n)      | O(n²)     | O(log n) | No |
| Heap Select | O(n + k log n) | O(n + k log n) | O(k) | No |
| Sort + Index | O(n log n) | O(n log n) | O(1) | Depends |

## Advanced Features

### Statistical Analysis
```go
stats, _ := CalculateStatistics(data)
// Returns: Min, Max, Median, Q1, Q3, IQR
```

### Custom Comparators
```go
// For custom data types or sorting criteria
comparator := func(a, b int) bool { return abs(a) < abs(b) }
result, _ := QuickSelectWithCustomComparator(arr, k, comparator)
```

### Iterative Implementation
```go
// Avoids recursion stack overhead
result, _ := QuickSelectIterative(arr, k)
```

## Testing

Comprehensive test suite included:

```bash
go test -v                 # Run all tests
go test -bench=.          # Run benchmarks
go test -race             # Test for race conditions
```

## Optimization Notes

- **Random Pivot**: Prevents worst-case on sorted inputs
- **Tail Recursion**: Can be optimized to iterative form
- **Cache Locality**: Good performance on modern processors
- **Introspective**: Can fall back to heap-select for worst cases

## Educational Value

This implementation demonstrates:
- Divide and conquer algorithms
- Randomization in algorithms
- Average vs. worst-case analysis
- Practical algorithm engineering
- Statistical computing applications
