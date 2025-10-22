package data_structures

import (
	"reflect"
	"testing"
)

func TestNewMaxHeap(t *testing.T) {
	heap := NewMaxHeap()
	if !heap.isMaxHeap {
		t.Error("Expected max heap, got min heap")
	}
	if heap.Size() != 0 {
		t.Errorf("Expected size 0, got %d", heap.Size())
	}
}

func TestNewMinHeap(t *testing.T) {
	heap := NewMinHeap()
	if heap.isMaxHeap {
		t.Error("Expected min heap, got max heap")
	}
	if heap.Size() != 0 {
		t.Errorf("Expected size 0, got %d", heap.Size())
	}
}

func TestMaxHeapInsertAndExtract(t *testing.T) {
	heap := NewMaxHeap()
	
	// Test insertions
	values := []int{10, 5, 15, 3, 7, 12, 20}
	for _, v := range values {
		heap.Insert(v)
	}
	
	if heap.Size() != len(values) {
		t.Errorf("Expected size %d, got %d", len(values), heap.Size())
	}
	
	// Check if heap property is maintained
	if !heap.Validate() {
		t.Error("Heap property violated after insertions")
	}
	
	// Test extractions (should come out in descending order)
	expected := []int{20, 15, 12, 10, 7, 5, 3}
	for i, expectedVal := range expected {
		val, err := heap.Extract()
		if err != nil {
			t.Errorf("Unexpected error during extraction %d: %v", i, err)
		}
		if val != expectedVal {
			t.Errorf("Expected %d, got %d at position %d", expectedVal, val, i)
		}
	}
	
	if !heap.IsEmpty() {
		t.Error("Heap should be empty after all extractions")
	}
}

func TestMinHeapInsertAndExtract(t *testing.T) {
	heap := NewMinHeap()
	
	values := []int{10, 5, 15, 3, 7, 12, 20}
	for _, v := range values {
		heap.Insert(v)
	}
	
	// Test extractions (should come out in ascending order)
	expected := []int{3, 5, 7, 10, 12, 15, 20}
	for i, expectedVal := range expected {
		val, err := heap.Extract()
		if err != nil {
			t.Errorf("Unexpected error during extraction %d: %v", i, err)
		}
		if val != expectedVal {
			t.Errorf("Expected %d, got %d at position %d", expectedVal, val, i)
		}
	}
}

func TestPeek(t *testing.T) {
	// Test max heap
	maxHeap := NewMaxHeap()
	values := []int{10, 5, 15, 3, 7}
	for _, v := range values {
		maxHeap.Insert(v)
	}
	
	max, err := maxHeap.Peek()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if max != 15 {
		t.Errorf("Expected max 15, got %d", max)
	}
	
	// Peek should not remove element
	if maxHeap.Size() != 5 {
		t.Errorf("Size should remain 5 after peek, got %d", maxHeap.Size())
	}
	
	// Test min heap
	minHeap := NewMinHeap()
	for _, v := range values {
		minHeap.Insert(v)
	}
	
	min, err := minHeap.Peek()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if min != 3 {
		t.Errorf("Expected min 3, got %d", min)
	}
}

func TestEmptyHeapOperations(t *testing.T) {
	heap := NewMaxHeap()
	
	// Test peek on empty heap
	_, err := heap.Peek()
	if err == nil {
		t.Error("Expected error when peeking empty heap")
	}
	
	// Test extract on empty heap
	_, err = heap.Extract()
	if err == nil {
		t.Error("Expected error when extracting from empty heap")
	}
	
	if !heap.IsEmpty() {
		t.Error("Heap should be empty")
	}
}

func TestBuildHeap(t *testing.T) {
	heap := NewMaxHeap()
	items := []int{4, 10, 3, 5, 1, 6, 9, 7, 8, 2}
	
	heap.BuildHeap(items)
	
	if heap.Size() != len(items) {
		t.Errorf("Expected size %d, got %d", len(items), heap.Size())
	}
	
	if !heap.Validate() {
		t.Error("Heap property violated after building heap")
	}
	
	// Extract all elements and verify they come out in sorted order
	var extracted []int
	for !heap.IsEmpty() {
		val, _ := heap.Extract()
		extracted = append(extracted, val)
	}
	
	// Should be in descending order for max heap
	for i := 1; i < len(extracted); i++ {
		if extracted[i] > extracted[i-1] {
			t.Errorf("Elements not in descending order: %v", extracted)
			break
		}
	}
}

func TestHeapSort(t *testing.T) {
	heap := NewMaxHeap()
	items := []int{64, 34, 25, 12, 22, 11, 90}
	expected := []int{11, 12, 22, 25, 34, 64, 90}
	
	sorted := heap.HeapSort(items)
	
	if !reflect.DeepEqual(sorted, expected) {
		t.Errorf("Expected %v, got %v", expected, sorted)
	}
	
	// Test with empty array
	empty := []int{}
	sortedEmpty := heap.HeapSort(empty)
	if len(sortedEmpty) != 0 {
		t.Error("Sorted empty array should be empty")
	}
	
	// Test with single element
	single := []int{42}
	sortedSingle := heap.HeapSort(single)
	if !reflect.DeepEqual(sortedSingle, single) {
		t.Errorf("Expected %v, got %v", single, sortedSingle)
	}
}

func TestGetKthLargest(t *testing.T) {
	heap := NewMaxHeap()
	items := []int{3, 2, 1, 5, 6, 4}
	
	// Test various k values
	tests := []struct {
		k        int
		expected int
	}{
		{1, 6}, // 1st largest
		{2, 5}, // 2nd largest
		{3, 4}, // 3rd largest
		{6, 1}, // 6th largest (smallest)
	}
	
	for _, test := range tests {
		result, err := heap.GetKthLargest(items, test.k)
		if err != nil {
			t.Errorf("Unexpected error for k=%d: %v", test.k, err)
		}
		if result != test.expected {
			t.Errorf("For k=%d, expected %d, got %d", test.k, test.expected, result)
		}
	}
	
	// Test invalid k values
	_, err := heap.GetKthLargest(items, 0)
	if err == nil {
		t.Error("Expected error for k=0")
	}
	
	_, err = heap.GetKthLargest(items, 7)
	if err == nil {
		t.Error("Expected error for k=7 (greater than array size)")
	}
}

func BenchmarkHeapInsert(b *testing.B) {
	heap := NewMaxHeap()
	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		heap.Insert(i)
	}
}

func BenchmarkHeapExtract(b *testing.B) {
	heap := NewMaxHeap()
	// Pre-fill heap
	for i := 0; i < b.N; i++ {
		heap.Insert(i)
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		heap.Extract()
	}
}

func BenchmarkHeapSort(b *testing.B) {
	heap := NewMaxHeap()
	items := make([]int, 1000)
	for i := range items {
		items[i] = 1000 - i // Reverse order for worst case
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		heap.HeapSort(items)
	}
}
