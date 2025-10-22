/**
 * Advanced Graph Algorithms in JavaScript
 * Implementation of Floyd-Warshall Algorithm for All Pairs Shortest Path
 * 
 * Time Complexity: O(V³) where V is the number of vertices
 * Space Complexity: O(V²) for the distance matrix
 * 
 * Author: Hacktoberfest 2025 Contributor
 * Date: October 2025
 */

class FloydWarshall {
    /**
     * Finds shortest paths between all pairs of vertices
     * @param {number[][]} graph - Adjacency matrix representation of graph
     * @returns {number[][]} Distance matrix with shortest paths
     */
    static findAllPairsShortestPath(graph) {
        if (!graph || graph.length === 0) {
            throw new Error("Graph cannot be empty");
        }

        const vertices = graph.length;
        
        // Initialize distance matrix
        const dist = Array(vertices).fill(null).map(() => Array(vertices).fill(Infinity));
        
        // Copy the input graph to distance matrix
        for (let i = 0; i < vertices; i++) {
            for (let j = 0; j < vertices; j++) {
                if (i === j) {
                    dist[i][j] = 0;
                } else if (graph[i][j] !== 0) {
                    dist[i][j] = graph[i][j];
                }
            }
        }

        // Floyd-Warshall algorithm core
        for (let k = 0; k < vertices; k++) {
            for (let i = 0; i < vertices; i++) {
                for (let j = 0; j < vertices; j++) {
                    // If vertex k is on the shortest path from i to j
                    if (dist[i][k] !== Infinity && dist[k][j] !== Infinity) {
                        dist[i][j] = Math.min(dist[i][j], dist[i][k] + dist[k][j]);
                    }
                }
            }
        }

        return dist;
    }

    /**
     * Detects negative weight cycles in the graph
     * @param {number[][]} dist - Distance matrix from Floyd-Warshall
     * @returns {boolean} True if negative cycle exists
     */
    static hasNegativeCycle(dist) {
        for (let i = 0; i < dist.length; i++) {
            if (dist[i][i] < 0) {
                return true;
            }
        }
        return false;
    }

    /**
     * Prints the shortest path matrix in a formatted way
     * @param {number[][]} dist - Distance matrix
     */
    static printSolution(dist) {
        console.log("Shortest distances between every pair of vertices:");
        console.log("     ", Array.from({length: dist.length}, (_, i) => i).join("    "));
        
        for (let i = 0; i < dist.length; i++) {
            let row = `${i}:   `;
            for (let j = 0; j < dist.length; j++) {
                if (dist[i][j] === Infinity) {
                    row += "INF  ";
                } else {
                    row += `${dist[i][j].toString().padStart(3)}  `;
                }
            }
            console.log(row);
        }
    }
}

// Test cases and examples
function runTests() {
    console.log("=== Floyd-Warshall Algorithm Tests ===\n");

    // Test Case 1: Simple 4-vertex graph
    const graph1 = [
        [0, 5, Infinity, 10],
        [Infinity, 0, 3, Infinity],
        [Infinity, Infinity, 0, 1],
        [Infinity, Infinity, Infinity, 0]
    ];

    console.log("Test Case 1: 4-vertex graph");
    console.log("Input graph:");
    FloydWarshall.printSolution(graph1);
    
    const result1 = FloydWarshall.findAllPairsShortestPath(graph1);
    console.log("\nShortest path distances:");
    FloydWarshall.printSolution(result1);
    console.log("Has negative cycle:", FloydWarshall.hasNegativeCycle(result1));
    console.log("\n" + "=".repeat(50) + "\n");

    // Test Case 2: Graph with negative weights (but no negative cycle)
    const graph2 = [
        [0, -1, 4],
        [Infinity, 0, 3],
        [Infinity, Infinity, 0]
    ];

    console.log("Test Case 2: Graph with negative weights");
    console.log("Input graph:");
    FloydWarshall.printSolution(graph2);
    
    const result2 = FloydWarshall.findAllPairsShortestPath(graph2);
    console.log("\nShortest path distances:");
    FloydWarshall.printSolution(result2);
    console.log("Has negative cycle:", FloydWarshall.hasNegativeCycle(result2));
    console.log("\n" + "=".repeat(50) + "\n");

    // Test Case 3: Disconnected graph
    const graph3 = [
        [0, 1, Infinity],
        [1, 0, Infinity],
        [Infinity, Infinity, 0]
    ];

    console.log("Test Case 3: Disconnected graph");
    console.log("Input graph:");
    FloydWarshall.printSolution(graph3);
    
    const result3 = FloydWarshall.findAllPairsShortestPath(graph3);
    console.log("\nShortest path distances:");
    FloydWarshall.printSolution(result3);
    console.log("Has negative cycle:", FloydWarshall.hasNegativeCycle(result3));
}

// Performance testing
function performanceTest() {
    console.log("\n=== Performance Test ===");
    
    const size = 100;
    const largeGraph = Array(size).fill(null).map((_, i) => 
        Array(size).fill(null).map((_, j) => {
            if (i === j) return 0;
            if (Math.random() < 0.3) return Math.floor(Math.random() * 20) + 1;
            return Infinity;
        })
    );

    console.log(`Testing Floyd-Warshall on ${size}x${size} graph...`);
    const startTime = performance.now();
    FloydWarshall.findAllPairsShortestPath(largeGraph);
    const endTime = performance.now();
    
    console.log(`Execution time: ${(endTime - startTime).toFixed(2)} milliseconds`);
}

// Export for module usage
if (typeof module !== 'undefined' && module.exports) {
    module.exports = FloydWarshall;
}

// Run tests if this file is executed directly
if (typeof require !== 'undefined' && require.main === module) {
    runTests();
    performanceTest();
}
