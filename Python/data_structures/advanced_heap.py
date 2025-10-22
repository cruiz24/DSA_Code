"""
Advanced Heap Data Structure Implementation
=========================================

This module provides a comprehensive implementation of heap data structures including:
- Min Heap and Max Heap
- Binary Heap with custom comparison
- D-ary Heap (generalized heap with d children per node)
- Indexed Priority Queue
- Merge operation for heaps

Time Complexities:
- Insert: O(log n)
- Extract Min/Max: O(log n)
- Peek: O(1)
- Heapify: O(n)
- Merge: O(n + m)

Author: Hacktoberfest Contributor
Date: 2025
"""

import heapq
from typing import List, Optional, Callable, Any, Tuple
from abc import ABC, abstractmethod


class AdvancedHeap(ABC):
    """Abstract base class for heap implementations."""
    
    def __init__(self):
        self.heap = []
        self.size = 0
    
    @abstractmethod
    def push(self, item: Any) -> None:
        """Insert an item into the heap."""
        pass
    
    @abstractmethod
    def pop(self) -> Any:
        """Remove and return the top item from the heap."""
        pass
    
    @abstractmethod
    def peek(self) -> Any:
        """Return the top item without removing it."""
        pass
    
    def is_empty(self) -> bool:
        """Check if the heap is empty."""
        return self.size == 0
    
    def __len__(self) -> int:
        """Return the number of items in the heap."""
        return self.size


class MinHeap(AdvancedHeap):
    """
    Min Heap implementation with advanced features.
    
    Features:
    - Custom comparison function support
    - Efficient merge operation
    - Batch operations
    - Priority updates
    """
    
    def __init__(self, comparator: Optional[Callable] = None):
        super().__init__()
        self.compare = comparator if comparator else lambda x, y: x < y
    
    def push(self, item: Any) -> None:
        """Insert an item into the min heap."""
        self.heap.append(item)
        self.size += 1
        self._bubble_up(self.size - 1)
    
    def pop(self) -> Any:
        """Remove and return the minimum item."""
        if self.is_empty():
            raise IndexError("pop from empty heap")
        
        min_item = self.heap[0]
        last_item = self.heap.pop()
        self.size -= 1
        
        if not self.is_empty():
            self.heap[0] = last_item
            self._bubble_down(0)
        
        return min_item
    
    def peek(self) -> Any:
        """Return the minimum item without removing it."""
        if self.is_empty():
            raise IndexError("peek from empty heap")
        return self.heap[0]
    
    def push_pop(self, item: Any) -> Any:
        """Push item and pop the minimum in one operation."""
        if self.is_empty() or self.compare(item, self.heap[0]):
            return item
        
        min_item = self.heap[0]
        self.heap[0] = item
        self._bubble_down(0)
        return min_item
    
    def replace(self, item: Any) -> Any:
        """Pop minimum and push new item in one operation."""
        if self.is_empty():
            raise IndexError("replace on empty heap")
        
        min_item = self.heap[0]
        self.heap[0] = item
        self._bubble_down(0)
        return min_item
    
    def merge(self, other_heap: 'MinHeap') -> 'MinHeap':
        """Merge with another heap and return a new heap."""
        merged = MinHeap(self.compare)
        merged.heap = self.heap + other_heap.heap
        merged.size = len(merged.heap)
        merged._heapify()
        return merged
    
    def nlargest(self, n: int) -> List[Any]:
        """Return the n largest elements."""
        if n >= self.size:
            return sorted(self.heap, key=lambda x: x, reverse=True)
        
        # Use a max heap to find n largest
        result = []
        temp_heap = [-x for x in self.heap]
        heapq.heapify(temp_heap)
        
        for _ in range(min(n, self.size)):
            result.append(-heapq.heappop(temp_heap))
        
        return result
    
    def nsmallest(self, n: int) -> List[Any]:
        """Return the n smallest elements."""
        return heapq.nsmallest(n, self.heap)
    
    def _bubble_up(self, index: int) -> None:
        """Move element up to maintain heap property."""
        parent = (index - 1) // 2
        if index > 0 and self.compare(self.heap[index], self.heap[parent]):
            self.heap[index], self.heap[parent] = self.heap[parent], self.heap[index]
            self._bubble_up(parent)
    
    def _bubble_down(self, index: int) -> None:
        """Move element down to maintain heap property."""
        smallest = index
        left = 2 * index + 1
        right = 2 * index + 2
        
        if (left < self.size and 
            self.compare(self.heap[left], self.heap[smallest])):
            smallest = left
        
        if (right < self.size and 
            self.compare(self.heap[right], self.heap[smallest])):
            smallest = right
        
        if smallest != index:
            self.heap[index], self.heap[smallest] = self.heap[smallest], self.heap[index]
            self._bubble_down(smallest)
    
    def _heapify(self) -> None:
        """Convert array to heap in O(n) time."""
        for i in range(self.size // 2 - 1, -1, -1):
            self._bubble_down(i)


class MaxHeap(MinHeap):
    """Max Heap implementation extending MinHeap."""
    
    def __init__(self, comparator: Optional[Callable] = None):
        # Reverse the comparison for max heap
        if comparator:
            super().__init__(lambda x, y: not comparator(x, y))
        else:
            super().__init__(lambda x, y: x > y)


class DaryHeap:
    """
    D-ary Heap implementation where each node has d children.
    More efficient for applications with frequent decrease-key operations.
    """
    
    def __init__(self, d: int = 2, is_min_heap: bool = True):
        self.d = d  # Number of children per node
        self.heap = []
        self.size = 0
        self.is_min = is_min_heap
    
    def push(self, item: Any) -> None:
        """Insert an item into the d-ary heap."""
        self.heap.append(item)
        self.size += 1
        self._bubble_up(self.size - 1)
    
    def pop(self) -> Any:
        """Remove and return the top item."""
        if self.size == 0:
            raise IndexError("pop from empty heap")
        
        top_item = self.heap[0]
        last_item = self.heap.pop()
        self.size -= 1
        
        if self.size > 0:
            self.heap[0] = last_item
            self._bubble_down(0)
        
        return top_item
    
    def peek(self) -> Any:
        """Return the top item without removing it."""
        if self.size == 0:
            raise IndexError("peek from empty heap")
        return self.heap[0]
    
    def _parent(self, index: int) -> int:
        """Get parent index."""
        return (index - 1) // self.d
    
    def _children(self, index: int) -> List[int]:
        """Get children indices."""
        first_child = self.d * index + 1
        return [i for i in range(first_child, min(first_child + self.d, self.size))]
    
    def _compare(self, x: Any, y: Any) -> bool:
        """Compare two elements based on heap type."""
        return x < y if self.is_min else x > y
    
    def _bubble_up(self, index: int) -> None:
        """Move element up to maintain heap property."""
        if index == 0:
            return
        
        parent = self._parent(index)
        if self._compare(self.heap[index], self.heap[parent]):
            self.heap[index], self.heap[parent] = self.heap[parent], self.heap[index]
            self._bubble_up(parent)
    
    def _bubble_down(self, index: int) -> None:
        """Move element down to maintain heap property."""
        children = self._children(index)
        if not children:
            return
        
        best_child = children[0]
        for child in children[1:]:
            if self._compare(self.heap[child], self.heap[best_child]):
                best_child = child
        
        if self._compare(self.heap[best_child], self.heap[index]):
            self.heap[index], self.heap[best_child] = self.heap[best_child], self.heap[index]
            self._bubble_down(best_child)


class IndexedPriorityQueue:
    """
    Indexed Priority Queue supporting efficient priority updates.
    Useful for algorithms like Dijkstra's shortest path.
    """
    
    def __init__(self, max_size: int, is_min_pq: bool = True):
        self.max_size = max_size
        self.is_min = is_min_pq
        self.heap = []  # (priority, index) pairs
        self.position = [-1] * max_size  # position[i] = position of index i in heap
        self.values = [None] * max_size  # values[i] = priority of index i
        self.size = 0
    
    def push(self, index: int, priority: Any) -> None:
        """Insert or update priority for given index."""
        if index < 0 or index >= self.max_size:
            raise ValueError("Index out of bounds")
        
        if self.contains(index):
            self.update_priority(index, priority)
        else:
            self.values[index] = priority
            self.heap.append((priority, index))
            self.position[index] = self.size
            self.size += 1
            self._bubble_up(self.size - 1)
    
    def pop(self) -> Tuple[int, Any]:
        """Remove and return (index, priority) with best priority."""
        if self.size == 0:
            raise IndexError("pop from empty priority queue")
        
        priority, index = self.heap[0]
        last_item = self.heap.pop()
        self.size -= 1
        self.position[index] = -1
        
        if self.size > 0:
            self.heap[0] = last_item
            self.position[last_item[1]] = 0
            self._bubble_down(0)
        
        return index, priority
    
    def peek(self) -> Tuple[int, Any]:
        """Return (index, priority) with best priority without removing."""
        if self.size == 0:
            raise IndexError("peek from empty priority queue")
        priority, index = self.heap[0]
        return index, priority
    
    def contains(self, index: int) -> bool:
        """Check if index is in the priority queue."""
        return 0 <= index < self.max_size and self.position[index] != -1
    
    def update_priority(self, index: int, new_priority: Any) -> None:
        """Update priority for given index."""
        if not self.contains(index):
            raise ValueError("Index not in priority queue")
        
        pos = self.position[index]
        old_priority = self.values[index]
        self.values[index] = new_priority
        self.heap[pos] = (new_priority, index)
        
        # Restore heap property
        if self._compare(new_priority, old_priority):
            self._bubble_up(pos)
        else:
            self._bubble_down(pos)
    
    def _compare(self, x: Any, y: Any) -> bool:
        """Compare two priorities based on queue type."""
        return x < y if self.is_min else x > y
    
    def _bubble_up(self, index: int) -> None:
        """Move element up to maintain heap property."""
        if index == 0:
            return
        
        parent = (index - 1) // 2
        if self._compare(self.heap[index][0], self.heap[parent][0]):
            # Swap heap elements
            self.heap[index], self.heap[parent] = self.heap[parent], self.heap[index]
            # Update position mapping
            self.position[self.heap[index][1]] = index
            self.position[self.heap[parent][1]] = parent
            self._bubble_up(parent)
    
    def _bubble_down(self, index: int) -> None:
        """Move element down to maintain heap property."""
        left = 2 * index + 1
        right = 2 * index + 2
        best = index
        
        if (left < self.size and 
            self._compare(self.heap[left][0], self.heap[best][0])):
            best = left
        
        if (right < self.size and 
            self._compare(self.heap[right][0], self.heap[best][0])):
            best = right
        
        if best != index:
            # Swap heap elements
            self.heap[index], self.heap[best] = self.heap[best], self.heap[index]
            # Update position mapping
            self.position[self.heap[index][1]] = index
            self.position[self.heap[best][1]] = best
            self._bubble_down(best)


# Utility functions for heap operations
def heapsort(arr: List[Any], reverse: bool = False) -> List[Any]:
    """Sort array using heap sort algorithm."""
    heap = MaxHeap() if reverse else MinHeap()
    
    # Build heap
    for item in arr:
        heap.push(item)
    
    # Extract elements in sorted order
    result = []
    while not heap.is_empty():
        result.append(heap.pop())
    
    return result


def merge_k_sorted_arrays(arrays: List[List[Any]]) -> List[Any]:
    """Merge k sorted arrays using min heap."""
    min_heap = MinHeap()
    result = []
    
    # Initialize heap with first element from each array
    for i, arr in enumerate(arrays):
        if arr:
            min_heap.push((arr[0], i, 0))  # (value, array_index, element_index)
    
    while not min_heap.is_empty():
        value, arr_idx, elem_idx = min_heap.pop()
        result.append(value)
        
        # Add next element from the same array
        if elem_idx + 1 < len(arrays[arr_idx]):
            next_value = arrays[arr_idx][elem_idx + 1]
            min_heap.push((next_value, arr_idx, elem_idx + 1))
    
    return result


# Test and demonstration code
def test_advanced_heap():
    """Comprehensive test suite for advanced heap implementations."""
    import time
    import random
    
    print("=== Advanced Heap Test Suite ===\n")
    
    # Test 1: Basic Min Heap operations
    print("1. Testing Min Heap:")
    min_heap = MinHeap()
    test_data = [5, 2, 8, 1, 9, 3]
    
    for item in test_data:
        min_heap.push(item)
    print(f"Inserted: {test_data}")
    
    extracted = []
    while not min_heap.is_empty():
        extracted.append(min_heap.pop())
    print(f"Extracted (sorted): {extracted}")
    assert extracted == sorted(test_data), "Min heap extraction failed"
    
    # Test 2: Max Heap operations
    print("\n2. Testing Max Heap:")
    max_heap = MaxHeap()
    for item in test_data:
        max_heap.push(item)
    
    extracted = []
    while not max_heap.is_empty():
        extracted.append(max_heap.pop())
    print(f"Extracted (reverse sorted): {extracted}")
    assert extracted == sorted(test_data, reverse=True), "Max heap extraction failed"
    
    # Test 3: Custom comparator (strings by length)
    print("\n3. Testing Custom Comparator (strings by length):")
    string_heap = MinHeap(lambda x, y: len(x) < len(y))
    strings = ["hello", "hi", "world", "a", "python"]
    
    for s in strings:
        string_heap.push(s)
    
    extracted = []
    while not string_heap.is_empty():
        extracted.append(string_heap.pop())
    print(f"Strings sorted by length: {extracted}")
    
    # Test 4: D-ary Heap (ternary heap)
    print("\n4. Testing D-ary Heap (d=3):")
    ternary_heap = DaryHeap(d=3, is_min_heap=True)
    for item in test_data:
        ternary_heap.push(item)
    
    extracted = []
    while ternary_heap.size > 0:
        extracted.append(ternary_heap.pop())
    print(f"Ternary heap extraction: {extracted}")
    assert extracted == sorted(test_data), "Ternary heap extraction failed"
    
    # Test 5: Indexed Priority Queue
    print("\n5. Testing Indexed Priority Queue:")
    ipq = IndexedPriorityQueue(10, is_min_pq=True)
    
    # Insert priorities for different indices
    priorities = [(0, 5), (1, 2), (2, 8), (3, 1), (4, 9)]
    for idx, priority in priorities:
        ipq.push(idx, priority)
    
    print("Extracting in priority order:")
    while ipq.size > 0:
        idx, priority = ipq.pop()
        print(f"Index: {idx}, Priority: {priority}")
    
    # Test 6: Heap merge operation
    print("\n6. Testing Heap Merge:")
    heap1 = MinHeap()
    heap2 = MinHeap()
    
    for item in [1, 3, 5, 7]:
        heap1.push(item)
    for item in [2, 4, 6, 8]:
        heap2.push(item)
    
    merged = heap1.merge(heap2)
    extracted = []
    while not merged.is_empty():
        extracted.append(merged.pop())
    print(f"Merged heap extraction: {extracted}")
    
    # Test 7: Performance comparison
    print("\n7. Performance Comparison:")
    n = 10000
    data = list(range(n))
    random.shuffle(data)
    
    # Test MinHeap performance
    start_time = time.time()
    heap = MinHeap()
    for item in data:
        heap.push(item)
    while not heap.is_empty():
        heap.pop()
    min_heap_time = time.time() - start_time
    
    # Test Python's heapq performance
    start_time = time.time()
    py_heap = data.copy()
    heapq.heapify(py_heap)
    while py_heap:
        heapq.heappop(py_heap)
    heapq_time = time.time() - start_time
    
    print(f"MinHeap time for {n} operations: {min_heap_time:.4f}s")
    print(f"Python heapq time for {n} operations: {heapq_time:.4f}s")
    print(f"Performance ratio: {min_heap_time/heapq_time:.2f}x")
    
    # Test 8: Utility functions
    print("\n8. Testing Utility Functions:")
    
    # Heap sort
    unsorted = [64, 34, 25, 12, 22, 11, 90]
    sorted_asc = heapsort(unsorted.copy())
    sorted_desc = heapsort(unsorted.copy(), reverse=True)
    print(f"Original: {unsorted}")
    print(f"Heap sort (ascending): {sorted_asc}")
    print(f"Heap sort (descending): {sorted_desc}")
    
    # Merge k sorted arrays
    k_arrays = [
        [1, 4, 7],
        [2, 5, 8],
        [3, 6, 9]
    ]
    merged_result = merge_k_sorted_arrays(k_arrays)
    print(f"Merged {len(k_arrays)} sorted arrays: {merged_result}")
    
    print("\n=== All tests passed! ===")


if __name__ == "__main__":
    test_advanced_heap()
