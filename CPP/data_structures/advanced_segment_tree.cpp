/**
 * Advanced Segment Tree Implementation in C++
 * Supports range queries and updates with lazy propagation
 * 
 * Time Complexities:
 * - Build: O(n)
 * - Point Update: O(log n)
 * - Range Update: O(log n) with lazy propagation
 * - Range Query: O(log n)
 * 
 * Space Complexity: O(4n) for segment tree array
 * 
 * Author: Hacktoberfest 2025 Contributor
 * Date: October 2025
 */

#include <iostream>
#include <vector>
#include <algorithm>
#include <climits>
#include <chrono>
#include <random>

template<typename T>
class AdvancedSegmentTree {
private:
    std::vector<T> tree;
    std::vector<T> lazy;
    int n;
    
    // Build the segment tree
    void build(const std::vector<T>& arr, int node, int start, int end) {
        if (start == end) {
            tree[node] = arr[start];
        } else {
            int mid = (start + end) / 2;
            build(arr, 2 * node, start, mid);
            build(arr, 2 * node + 1, mid + 1, end);
            tree[node] = tree[2 * node] + tree[2 * node + 1];
        }
    }
    
    // Update lazy propagation
    void updateLazy(int node, int start, int end) {
        if (lazy[node] != 0) {
            tree[node] += lazy[node] * (end - start + 1);
            if (start != end) {
                lazy[2 * node] += lazy[node];
                lazy[2 * node + 1] += lazy[node];
            }
            lazy[node] = 0;
        }
    }
    
    // Range update with lazy propagation
    void updateRange(int node, int start, int end, int l, int r, T val) {
        updateLazy(node, start, end);
        if (start > r || end < l) {
            return;
        }
        
        if (start >= l && end <= r) {
            lazy[node] += val;
            updateLazy(node, start, end);
            return;
        }
        
        int mid = (start + end) / 2;
        updateRange(2 * node, start, mid, l, r, val);
        updateRange(2 * node + 1, mid + 1, end, l, r, val);
        
        updateLazy(2 * node, start, mid);
        updateLazy(2 * node + 1, mid + 1, end);
        tree[node] = tree[2 * node] + tree[2 * node + 1];
    }
    
    // Range sum query with lazy propagation
    T queryRange(int node, int start, int end, int l, int r) {
        if (start > r || end < l) {
            return 0;
        }
        
        updateLazy(node, start, end);
        
        if (start >= l && end <= r) {
            return tree[node];
        }
        
        int mid = (start + end) / 2;
        T p1 = queryRange(2 * node, start, mid, l, r);
        T p2 = queryRange(2 * node + 1, mid + 1, end, l, r);
        return p1 + p2;
    }
    
    // Point update
    void updatePoint(int node, int start, int end, int idx, T val) {
        updateLazy(node, start, end);
        
        if (start == end) {
            tree[node] = val;
        } else {
            int mid = (start + end) / 2;
            if (idx <= mid) {
                updatePoint(2 * node, start, mid, idx, val);
            } else {
                updatePoint(2 * node + 1, mid + 1, end, idx, val);
            }
            
            updateLazy(2 * node, start, mid);
            updateLazy(2 * node + 1, mid + 1, end);
            tree[node] = tree[2 * node] + tree[2 * node + 1];
        }
    }
    
    // Find minimum in range
    T queryMin(int node, int start, int end, int l, int r) {
        if (start > r || end < l) {
            return std::numeric_limits<T>::max();
        }
        
        updateLazy(node, start, end);
        
        if (start >= l && end <= r) {
            return tree[node];
        }
        
        int mid = (start + end) / 2;
        T p1 = queryMin(2 * node, start, mid, l, r);
        T p2 = queryMin(2 * node + 1, mid + 1, end, l, r);
        return std::min(p1, p2);
    }

public:
    // Constructor
    AdvancedSegmentTree(const std::vector<T>& arr) {
        n = arr.size();
        tree.resize(4 * n);
        lazy.resize(4 * n, 0);
        build(arr, 1, 0, n - 1);
    }
    
    // Public interface for range update
    void updateRange(int l, int r, T val) {
        updateRange(1, 0, n - 1, l, r, val);
    }
    
    // Public interface for range sum query
    T querySum(int l, int r) {
        return queryRange(1, 0, n - 1, l, r);
    }
    
    // Public interface for point update
    void updatePoint(int idx, T val) {
        updatePoint(1, 0, n - 1, idx, val);
    }
    
    // Public interface for range minimum query
    T queryMin(int l, int r) {
        return queryMin(1, 0, n - 1, l, r);
    }
    
    // Get current array state
    std::vector<T> getCurrentArray() {
        std::vector<T> result(n);
        for (int i = 0; i < n; i++) {
            result[i] = querySum(i, i);
        }
        return result;
    }
    
    // Print tree (for debugging)
    void printTree() {
        std::cout << "Segment Tree: ";
        for (int i = 1; i < 4 * n && tree[i] != 0; i++) {
            std::cout << tree[i] << " ";
        }
        std::cout << std::endl;
    }
};

// Binary Indexed Tree (Fenwick Tree) for comparison
template<typename T>
class BinaryIndexedTree {
private:
    std::vector<T> bit;
    int n;

public:
    BinaryIndexedTree(const std::vector<T>& arr) {
        n = arr.size();
        bit.assign(n + 1, 0);
        for (int i = 0; i < n; i++) {
            update(i, arr[i]);
        }
    }
    
    void update(int idx, T val) {
        for (++idx; idx <= n; idx += idx & -idx) {
            bit[idx] += val;
        }
    }
    
    T query(int idx) {
        T sum = 0;
        for (++idx; idx > 0; idx -= idx & -idx) {
            sum += bit[idx];
        }
        return sum;
    }
    
    T rangeQuery(int l, int r) {
        return query(r) - (l > 0 ? query(l - 1) : 0);
    }
};

// Test functions
void testSegmentTree() {
    std::cout << "=== Advanced Segment Tree Tests ===\n\n";
    
    // Test 1: Basic operations
    std::vector<int> arr = {1, 3, 5, 7, 9, 11, 13, 15};
    AdvancedSegmentTree<int> segTree(arr);
    
    std::cout << "Test 1: Basic Operations\n";
    std::cout << "Initial array: ";
    for (int x : arr) std::cout << x << " ";
    std::cout << std::endl;
    
    // Range sum queries
    std::cout << "Range sum [1, 5]: " << segTree.querySum(1, 5) << std::endl;
    std::cout << "Range sum [0, 7]: " << segTree.querySum(0, 7) << std::endl;
    
    // Range updates
    std::cout << "\nApplying range update [2, 6] += 10\n";
    segTree.updateRange(2, 6, 10);
    
    std::vector<int> updated = segTree.getCurrentArray();
    std::cout << "Updated array: ";
    for (int x : updated) std::cout << x << " ";
    std::cout << std::endl;
    
    std::cout << "Range sum [1, 5] after update: " << segTree.querySum(1, 5) << std::endl;
    
    std::cout << "\n" << std::string(50, '=') << "\n\n";
    
    // Test 2: Multiple range updates
    std::cout << "Test 2: Multiple Range Updates\n";
    segTree.updateRange(0, 3, 5);  // Add 5 to first 4 elements
    segTree.updateRange(4, 7, -3); // Subtract 3 from last 4 elements
    
    updated = segTree.getCurrentArray();
    std::cout << "After multiple updates: ";
    for (int x : updated) std::cout << x << " ";
    std::cout << std::endl;
    
    std::cout << "\n" << std::string(50, '=') << "\n\n";
    
    // Test 3: Point updates
    std::cout << "Test 3: Point Updates\n";
    segTree.updatePoint(0, 100);
    segTree.updatePoint(7, 200);
    
    updated = segTree.getCurrentArray();
    std::cout << "After point updates: ";
    for (int x : updated) std::cout << x << " ";
    std::cout << std::endl;
    
    std::cout << "Range sum [0, 7]: " << segTree.querySum(0, 7) << std::endl;
}

void performanceComparison() {
    std::cout << "\n=== Performance Comparison ===\n";
    
    const int n = 100000;
    const int operations = 10000;
    
    // Generate random array
    std::vector<int> arr(n);
    std::random_device rd;
    std::mt19937 gen(rd());
    std::uniform_int_distribution<> dis(1, 1000);
    
    for (int i = 0; i < n; i++) {
        arr[i] = dis(gen);
    }
    
    // Segment Tree performance
    auto start = std::chrono::high_resolution_clock::now();
    AdvancedSegmentTree<int> segTree(arr);
    auto build_time = std::chrono::high_resolution_clock::now();
    
    // Perform operations
    std::uniform_int_distribution<> index_dis(0, n-1);
    for (int i = 0; i < operations; i++) {
        int l = index_dis(gen);
        int r = index_dis(gen);
        if (l > r) std::swap(l, r);
        
        if (i % 2 == 0) {
            segTree.updateRange(l, r, dis(gen) % 10);
        } else {
            segTree.querySum(l, r);
        }
    }
    auto seg_end = std::chrono::high_resolution_clock::now();
    
    // BIT performance
    auto bit_start = std::chrono::high_resolution_clock::now();
    BinaryIndexedTree<int> bit(arr);
    auto bit_build = std::chrono::high_resolution_clock::now();
    
    for (int i = 0; i < operations / 2; i++) {  // BIT doesn't support range updates
        int l = index_dis(gen);
        int r = index_dis(gen);
        if (l > r) std::swap(l, r);
        bit.rangeQuery(l, r);
    }
    auto bit_end = std::chrono::high_resolution_clock::now();
    
    // Print results
    auto seg_build_ms = std::chrono::duration_cast<std::chrono::milliseconds>(build_time - start).count();
    auto seg_ops_ms = std::chrono::duration_cast<std::chrono::milliseconds>(seg_end - build_time).count();
    auto bit_build_ms = std::chrono::duration_cast<std::chrono::milliseconds>(bit_build - bit_start).count();
    auto bit_ops_ms = std::chrono::duration_cast<std::chrono::milliseconds>(bit_end - bit_build).count();
    
    std::cout << "Array size: " << n << " elements\n";
    std::cout << "Operations: " << operations << "\n\n";
    
    std::cout << "Segment Tree:\n";
    std::cout << "  Build time: " << seg_build_ms << " ms\n";
    std::cout << "  Operations time: " << seg_ops_ms << " ms\n";
    std::cout << "  Total time: " << (seg_build_ms + seg_ops_ms) << " ms\n\n";
    
    std::cout << "Binary Indexed Tree:\n";
    std::cout << "  Build time: " << bit_build_ms << " ms\n";
    std::cout << "  Operations time: " << bit_ops_ms << " ms\n";
    std::cout << "  Total time: " << (bit_build_ms + bit_ops_ms) << " ms\n";
    
    std::cout << "\nNote: Segment Tree supports range updates with lazy propagation,\n";
    std::cout << "while BIT only supports point updates and range queries.\n";
}

void advancedUseCases() {
    std::cout << "\n=== Advanced Use Cases ===\n\n";
    
    // Use case 1: Employee salary management
    std::cout << "Use Case 1: Employee Salary Management System\n";
    std::vector<int> salaries = {50000, 60000, 55000, 70000, 80000, 75000};
    AdvancedSegmentTree<int> salaryTree(salaries);
    
    std::cout << "Initial salaries: ";
    for (int s : salaries) std::cout << "$" << s << " ";
    std::cout << std::endl;
    
    // Department-wide bonus
    std::cout << "Giving $5000 bonus to department 1 (employees 1-3)\n";
    salaryTree.updateRange(1, 3, 5000);
    
    // Individual raise
    std::cout << "Giving employee 0 a $10000 raise\n";
    salaryTree.updatePoint(0, salaryTree.querySum(0, 0) + 10000);
    
    std::cout << "Total payroll: $" << salaryTree.querySum(0, 5) << std::endl;
    std::cout << "Department 1 payroll: $" << salaryTree.querySum(1, 3) << std::endl;
    
    auto updated_salaries = salaryTree.getCurrentArray();
    std::cout << "Updated salaries: ";
    for (int s : updated_salaries) std::cout << "$" << s << " ";
    std::cout << std::endl;
    
    std::cout << "\n" << std::string(40, '=') << "\n\n";
    
    // Use case 2: Stock price analysis
    std::cout << "Use Case 2: Stock Price Analysis\n";
    std::vector<int> prices = {100, 105, 102, 108, 95, 112, 118, 115};
    AdvancedSegmentTree<int> priceTree(prices);
    
    std::cout << "Stock prices: ";
    for (int p : prices) std::cout << "$" << p << " ";
    std::cout << std::endl;
    
    std::cout << "Total value of first 4 days: $" << priceTree.querySum(0, 3) << std::endl;
    std::cout << "Total value of last 4 days: $" << priceTree.querySum(4, 7) << std::endl;
    
    // Market correction (reduce all prices by 5%)
    std::cout << "Applying 5% market correction...\n";
    for (int i = 0; i < 8; i++) {
        int current = priceTree.querySum(i, i);
        int corrected = current * 0.95;
        priceTree.updatePoint(i, corrected);
    }
    
    auto corrected_prices = priceTree.getCurrentArray();
    std::cout << "Corrected prices: ";
    for (int p : corrected_prices) std::cout << "$" << p << " ";
    std::cout << std::endl;
}

int main() {
    std::cout << "Advanced Segment Tree Implementation\n";
    std::cout << "====================================\n\n";
    
    testSegmentTree();
    performanceComparison();
    advancedUseCases();
    
    std::cout << "\nImplementation complete! ✅\n";
    std::cout << "Features demonstrated:\n";
    std::cout << "✓ Range sum queries with lazy propagation\n";
    std::cout << "✓ Range updates with lazy propagation\n";
    std::cout << "✓ Point updates and queries\n";
    std::cout << "✓ Performance comparison with BIT\n";
    std::cout << "✓ Real-world use case examples\n";
    
    return 0;
}
