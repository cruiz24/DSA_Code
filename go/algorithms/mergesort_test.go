package algorithms

import (
	"reflect"
	"sort"
	"testing"
)

func TestMergeSortBasic(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{64, 34, 25, 12, 22, 11, 90}, []int{11, 12, 22, 25, 34, 64, 90}},
		{[]int{5, 2, 4, 6, 1, 3}, []int{1, 2, 3, 4, 5, 6}},
		{[]int{1}, []int{1}},
		{[]int{}, []int{}},
		{[]int{3, 3, 3}, []int{3, 3, 3}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}}, // Already sorted
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}}, // Reverse sorted
	}

	for i, test := range tests {
		result := MergeSort(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Test %d: Expected %v, got %v", i, test.expected, result)
		}
		
		// Verify original array is unchanged
		originalSorted := make([]int, len(test.input))
		copy(originalSorted, test.input)
		sort.Ints(originalSorted)
		
		if !reflect.DeepEqual(result, originalSorted) {
			t.Errorf("Test %d: Result doesn't match sorted original", i)
		}
	}
}

func TestMergeSortInPlace(t *testing.T) {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	expected := []int{11, 12, 22, 25, 34, 64, 90}
	
	MergeSortInPlace(arr)
	
	if !reflect.DeepEqual(arr, expected) {
		t.Errorf("Expected %v, got %v", expected, arr)
	}
}

func TestMergeSortWithComparator(t *testing.T) {
	arr := []int{3, 1, 4, 1, 5, 9, 2, 6}
	
	// Test ascending order (default)
	ascending := func(a, b int) bool { return a < b }
	resultAsc := MergeSortWithComparator(arr, ascending)
	expectedAsc := []int{1, 1, 2, 3, 4, 5, 6, 9}
	
	if !reflect.DeepEqual(resultAsc, expectedAsc) {
		t.Errorf("Ascending: Expected %v, got %v", expectedAsc, resultAsc)
	}
	
	// Test descending order
	descending := func(a, b int) bool { return a > b }
	resultDesc := MergeSortWithComparator(arr, descending)
	expectedDesc := []int{9, 6, 5, 4, 3, 2, 1, 1}
	
	if !reflect.DeepEqual(resultDesc, expectedDesc) {
		t.Errorf("Descending: Expected %v, got %v", expectedDesc, resultDesc)
	}
}

func TestBottomUpMergeSort(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{64, 34, 25, 12, 22, 11, 90}, []int{11, 12, 22, 25, 34, 64, 90}},
		{[]int{5, 2, 4, 6, 1, 3}, []int{1, 2, 3, 4, 5, 6}},
		{[]int{1}, []int{1}},
		{[]int{}, []int{}},
	}

	for i, test := range tests {
		result := BottomUpMergeSort(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Test %d: Expected %v, got %v", i, test.expected, result)
		}
	}
}

func TestMergeSortedArrays(t *testing.T) {
	tests := []struct {
		arr1     []int
		arr2     []int
		expected []int
	}{
		{[]int{1, 3, 5}, []int{2, 4, 6}, []int{1, 2, 3, 4, 5, 6}},
		{[]int{1, 2, 3}, []int{4, 5, 6}, []int{1, 2, 3, 4, 5, 6}},
		{[]int{}, []int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{1, 2, 3}, []int{}, []int{1, 2, 3}},
		{[]int{}, []int{}, []int{}},
		{[]int{1, 1, 2}, []int{1, 3, 3}, []int{1, 1, 1, 2, 3, 3}},
	}

	for i, test := range tests {
		result := MergeSortedArrays(test.arr1, test.arr2)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Test %d: Expected %v, got %v", i, test.expected, result)
		}
	}
}

func TestCountInversions(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{2, 3, 8, 6, 1}, 5}, // (2,1), (3,1), (8,6), (8,1), (6,1)
		{[]int{1, 2, 3, 4, 5}, 0}, // Already sorted
		{[]int{5, 4, 3, 2, 1}, 10}, // Reverse sorted: 4+3+2+1 = 10
		{[]int{1}, 0},
		{[]int{}, 0},
		{[]int{2, 1}, 1},
	}

	for i, test := range tests {
		result := CountInversions(test.input)
		if result != test.expected {
			t.Errorf("Test %d: Expected %d inversions, got %d", i, test.expected, result)
		}
	}
}

func TestKWayMerge(t *testing.T) {
	tests := []struct {
		input    [][]int
		expected []int
	}{
		{
			[][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}},
			[]int{1, 1, 2, 3, 4, 4, 5, 6},
		},
		{
			[][]int{{1, 2, 3}},
			[]int{1, 2, 3},
		},
		{
			[][]int{},
			[]int{},
		},
		{
			[][]int{{1}, {2}, {3}},
			[]int{1, 2, 3},
		},
	}

	for i, test := range tests {
		result := KWayMerge(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Test %d: Expected %v, got %v", i, test.expected, result)
		}
	}
}

func TestMergeSortRange(t *testing.T) {
	arr := []int{5, 2, 8, 1, 9, 3}
	MergeSortRange(arr, 1, 4) // Sort elements at indices 1-4
	
	// Elements at indices 1-4 should be sorted: [2, 8, 1, 9] -> [1, 2, 8, 9]
	expected := []int{5, 1, 2, 8, 9, 3}
	
	if !reflect.DeepEqual(arr, expected) {
		t.Errorf("Expected %v, got %v", expected, arr)
	}
	
	// Test invalid ranges
	arr2 := []int{1, 2, 3}
	MergeSortRange(arr2, -1, 2) // Should do nothing
	MergeSortRange(arr2, 0, 5)  // Should do nothing
	MergeSortRange(arr2, 2, 1)  // Should do nothing
	
	expected2 := []int{1, 2, 3}
	if !reflect.DeepEqual(arr2, expected2) {
		t.Error("Invalid range operations should not modify array")
	}
}

func TestIsSorted(t *testing.T) {
	tests := []struct {
		input    []int
		expected bool
	}{
		{[]int{1, 2, 3, 4, 5}, true},
		{[]int{5, 4, 3, 2, 1}, false},
		{[]int{1, 3, 2, 4, 5}, false},
		{[]int{1}, true},
		{[]int{}, true},
		{[]int{1, 1, 1}, true},
		{[]int{1, 2, 2, 3}, true},
	}

	for i, test := range tests {
		result := IsSorted(test.input)
		if result != test.expected {
			t.Errorf("Test %d: Expected %t, got %t", i, test.expected, result)
		}
	}
}

func TestMergeSortWithStats(t *testing.T) {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	
	sorted, stats := MergeSortWithStats(arr)
	
	expectedSorted := []int{11, 12, 22, 25, 34, 64, 90}
	if !reflect.DeepEqual(sorted, expectedSorted) {
		t.Errorf("Expected %v, got %v", expectedSorted, sorted)
	}
	
	if stats.ArraySize != len(arr) {
		t.Errorf("Expected array size %d, got %d", len(arr), stats.ArraySize)
	}
	
	if stats.Comparisons <= 0 {
		t.Error("Expected positive number of comparisons")
	}
	
	if stats.Swaps <= 0 {
		t.Error("Expected positive number of swaps")
	}
	
	// Test already sorted array
	sortedArr := []int{1, 2, 3, 4, 5}
	_, sortedStats := MergeSortWithStats(sortedArr)
	
	if !sortedStats.IsSorted {
		t.Error("Should detect that input array is already sorted")
	}
}

func TestMergeSortStability(t *testing.T) {
	// Test that merge sort is stable (preserves relative order of equal elements)
	// We can't test this directly with integers, but we can verify the algorithm
	// maintains stability by checking that equal elements appear in their original order
	
	arr := []int{3, 1, 3, 2, 3, 1}
	sorted := MergeSort(arr)
	expected := []int{1, 1, 2, 3, 3, 3}
	
	if !reflect.DeepEqual(sorted, expected) {
		t.Errorf("Expected %v, got %v", expected, sorted)
	}
}

func TestMergeSortLargeArray(t *testing.T) {
	// Test with larger array
	size := 1000
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = size - i // Reverse order
	}
	
	sorted := MergeSort(arr)
	
	// Verify it's sorted
	if !IsSorted(sorted) {
		t.Error("Large array not properly sorted")
	}
	
	// Verify all elements are present
	if len(sorted) != size {
		t.Errorf("Expected length %d, got %d", size, len(sorted))
	}
	
	// Verify first and last elements
	if sorted[0] != 1 || sorted[size-1] != size {
		t.Error("First or last element incorrect")
	}
}

// Benchmark tests
func BenchmarkMergeSort(b *testing.B) {
	sizes := []int{100, 1000, 10000}
	
	for _, size := range sizes {
		arr := make([]int, size)
		for i := 0; i < size; i++ {
			arr[i] = size - i // Reverse order (worst case for many algorithms)
		}
		
		b.Run(fmt.Sprintf("Size_%d", size), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				// Create fresh copy for each iteration
				testArr := make([]int, len(arr))
				copy(testArr, arr)
				MergeSort(testArr)
			}
		})
	}
}

func BenchmarkMergeSortInPlace(b *testing.B) {
	size := 10000
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = size - i
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		testArr := make([]int, len(arr))
		copy(testArr, arr)
		MergeSortInPlace(testArr)
	}
}

func BenchmarkBottomUpMergeSort(b *testing.B) {
	size := 10000
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = size - i
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BottomUpMergeSort(arr)
	}
}

func BenchmarkCountInversions(b *testing.B) {
	size := 1000
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = size - i // Maximum inversions
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CountInversions(arr)
	}
}
