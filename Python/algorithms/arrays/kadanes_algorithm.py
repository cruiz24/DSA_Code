"""
Kadane's Algorithm - Maximum Subarray Sum
==========================================

Algorithm Description:
---------------------
Kadane's Algorithm finds the maximum sum of a contiguous subarray within a one-dimensional 
numeric array. It uses dynamic programming to efficiently solve this problem in linear time.

The key insight is that at any position, the maximum subarray ending at that position is 
either the element itself or the element plus the maximum subarray ending at the previous position.

Time Complexity: O(n) - Single pass through the array
Space Complexity: O(1) - Only constant extra space used

Algorithm Steps:
1. Initialize max_current and max_global with the first element
2. Iterate through array from index 1 to n-1
3. At each position, decide whether to extend the existing subarray or start new
4. Update max_current = max(arr[i], max_current + arr[i])
5. Update max_global if max_current is greater
6. Return max_global

Use Cases:
- Stock profit maximization (buy/sell once)
- Finding best performing period in time series data
- Image processing for finding regions of interest
- Network flow optimization

Example:
--------
Input:  [-2, 1, -3, 4, -1, 2, 1, -5, 4]
Output: 6
Explanation: Subarray [4, -1, 2, 1] has the maximum sum of 6

Author: [Your Name]
Date: October 2024
Hacktoberfest 2024
"""

def kadanes_algorithm(arr):
    """
    Find maximum sum of contiguous subarray using Kadane's Algorithm.
    
    Args:
        arr (list): List of integers (can contain negative numbers)
        
    Returns:
        int: Maximum sum of contiguous subarray
        
    Raises:
        ValueError: If array is empty or None
        TypeError: If input is not a list
    """
    # Input validation
    if arr is None:
        raise ValueError("Array cannot be None")
    
    if not isinstance(arr, list):
        raise TypeError("Input must be a list")
    
    if len(arr) == 0:
        raise ValueError("Array cannot be empty")
    
    # Initialize variables
    max_current = arr[0]  # Maximum sum ending at current position
    max_global = arr[0]   # Maximum sum found so far
    
    # Track the subarray indices (optional, for returning the actual subarray)
    start = 0
    end = 0
    temp_start = 0
    
    # Iterate through array starting from second element
    for i in range(1, len(arr)):
        # Decide: extend existing subarray or start new one
        if arr[i] > max_current + arr[i]:
            max_current = arr[i]
            temp_start = i
        else:
            max_current = max_current + arr[i]
        
        # Update global maximum if current is better
        if max_current > max_global:
            max_global = max_current
            start = temp_start
            end = i
    
    return max_global


def kadanes_algorithm_with_subarray(arr):
    """
    Find maximum sum and return the actual subarray.
    
    Args:
        arr (list): List of integers
        
    Returns:
        tuple: (max_sum, start_index, end_index, subarray)
    """
    if not arr:
        raise ValueError("Array cannot be empty")
    
    max_current = arr[0]
    max_global = arr[0]
    start = 0
    end = 0
    temp_start = 0
    
    for i in range(1, len(arr)):
        if arr[i] > max_current + arr[i]:
            max_current = arr[i]
            temp_start = i
        else:
            max_current = max_current + arr[i]
        
        if max_current > max_global:
            max_global = max_current
            start = temp_start
            end = i
    
    subarray = arr[start:end + 1]
    return max_global, start, end, subarray


def kadanes_algorithm_all_negative(arr):
    """
    Variation that handles all-negative arrays by returning least negative.
    
    Args:
        arr (list): List of integers
        
    Returns:
        int: Maximum sum (least negative if all elements negative)
    """
    if not arr:
        raise ValueError("Array cannot be empty")
    
    max_current = arr[0]
    max_global = arr[0]
    
    for i in range(1, len(arr)):
        max_current = max(arr[i], max_current + arr[i])
        max_global = max(max_global, max_current)
    
    return max_global


# Test Cases
# ==========

def test_kadanes_algorithm():
    """
    Comprehensive test cases for Kadane's Algorithm.
    """
    print("=" * 60)
    print("Testing Kadane's Algorithm - Maximum Subarray Sum")
    print("=" * 60)
    
    # Test Case 1: Mixed positive and negative numbers
    test1 = [-2, 1, -3, 4, -1, 2, 1, -5, 4]
    result1 = kadanes_algorithm(test1)
    print(f"\nTest 1: Mixed numbers")
    print(f"Input:  {test1}")
    print(f"Output: {result1}")
    print(f"Expected: 6 (subarray [4, -1, 2, 1])")
    assert result1 == 6, "Test 1 failed"
    print("✅ Passed")
    
    # Test Case 2: All positive numbers
    test2 = [1, 2, 3, 4, 5]
    result2 = kadanes_algorithm(test2)
    print(f"\nTest 2: All positive")
    print(f"Input:  {test2}")
    print(f"Output: {result2}")
    print(f"Expected: 15 (entire array)")
    assert result2 == 15, "Test 2 failed"
    print("✅ Passed")
    
    # Test Case 3: All negative numbers
    test3 = [-5, -2, -8, -1, -4]
    result3 = kadanes_algorithm(test3)
    print(f"\nTest 3: All negative")
    print(f"Input:  {test3}")
    print(f"Output: {result3}")
    print(f"Expected: -1 (least negative)")
    assert result3 == -1, "Test 3 failed"
    print("✅ Passed")
    
    # Test Case 4: Single element
    test4 = [5]
    result4 = kadanes_algorithm(test4)
    print(f"\nTest 4: Single element")
    print(f"Input:  {test4}")
    print(f"Output: {result4}")
    print(f"Expected: 5")
    assert result4 == 5, "Test 4 failed"
    print("✅ Passed")
    
    # Test Case 5: Large numbers
    test5 = [-2, -3, 4, -1, -2, 1, 5, -3]
    result5 = kadanes_algorithm(test5)
    print(f"\nTest 5: Large mixed numbers")
    print(f"Input:  {test5}")
    print(f"Output: {result5}")
    print(f"Expected: 7 (subarray [4, -1, -2, 1, 5])")
    assert result5 == 7, "Test 5 failed"
    print("✅ Passed")
    
    # Test Case 6: With subarray information
    test6 = [-2, 1, -3, 4, -1, 2, 1, -5, 4]
    max_sum, start, end, subarray = kadanes_algorithm_with_subarray(test6)
    print(f"\nTest 6: With subarray details")
    print(f"Input:  {test6}")
    print(f"Max Sum: {max_sum}")
    print(f"Subarray: {subarray}")
    print(f"Indices: [{start}:{end}]")
    assert max_sum == 6 and subarray == [4, -1, 2, 1], "Test 6 failed"
    print("✅ Passed")
    
    # Test Case 7: Edge case - two elements
    test7 = [-1, 2]
    result7 = kadanes_algorithm(test7)
    print(f"\nTest 7: Two elements")
    print(f"Input:  {test7}")
    print(f"Output: {result7}")
    print(f"Expected: 2")
    assert result7 == 2, "Test 7 failed"
    print("✅ Passed")
    
    # Test Case 8: Alternating positive and negative
    test8 = [5, -3, 5, -3, 5]
    result8 = kadanes_algorithm(test8)
    print(f"\nTest 8: Alternating")
    print(f"Input:  {test8}")
    print(f"Output: {result8}")
    print(f"Expected: 9 (entire array)")
    assert result8 == 9, "Test 8 failed"
    print("✅ Passed")
    
    print("\n" + "=" * 60)
    print("All Tests Passed! ✅")
    print("=" * 60)


def test_edge_cases():
    """
    Test edge cases and error handling.
    """
    print("\n" + "=" * 60)
    print("Testing Edge Cases")
    print("=" * 60)
    
    # Test empty array
    try:
        kadanes_algorithm([])
        print("❌ Empty array test failed - should raise ValueError")
    except ValueError as e:
        print(f"\n✅ Empty array handled correctly: {e}")
    
    # Test None input
    try:
        kadanes_algorithm(None)
        print("❌ None input test failed - should raise ValueError")
    except ValueError as e:
        print(f"✅ None input handled correctly: {e}")
    
    # Test non-list input
    try:
        kadanes_algorithm("not a list")
        print("❌ Non-list input test failed - should raise TypeError")
    except TypeError as e:
        print(f"✅ Non-list input handled correctly: {e}")
    
    print("\n" + "=" * 60)
    print("Edge Cases Handled Successfully! ✅")
    print("=" * 60)


# Main execution
if __name__ == "__main__":
    # Run all tests
    test_kadanes_algorithm()
    test_edge_cases()
    
    # Interactive example
    print("\n" + "=" * 60)
    print("Interactive Example")
    print("=" * 60)
    
    example_array = [-2, 1, -3, 4, -1, 2, 1, -5, 4]
    print(f"\nFinding maximum subarray sum for: {example_array}")
    
    max_sum, start, end, subarray = kadanes_algorithm_with_subarray(example_array)
    
    print(f"\n📊 Results:")
    print(f"   Maximum Sum: {max_sum}")
    print(f"   Subarray: {subarray}")
    print(f"   Start Index: {start}")
    print(f"   End Index: {end}")
    print(f"\n✨ The subarray {subarray} has the maximum sum of {max_sum}")
    
    print("\n" + "=" * 60)
    print("Kadane's Algorithm Demo Complete! 🎉")
    print("=" * 60)