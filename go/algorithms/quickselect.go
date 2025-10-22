package algorithms

import (
	"fmt"
	"math/rand"
	"time"
)

// QuickSelect finds the kth smallest element in an array using quickselect algorithm
// Time complexity: Average O(n), Worst case O(n²)
// Space complexity: O(log n) for recursion stack
func QuickSelect(arr []int, k int) (int, error) {
	if k < 1 || k > len(arr) {
		return 0, fmt.Errorf("k must be between 1 and %d", len(arr))
	}
	if len(arr) == 0 {
		return 0, fmt.Errorf("array cannot be empty")
	}
	
	// Create a copy to avoid modifying original array
	nums := make([]int, len(arr))
	copy(nums, arr)
	
	return quickSelectHelper(nums, 0, len(nums)-1, k-1)
}

// quickSelectHelper performs the quickselect algorithm recursively
func quickSelectHelper(arr []int, left, right, k int) (int, error) {
	if left == right {
		return arr[left], nil
	}
	
	// Choose random pivot to avoid worst case
	pivotIndex := randomPartition(arr, left, right)
	
	if k == pivotIndex {
		return arr[k], nil
	} else if k < pivotIndex {
		return quickSelectHelper(arr, left, pivotIndex-1, k)
	} else {
		return quickSelectHelper(arr, pivotIndex+1, right, k)
	}
}

// QuickSelectIterative is an iterative implementation of quickselect
func QuickSelectIterative(arr []int, k int) (int, error) {
	if k < 1 || k > len(arr) {
		return 0, fmt.Errorf("k must be between 1 and %d", len(arr))
	}
	if len(arr) == 0 {
		return 0, fmt.Errorf("array cannot be empty")
	}
	
	nums := make([]int, len(arr))
	copy(nums, arr)
	
	left, right := 0, len(nums)-1
	k = k - 1 // Convert to 0-based index
	
	for left < right {
		pivotIndex := randomPartition(nums, left, right)
		
		if k == pivotIndex {
			return nums[k], nil
		} else if k < pivotIndex {
			right = pivotIndex - 1
		} else {
			left = pivotIndex + 1
		}
	}
	
	return nums[k], nil
}

// FindKthLargest finds the kth largest element (1-indexed)
func FindKthLargest(arr []int, k int) (int, error) {
	if k < 1 || k > len(arr) {
		return 0, fmt.Errorf("k must be between 1 and %d", len(arr))
	}
	
	// kth largest = (n-k+1)th smallest
	return QuickSelect(arr, len(arr)-k+1)
}

// FindMedian finds the median of the array
func FindMedian(arr []int) (float64, error) {
	if len(arr) == 0 {
		return 0, fmt.Errorf("array cannot be empty")
	}
	
	n := len(arr)
	if n%2 == 1 {
		// Odd length - return middle element
		mid, err := QuickSelect(arr, (n+1)/2)
		return float64(mid), err
	} else {
		// Even length - return average of two middle elements
		mid1, err1 := QuickSelect(arr, n/2)
		if err1 != nil {
			return 0, err1
		}
		mid2, err2 := QuickSelect(arr, n/2+1)
		if err2 != nil {
			return 0, err2
		}
		return float64(mid1+mid2) / 2.0, nil
	}
}

// TopK finds the k smallest elements (not necessarily sorted)
func TopK(arr []int, k int) ([]int, error) {
	if k < 1 || k > len(arr) {
		return nil, fmt.Errorf("k must be between 1 and %d", len(arr))
	}
	if len(arr) == 0 {
		return nil, fmt.Errorf("array cannot be empty")
	}
	
	nums := make([]int, len(arr))
	copy(nums, arr)
	
	// Find the kth smallest element
	kthElement, err := QuickSelect(nums, k)
	if err != nil {
		return nil, err
	}
	
	// Collect all elements <= kth element
	result := make([]int, 0, k)
	count := 0
	
	for _, num := range nums {
		if num < kthElement {
			result = append(result, num)
			count++
		}
	}
	
	// Add remaining spots with kth element value
	for count < k {
		result = append(result, kthElement)
		count++
	}
	
	return result, nil
}

// randomPartition partitions array with random pivot
func randomPartition(arr []int, left, right int) int {
	// Choose random pivot
	rand.Seed(time.Now().UnixNano())
	randomIndex := left + rand.Intn(right-left+1)
	
	// Swap with last element
	arr[randomIndex], arr[right] = arr[right], arr[randomIndex]
	
	return partition(arr, left, right)
}

// partition partitions the array around the last element as pivot
func partition(arr []int, left, right int) int {
	pivot := arr[right]
	i := left
	
	for j := left; j < right; j++ {
		if arr[j] <= pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	
	arr[i], arr[right] = arr[right], arr[i]
	return i
}

// QuickSelectWithCustomComparator allows custom comparison function
func QuickSelectWithCustomComparator(arr []int, k int, less func(a, b int) bool) (int, error) {
	if k < 1 || k > len(arr) {
		return 0, fmt.Errorf("k must be between 1 and %d", len(arr))
	}
	if len(arr) == 0 {
		return 0, fmt.Errorf("array cannot be empty")
	}
	
	nums := make([]int, len(arr))
	copy(nums, arr)
	
	return quickSelectWithComparator(nums, 0, len(nums)-1, k-1, less)
}

// quickSelectWithComparator helper function for custom comparator
func quickSelectWithComparator(arr []int, left, right, k int, less func(a, b int) bool) (int, error) {
	if left == right {
		return arr[left], nil
	}
	
	pivotIndex := partitionWithComparator(arr, left, right, less)
	
	if k == pivotIndex {
		return arr[k], nil
	} else if k < pivotIndex {
		return quickSelectWithComparator(arr, left, pivotIndex-1, k, less)
	} else {
		return quickSelectWithComparator(arr, pivotIndex+1, right, k, less)
	}
}

// partitionWithComparator partitions with custom comparator
func partitionWithComparator(arr []int, left, right int, less func(a, b int) bool) int {
	pivot := arr[right]
	i := left
	
	for j := left; j < right; j++ {
		if less(arr[j], pivot) || arr[j] == pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	
	arr[i], arr[right] = arr[right], arr[i]
	return i
}

// Statistics provides various statistical measures using quickselect
type Statistics struct {
	Min, Max     int
	Median       float64
	Q1, Q3       int // First and third quartiles
	IQR          int // Interquartile range
}

// CalculateStatistics computes various statistics for the array
func CalculateStatistics(arr []int) (*Statistics, error) {
	if len(arr) == 0 {
		return nil, fmt.Errorf("array cannot be empty")
	}
	
	n := len(arr)
	stats := &Statistics{}
	
	// Min and Max using quickselect
	min, err := QuickSelect(arr, 1)
	if err != nil {
		return nil, err
	}
	stats.Min = min
	
	max, err := QuickSelect(arr, n)
	if err != nil {
		return nil, err
	}
	stats.Max = max
	
	// Median
	median, err := FindMedian(arr)
	if err != nil {
		return nil, err
	}
	stats.Median = median
	
	// Quartiles
	q1, err := QuickSelect(arr, n/4+1)
	if err != nil {
		return nil, err
	}
	stats.Q1 = q1
	
	q3, err := QuickSelect(arr, 3*n/4+1)
	if err != nil {
		return nil, err
	}
	stats.Q3 = q3
	
	stats.IQR = stats.Q3 - stats.Q1
	
	return stats, nil
}
