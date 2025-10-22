package algorithms

import (
	"fmt"
)

// MergeSort sorts an array using the merge sort algorithm
// Time complexity: O(n log n) in all cases
// Space complexity: O(n)
func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	
	// Create a copy to avoid modifying original
	result := make([]int, len(arr))
	copy(result, arr)
	
	mergeSortHelper(result, 0, len(result)-1)
	return result
}

// MergeSortInPlace sorts the array in place
func MergeSortInPlace(arr []int) {
	if len(arr) <= 1 {
		return
	}
	mergeSortHelper(arr, 0, len(arr)-1)
}

// mergeSortHelper recursively sorts the array
func mergeSortHelper(arr []int, left, right int) {
	if left >= right {
		return
	}
	
	mid := left + (right-left)/2
	
	// Recursively sort left and right halves
	mergeSortHelper(arr, left, mid)
	mergeSortHelper(arr, mid+1, right)
	
	// Merge the sorted halves
	merge(arr, left, mid, right)
}

// merge combines two sorted subarrays
func merge(arr []int, left, mid, right int) {
	// Create temporary arrays for left and right subarrays
	leftSize := mid - left + 1
	rightSize := right - mid
	
	leftArr := make([]int, leftSize)
	rightArr := make([]int, rightSize)
	
	// Copy data to temporary arrays
	copy(leftArr, arr[left:mid+1])
	copy(rightArr, arr[mid+1:right+1])
	
	// Merge the temporary arrays back into arr[left..right]
	i, j, k := 0, 0, left
	
	for i < leftSize && j < rightSize {
		if leftArr[i] <= rightArr[j] {
			arr[k] = leftArr[i]
			i++
		} else {
			arr[k] = rightArr[j]
			j++
		}
		k++
	}
	
	// Copy remaining elements of leftArr, if any
	for i < leftSize {
		arr[k] = leftArr[i]
		i++
		k++
	}
	
	// Copy remaining elements of rightArr, if any
	for j < rightSize {
		arr[k] = rightArr[j]
		j++
		k++
	}
}

// MergeSortWithComparator sorts with custom comparison function
func MergeSortWithComparator(arr []int, less func(a, b int) bool) []int {
	if len(arr) <= 1 {
		return arr
	}
	
	result := make([]int, len(arr))
	copy(result, arr)
	
	mergeSortWithComparatorHelper(result, 0, len(result)-1, less)
	return result
}

// mergeSortWithComparatorHelper with custom comparator
func mergeSortWithComparatorHelper(arr []int, left, right int, less func(a, b int) bool) {
	if left >= right {
		return
	}
	
	mid := left + (right-left)/2
	
	mergeSortWithComparatorHelper(arr, left, mid, less)
	mergeSortWithComparatorHelper(arr, mid+1, right, less)
	
	mergeWithComparator(arr, left, mid, right, less)
}

// mergeWithComparator with custom comparison
func mergeWithComparator(arr []int, left, mid, right int, less func(a, b int) bool) {
	leftSize := mid - left + 1
	rightSize := right - mid
	
	leftArr := make([]int, leftSize)
	rightArr := make([]int, rightSize)
	
	copy(leftArr, arr[left:mid+1])
	copy(rightArr, arr[mid+1:right+1])
	
	i, j, k := 0, 0, left
	
	for i < leftSize && j < rightSize {
		if less(leftArr[i], rightArr[j]) || leftArr[i] == rightArr[j] {
			arr[k] = leftArr[i]
			i++
		} else {
			arr[k] = rightArr[j]
			j++
		}
		k++
	}
	
	for i < leftSize {
		arr[k] = leftArr[i]
		i++
		k++
	}
	
	for j < rightSize {
		arr[k] = rightArr[j]
		j++
		k++
	}
}

// BottomUpMergeSort iterative implementation
func BottomUpMergeSort(arr []int) []int {
	n := len(arr)
	if n <= 1 {
		return arr
	}
	
	result := make([]int, len(arr))
	copy(result, arr)
	
	// Merge subarrays of size 1, 2, 4, 8, ...
	for size := 1; size < n; size *= 2 {
		for left := 0; left < n-1; left += 2 * size {
			mid := min(left+size-1, n-1)
			right := min(left+2*size-1, n-1)
			
			if mid < right {
				merge(result, left, mid, right)
			}
		}
	}
	
	return result
}

// MergeSortedArrays merges two already sorted arrays
func MergeSortedArrays(arr1, arr2 []int) []int {
	result := make([]int, len(arr1)+len(arr2))
	i, j, k := 0, 0, 0
	
	// Merge elements from both arrays
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] <= arr2[j] {
			result[k] = arr1[i]
			i++
		} else {
			result[k] = arr2[j]
			j++
		}
		k++
	}
	
	// Copy remaining elements
	for i < len(arr1) {
		result[k] = arr1[i]
		i++
		k++
	}
	
	for j < len(arr2) {
		result[k] = arr2[j]
		j++
		k++
	}
	
	return result
}

// CountInversions counts the number of inversions using merge sort
func CountInversions(arr []int) int {
	if len(arr) <= 1 {
		return 0
	}
	
	temp := make([]int, len(arr))
	copy(temp, arr)
	
	return countInversionsHelper(temp, 0, len(temp)-1)
}

// countInversionsHelper counts inversions recursively
func countInversionsHelper(arr []int, left, right int) int {
	if left >= right {
		return 0
	}
	
	mid := left + (right-left)/2
	
	inversions := countInversionsHelper(arr, left, mid)
	inversions += countInversionsHelper(arr, mid+1, right)
	inversions += mergeAndCount(arr, left, mid, right)
	
	return inversions
}

// mergeAndCount merges and counts cross inversions
func mergeAndCount(arr []int, left, mid, right int) int {
	leftArr := make([]int, mid-left+1)
	rightArr := make([]int, right-mid)
	
	copy(leftArr, arr[left:mid+1])
	copy(rightArr, arr[mid+1:right+1])
	
	i, j, k := 0, 0, left
	inversions := 0
	
	for i < len(leftArr) && j < len(rightArr) {
		if leftArr[i] <= rightArr[j] {
			arr[k] = leftArr[i]
			i++
		} else {
			arr[k] = rightArr[j]
			// All elements from leftArr[i:] are greater than rightArr[j]
			inversions += len(leftArr) - i
			j++
		}
		k++
	}
	
	for i < len(leftArr) {
		arr[k] = leftArr[i]
		i++
		k++
	}
	
	for j < len(rightArr) {
		arr[k] = rightArr[j]
		j++
		k++
	}
	
	return inversions
}

// KWayMerge merges k sorted arrays
func KWayMerge(arrays [][]int) []int {
	if len(arrays) == 0 {
		return []int{}
	}
	
	result := arrays[0]
	
	for i := 1; i < len(arrays); i++ {
		result = MergeSortedArrays(result, arrays[i])
	}
	
	return result
}

// MergeSortRange sorts a specific range in the array
func MergeSortRange(arr []int, start, end int) {
	if start < 0 || end >= len(arr) || start >= end {
		return
	}
	
	mergeSortHelper(arr, start, end)
}

// IsSorted checks if array is sorted
func IsSorted(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			return false
		}
	}
	return true
}

// MergeSortStats returns sorting statistics
type MergeSortStats struct {
	Comparisons int
	Swaps       int
	ArraySize   int
	IsSorted    bool
}

// MergeSortWithStats sorts and returns statistics
func MergeSortWithStats(arr []int) ([]int, MergeSortStats) {
	stats := MergeSortStats{
		ArraySize: len(arr),
		IsSorted:  IsSorted(arr),
	}
	
	if len(arr) <= 1 {
		return arr, stats
	}
	
	result := make([]int, len(arr))
	copy(result, arr)
	
	mergeSortWithStats(result, 0, len(result)-1, &stats)
	
	return result, stats
}

// mergeSortWithStats with statistics tracking
func mergeSortWithStats(arr []int, left, right int, stats *MergeSortStats) {
	if left >= right {
		return
	}
	
	mid := left + (right-left)/2
	
	mergeSortWithStats(arr, left, mid, stats)
	mergeSortWithStats(arr, mid+1, right, stats)
	mergeWithStats(arr, left, mid, right, stats)
}

// mergeWithStats with statistics tracking
func mergeWithStats(arr []int, left, mid, right int, stats *MergeSortStats) {
	leftSize := mid - left + 1
	rightSize := right - mid
	
	leftArr := make([]int, leftSize)
	rightArr := make([]int, rightSize)
	
	copy(leftArr, arr[left:mid+1])
	copy(rightArr, arr[mid+1:right+1])
	
	i, j, k := 0, 0, left
	
	for i < leftSize && j < rightSize {
		stats.Comparisons++
		if leftArr[i] <= rightArr[j] {
			arr[k] = leftArr[i]
			i++
		} else {
			arr[k] = rightArr[j]
			j++
		}
		stats.Swaps++
		k++
	}
	
	for i < leftSize {
		arr[k] = leftArr[i]
		stats.Swaps++
		i++
		k++
	}
	
	for j < rightSize {
		arr[k] = rightArr[j]
		stats.Swaps++
		j++
		k++
	}
}

// Helper function
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
