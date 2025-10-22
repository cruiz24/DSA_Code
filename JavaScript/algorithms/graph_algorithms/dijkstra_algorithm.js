/**
 * Advanced Dijkstra's Algorithm Implementation with Priority Queue
 * Finds shortest path from source to all other vertices in weighted graph
 * 
 * Time Complexity: O((V + E) log V) using binary heap
 * Space Complexity: O(V) for distance array and priority queue
 * 
 * Author: Hacktoberfest 2025 Contributor
 * Date: October 2025
 */

class MinHeap {
    constructor() {
        this.heap = [];
    }

    parent(i) { return Math.floor((i - 1) / 2); }
    leftChild(i) { return 2 * i + 1; }
    rightChild(i) { return 2 * i + 2; }

    insert(node) {
        this.heap.push(node);
        this.heapifyUp(this.heap.length - 1);
    }

    extractMin() {
        if (this.heap.length === 0) return null;
        if (this.heap.length === 1) return this.heap.pop();

        const min = this.heap[0];
        this.heap[0] = this.heap.pop();
        this.heapifyDown(0);
        return min;
    }

    heapifyUp(i) {
        while (i > 0 && this.heap[this.parent(i)].distance > this.heap[i].distance) {
            [this.heap[this.parent(i)], this.heap[i]] = [this.heap[i], this.heap[this.parent(i)]];
            i = this.parent(i);
        }
    }

    heapifyDown(i) {
        let minIndex = i;
        const left = this.leftChild(i);
        const right = this.rightChild(i);

        if (left < this.heap.length && this.heap[left].distance < this.heap[minIndex].distance) {
            minIndex = left;
        }

        if (right < this.heap.length && this.heap[right].distance < this.heap[minIndex].distance) {
            minIndex = right;
        }

        if (i !== minIndex) {
            [this.heap[i], this.heap[minIndex]] = [this.heap[minIndex], this.heap[i]];
            this.heapifyDown(minIndex);
        }
    }

    isEmpty() {
        return this.heap.length === 0;
    }
}

class DijkstraAlgorithm {
    /**
     * Finds shortest path from source to all vertices
     * @param {Array} graph - Adjacency list representation
     * @param {number} source - Source vertex
     * @returns {Object} Object containing distances and previous vertices
     */
    static findShortestPath(graph, source) {
        if (!graph || graph.length === 0) {
            throw new Error("Graph cannot be empty");
        }

        if (source < 0 || source >= graph.length) {
            throw new Error("Invalid source vertex");
        }

        const vertices = graph.length;
        const distances = Array(vertices).fill(Infinity);
        const previous = Array(vertices).fill(null);
        const visited = Array(vertices).fill(false);
        const pq = new MinHeap();

        // Initialize source
        distances[source] = 0;
        pq.insert({ vertex: source, distance: 0 });

        while (!pq.isEmpty()) {
            const current = pq.extractMin();
            const u = current.vertex;

            if (visited[u]) continue;
            visited[u] = true;

            // Process all neighbors
            for (const edge of graph[u] || []) {
                const v = edge.vertex;
                const weight = edge.weight;

                if (!visited[v] && distances[u] + weight < distances[v]) {
                    distances[v] = distances[u] + weight;
                    previous[v] = u;
                    pq.insert({ vertex: v, distance: distances[v] });
                }
            }
        }

        return { distances, previous };
    }

    /**
     * Reconstructs path from source to target
     * @param {Array} previous - Previous vertices array
     * @param {number} target - Target vertex
     * @returns {Array} Path from source to target
     */
    static getPath(previous, target) {
        const path = [];
        let current = target;

        while (current !== null) {
            path.unshift(current);
            current = previous[current];
        }

        return path.length > 1 ? path : [];
    }

    /**
     * Prints the shortest distances and paths from source
     * @param {Array} distances - Distances array
     * @param {Array} previous - Previous vertices array
     * @param {number} source - Source vertex
     */
    static printSolution(distances, previous, source) {
        console.log(`Shortest distances from vertex ${source}:`);
        console.log("Vertex\tDistance\tPath");
        console.log("------\t--------\t----");

        for (let i = 0; i < distances.length; i++) {
            const path = this.getPath(previous, i);
            const pathStr = path.length > 0 ? path.join(" → ") : "No path";
            const distStr = distances[i] === Infinity ? "∞" : distances[i].toString();
            
            console.log(`${i}\t${distStr}\t\t${pathStr}`);
        }
    }
}

// Test cases and examples
function runTests() {
    console.log("=== Dijkstra's Algorithm Tests ===\n");

    // Test Case 1: Standard weighted graph
    const graph1 = [
        [{ vertex: 1, weight: 4 }, { vertex: 2, weight: 1 }],              // 0
        [{ vertex: 3, weight: 1 }],                                        // 1
        [{ vertex: 1, weight: 2 }, { vertex: 3, weight: 5 }],              // 2
        []                                                                  // 3
    ];

    console.log("Test Case 1: Standard weighted graph");
    console.log("Graph structure:");
    console.log("0 → 1(4), 2(1)");
    console.log("1 → 3(1)");
    console.log("2 → 1(2), 3(5)");
    console.log("3 → (no edges)");
    
    const result1 = DijkstraAlgorithm.findShortestPath(graph1, 0);
    DijkstraAlgorithm.printSolution(result1.distances, result1.previous, 0);
    console.log("\n" + "=".repeat(50) + "\n");

    // Test Case 2: Larger graph with multiple paths
    const graph2 = [
        [{ vertex: 1, weight: 10 }, { vertex: 4, weight: 5 }],             // 0
        [{ vertex: 2, weight: 1 }, { vertex: 4, weight: 2 }],              // 1
        [{ vertex: 3, weight: 4 }],                                        // 2
        [{ vertex: 2, weight: 6 }, { vertex: 0, weight: 7 }],              // 3
        [{ vertex: 1, weight: 3 }, { vertex: 2, weight: 9 }, { vertex: 3, weight: 2 }] // 4
    ];

    console.log("Test Case 2: Complex graph with multiple paths");
    const result2 = DijkstraAlgorithm.findShortestPath(graph2, 0);
    DijkstraAlgorithm.printSolution(result2.distances, result2.previous, 0);
    console.log("\n" + "=".repeat(50) + "\n");

    // Test Case 3: Disconnected graph
    const graph3 = [
        [{ vertex: 1, weight: 1 }],     // 0
        [],                             // 1
        [{ vertex: 3, weight: 2 }],     // 2
        []                              // 3
    ];

    console.log("Test Case 3: Disconnected graph");
    const result3 = DijkstraAlgorithm.findShortestPath(graph3, 0);
    DijkstraAlgorithm.printSolution(result3.distances, result3.previous, 0);
}

// Performance comparison
function performanceTest() {
    console.log("\n=== Performance Test ===");
    
    // Generate random graph
    const vertices = 1000;
    const edges = 5000;
    const graph = Array(vertices).fill(null).map(() => []);

    // Add random edges
    for (let i = 0; i < edges; i++) {
        const from = Math.floor(Math.random() * vertices);
        const to = Math.floor(Math.random() * vertices);
        const weight = Math.floor(Math.random() * 100) + 1;
        
        if (from !== to) {
            graph[from].push({ vertex: to, weight });
        }
    }

    console.log(`Testing Dijkstra on graph with ${vertices} vertices and ~${edges} edges...`);
    const startTime = performance.now();
    DijkstraAlgorithm.findShortestPath(graph, 0);
    const endTime = performance.now();
    
    console.log(`Execution time: ${(endTime - startTime).toFixed(2)} milliseconds`);
}

// Export for module usage
if (typeof module !== 'undefined' && module.exports) {
    module.exports = { DijkstraAlgorithm, MinHeap };
}

// Run tests if this file is executed directly
if (typeof require !== 'undefined' && require.main === module) {
    runTests();
    performanceTest();
}
