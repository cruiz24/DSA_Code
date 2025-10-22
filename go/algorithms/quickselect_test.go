package algorithms

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestQuickSelect(t *testing.T) {
	tests := []struct {
		arr      []int
		k        int
		expected int
	}{
		{[]int{3, 6, 8, 10, 1, 2, 1}, 3, 3},
		{[]int{7, 10, 4, 3, 20, 15}, 3, 7},
		{[]int{1, 2, 3, 4, 5}, 1, 1},
		{[]int{1, 2, 3, 4, 5}, 5, 5},
		{[]int{5, 4, 3, 2, 1}, 3, 3},
		{[]int{1}, 1, 1},
		{[]int{2, 1}, 1, 1},
		{[]int{2, 1}, 2, 2},
	}

	for i, test := range tests {
		result, err := QuickSelect(test.arr, test.k)
		if err != nil {
			t.Errorf("Test %d: Unexpected error: %v", i, err)
			continue
		}

		// Verify by sorting and checking kth element
		sorted := make([]int, len(test.arr))
		copy(sorted, test.arr)
		sort.Ints(sorted)

		if result != sorted[test.k-1] {
			t.Errorf("Test %d: Expected %d, got %d", i, sorted[test.k-1], result)
		}
	}
}

func TestQuickSelectIterative(t *testing.T) {
	tests := []struct {
		arr []int
		k   int
	}{
		{[]int{3, 6, 8, 10, 1, 2, 1}, 3},
		{[]int{7, 10, 4, 3, 20, 15}, 3},
		{[]int{1, 2, 3, 4, 5}, 1},
		{[]int{1, 2, 3, 4, 5}, 5},
	}

	for i, test := range tests {
		recursive, _ := QuickSelect(test.arr, test.k)
		iterative, err := QuickSelectIterative(test.arr, test.k)
		
		if err != nil {
			t.Errorf("Test %d: Unexpected error in iterative: %v", i, err)
			continue
		}

		if recursive != iterative {
			t.Errorf("Test %d: Recursive (%d) != Iterative (%d)", i, recursive, iterative)
		}
	}
}

func TestQuickSelectErrors(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}

	// Test k out of bounds
	_, err := QuickSelect(arr, 0)
	if err == nil {
		t.Error("Expected error for k=0")
	}

	_, err = QuickSelect(arr, 6)
	if err == nil {
		t.Error("Expected error for k=6 (greater than array length)")
	}

	// Test empty array
	_, err = QuickSelect([]int{}, 1)
	if err == nil {
		t.Error("Expected error for empty array")
	}
}

func TestFindKthLargest(t *testing.T) {
	tests := []struct {
		arr      []int
		k        int
		expected int
	}{
		{[]int{3, 2, 1, 5, 6, 4}, 2, 5}, // 2nd largest
		{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4, 4}, // 4th largest
		{[]int{1}, 1, 1},
		{[]int{1, 2}, 1, 2}, // 1st largest
		{[]int{1, 2}, 2, 1}, // 2nd largest
	}

	for i, test := range tests {
		result, err := FindKthLargest(test.arr, test.k)
		if err != nil {
			t.Errorf("Test %d: Unexpected error: %v", i, err)
			continue
		}

		// Verify by sorting in descending order
		sorted := make([]int, len(test.arr))
		copy(sorted, test.arr)
		sort.Sort(sort.Reverse(sort.IntSlice(sorted)))

		if result != sorted[test.k-1] {
			t.Errorf("Test %d: Expected %d, got %d", i, sorted[test.k-1], result)
		}
	}
}

func TestFindMedian(t *testing.T) {
	tests := []struct {
		arr      []int
		expected float64
	}{
		{[]int{1, 2, 3, 4, 5}, 3.0},     // Odd length
		{[]int{1, 2, 3, 4}, 2.5},        // Even length
		{[]int{1}, 1.0},                 // Single element
		{[]int{2, 1}, 1.5},              // Two elements
		{[]int{3, 1, 2}, 2.0},           // Odd length, unsorted
		{[]int{4, 1, 3, 2}, 2.5},        // Even length, unsorted
	}

	for i, test := range tests {
		result, err := FindMedian(test.arr)
		if err != nil {
			t.Errorf("Test %d: Unexpected error: %v", i, err)
			continue
		}

		if result != test.expected {
			t.Errorf("Test %d: Expected %.1f, got %.1f", i, test.expected, result)
		}
	}

	// Test empty array
	_, err := FindMedian([]int{})
	if err == nil {
		t.Error("Expected error for empty array")
	}
}

func TestTopK(t *testing.T) {
	tests := []struct {
		arr []int
		k   int
	}{
		{[]int{4, 5, 8, 2}, 3},
		{[]int{3, 2, 1, 5, 6, 4}, 2},
		{[]int{1}, 1},
		{[]int{1, 1, 1, 1}, 2},
	}

	for i, test := range tests {
		result, err := TopK(test.arr, test.k)
		if err != nil {
			t.Errorf("Test %d: Unexpected error: %v", i, err)
			continue
		}

		if len(result) != test.k {
			t.Errorf("Test %d: Expected %d elements, got %d", i, test.k, len(result))
			continue
		}

		// Verify by sorting original and comparing
		sorted := make([]int, len(test.arr))
		copy(sorted, test.arr)
		sort.Ints(sorted)
		expectedTopK := sorted[:test.k]

		sort.Ints(result)
		sort.Ints(expectedTopK)

		if !reflect.DeepEqual(result, expectedTopK) {
			t.Errorf("Test %d: Expected %v, got %v", i, expectedTopK, result)
		}
	}
}

func TestQuickSelectWithCustomComparator(t *testing.T) {
	arr := []int{3, 1, 4, 1, 5, 9, 2, 6}
	
	// Test with reverse comparator (descending order)
	reverseComparator := func(a, b int) bool {
		return a > b
	}
	
	result, err := QuickSelectWithCustomComparator(arr, 3, reverseComparator)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	
	// With reverse comparator, 3rd element should be 3rd largest
	sorted := make([]int, len(arr))
	copy(sorted, arr)
	sort.Sort(sort.Reverse(sort.IntSlice(sorted)))
	
	expected := sorted[2] // 3rd largest (0-indexed)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestCalculateStatistics(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	
	stats, err := CalculateStatistics(arr)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	
	if stats.Min != 1 {
		t.Errorf("Expected min 1, got %d", stats.Min)
	}
	
	if stats.Max != 10 {
		t.Errorf("Expected max 10, got %d", stats.Max)
	}
	
	if stats.Median != 5.5 {
		t.Errorf("Expected median 5.5, got %.1f", stats.Median)
	}
	
	// Test with empty array
	_, err = CalculateStatistics([]int{})
	if err == nil {
		t.Error("Expected error for empty array")
	}
}

// Benchmark tests
func BenchmarkQuickSelect(b *testing.B) {
	sizes := []int{100, 1000, 10000}
	
	for _, size := range sizes {
		arr := make([]int, size)
		for i := range arr {
			arr[i] = size - i // Worst case: reverse sorted
		}
		
		b.Run(fmt.Sprintf("Size_%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				QuickSelect(arr, size/2)
			}
		})
	}
}

func BenchmarkQuickSelectIterative(b *testing.B) {
	arr := make([]int, 10000)
	for i := range arr {
		arr[i] = 10000 - i
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		QuickSelectIterative(arr, 5000)
	}
}

func BenchmarkFindMedian(b *testing.B) {
	arr := make([]int, 10000)
	for i := range arr {
		arr[i] = i
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FindMedian(arr)
	}
}
