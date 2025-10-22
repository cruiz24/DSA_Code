package data_structures

import (
	"errors"
	"fmt"
)

// BinaryHeap represents a binary heap data structure
type BinaryHeap struct {
	items   []int
	isMaxHeap bool
}

// NewMaxHeap creates a new max heap
func NewMaxHeap() *BinaryHeap {
	return &BinaryHeap{
		items:   make([]int, 0),
		isMaxHeap: true,
	}
}

// NewMinHeap creates a new min heap
func NewMinHeap() *BinaryHeap {
	return &BinaryHeap{
		items:   make([]int, 0),
		isMaxHeap: false,
	}
}

// Insert adds a new element to the heap
func (h *BinaryHeap) Insert(item int) {
	h.items = append(h.items, item)
	h.heapifyUp(len(h.items) - 1)
}

// Extract removes and returns the root element
func (h *BinaryHeap) Extract() (int, error) {
	if len(h.items) == 0 {
		return 0, errors.New("heap is empty")
	}

	root := h.items[0]
	lastIndex := len(h.items) - 1
	h.items[0] = h.items[lastIndex]
	h.items = h.items[:lastIndex]
	
	if len(h.items) > 0 {
		h.heapifyDown(0)
	}
	
	return root, nil
}

// Peek returns the root element without removing it
func (h *BinaryHeap) Peek() (int, error) {
	if len(h.items) == 0 {
		return 0, errors.New("heap is empty")
	}
	return h.items[0], nil
}

// Size returns the number of elements in the heap
func (h *BinaryHeap) Size() int {
	return len(h.items)
}

// IsEmpty checks if the heap is empty
func (h *BinaryHeap) IsEmpty() bool {
	return len(h.items) == 0
}

// BuildHeap creates a heap from an existing slice
func (h *BinaryHeap) BuildHeap(items []int) {
	h.items = make([]int, len(items))
	copy(h.items, items)
	
	// Start from the last non-leaf node and heapify down
	for i := (len(h.items) / 2) - 1; i >= 0; i-- {
		h.heapifyDown(i)
	}
}

// HeapSort sorts the array using heap sort algorithm
func (h *BinaryHeap) HeapSort(items []int) []int {
	if len(items) == 0 {
		return items
	}
	
	// Create a max heap for ascending sort
	heap := NewMaxHeap()
	heap.BuildHeap(items)
	
	sorted := make([]int, len(items))
	originalSize := len(heap.items)
	
	for i := originalSize - 1; i >= 0; i-- {
		// Extract max and place at the end
		max, _ := heap.Extract()
		sorted[i] = max
	}
	
	return sorted
}

// GetKthLargest finds the kth largest element
func (h *BinaryHeap) GetKthLargest(items []int, k int) (int, error) {
	if k <= 0 || k > len(items) {
		return 0, errors.New("invalid k value")
	}
	
	// Use min heap to find kth largest
	minHeap := NewMinHeap()
	
	for _, item := range items {
		if minHeap.Size() < k {
			minHeap.Insert(item)
		} else if peek, _ := minHeap.Peek(); item > peek {
			minHeap.Extract()
			minHeap.Insert(item)
		}
	}
	
	return minHeap.Peek()
}

// heapifyUp maintains heap property by moving element up
func (h *BinaryHeap) heapifyUp(index int) {
	parentIndex := (index - 1) / 2
	
	if index > 0 && h.shouldSwap(index, parentIndex) {
		h.items[index], h.items[parentIndex] = h.items[parentIndex], h.items[index]
		h.heapifyUp(parentIndex)
	}
}

// heapifyDown maintains heap property by moving element down
func (h *BinaryHeap) heapifyDown(index int) {
	leftChild := 2*index + 1
	rightChild := 2*index + 2
	targetIndex := index
	
	if leftChild < len(h.items) && h.shouldSwap(leftChild, targetIndex) {
		targetIndex = leftChild
	}
	
	if rightChild < len(h.items) && h.shouldSwap(rightChild, targetIndex) {
		targetIndex = rightChild
	}
	
	if targetIndex != index {
		h.items[index], h.items[targetIndex] = h.items[targetIndex], h.items[index]
		h.heapifyDown(targetIndex)
	}
}

// shouldSwap determines if two elements should be swapped based on heap type
func (h *BinaryHeap) shouldSwap(childIndex, parentIndex int) bool {
	if h.isMaxHeap {
		return h.items[childIndex] > h.items[parentIndex]
	}
	return h.items[childIndex] < h.items[parentIndex]
}

// Print displays the heap structure
func (h *BinaryHeap) Print() {
	if len(h.items) == 0 {
		fmt.Println("Empty heap")
		return
	}
	
	heapType := "Max"
	if !h.isMaxHeap {
		heapType = "Min"
	}
	
	fmt.Printf("%s Heap: %v\n", heapType, h.items)
}

// Validate checks if the heap property is maintained
func (h *BinaryHeap) Validate() bool {
	for i := 0; i < len(h.items); i++ {
		leftChild := 2*i + 1
		rightChild := 2*i + 2
		
		if leftChild < len(h.items) {
			if h.isMaxHeap && h.items[i] < h.items[leftChild] {
				return false
			}
			if !h.isMaxHeap && h.items[i] > h.items[leftChild] {
				return false
			}
		}
		
		if rightChild < len(h.items) {
			if h.isMaxHeap && h.items[i] < h.items[rightChild] {
				return false
			}
			if !h.isMaxHeap && h.items[i] > h.items[rightChild] {
				return false
			}
		}
	}
	return true
}
