/**
 * Advanced Red-Black Tree Implementation in Java
 * Self-balancing binary search tree with guaranteed O(log n) operations
 * 
 * Time Complexities:
 * - Search: O(log n)
 * - Insert: O(log n) 
 * - Delete: O(log n)
 * - Traversal: O(n)
 * 
 * Space Complexity: O(n) for n nodes
 * 
 * Red-Black Tree Properties:
 * 1. Every node is either red or black
 * 2. Root is always black
 * 3. All leaves (NIL) are black
 * 4. Red nodes cannot have red children
 * 5. All paths from root to leaves contain same number of black nodes
 * 
 * @author Hacktoberfest 2025 Contributor
 * @date October 2025
 */

import java.util.*;
import java.util.function.Consumer;

public class AdvancedRedBlackTree<T extends Comparable<T>> {
    
    private enum Color {
        RED, BLACK
    }
    
    private class Node {
        T data;
        Color color;
        Node left, right, parent;
        
        Node(T data) {
            this.data = data;
            this.color = Color.RED; // New nodes are always red initially
            this.left = NIL;
            this.right = NIL;
            this.parent = null;
        }
        
        @Override
        public String toString() {
            return data + "(" + (color == Color.RED ? "R" : "B") + ")";
        }
    }
    
    private final Node NIL; // Sentinel node representing null
    private Node root;
    private int size;
    
    public AdvancedRedBlackTree() {
        NIL = new Node(null);
        NIL.color = Color.BLACK;
        NIL.left = NIL.right = NIL.parent = NIL;
        root = NIL;
        size = 0;
    }
    
    /**
     * Insert a value into the tree
     * @param data Value to insert
     * @return true if inserted, false if already exists
     */
    public boolean insert(T data) {
        if (data == null) return false;
        
        Node newNode = new Node(data);
        Node parent = NIL;
        Node current = root;
        
        // Find position for new node
        while (current != NIL) {
            parent = current;
            int cmp = data.compareTo(current.data);
            if (cmp < 0) {
                current = current.left;
            } else if (cmp > 0) {
                current = current.right;
            } else {
                return false; // Duplicate value
            }
        }
        
        newNode.parent = parent;
        
        if (parent == NIL) {
            root = newNode; // Tree was empty
        } else if (data.compareTo(parent.data) < 0) {
            parent.left = newNode;
        } else {
            parent.right = newNode;
        }
        
        size++;
        
        // Fix Red-Black tree properties
        insertFixup(newNode);
        
        return true;
    }
    
    /**
     * Fix Red-Black tree properties after insertion
     */
    private void insertFixup(Node node) {
        while (node.parent.color == Color.RED) {
            if (node.parent == node.parent.parent.left) {
                Node uncle = node.parent.parent.right;
                
                if (uncle.color == Color.RED) {
                    // Case 1: Uncle is red
                    node.parent.color = Color.BLACK;
                    uncle.color = Color.BLACK;
                    node.parent.parent.color = Color.RED;
                    node = node.parent.parent;
                } else {
                    if (node == node.parent.right) {
                        // Case 2: Uncle is black, node is right child
                        node = node.parent;
                        leftRotate(node);
                    }
                    // Case 3: Uncle is black, node is left child
                    node.parent.color = Color.BLACK;
                    node.parent.parent.color = Color.RED;
                    rightRotate(node.parent.parent);
                }
            } else {
                // Symmetric cases when parent is right child
                Node uncle = node.parent.parent.left;
                
                if (uncle.color == Color.RED) {
                    node.parent.color = Color.BLACK;
                    uncle.color = Color.BLACK;
                    node.parent.parent.color = Color.RED;
                    node = node.parent.parent;
                } else {
                    if (node == node.parent.left) {
                        node = node.parent;
                        rightRotate(node);
                    }
                    node.parent.color = Color.BLACK;
                    node.parent.parent.color = Color.RED;
                    leftRotate(node.parent.parent);
                }
            }
        }
        root.color = Color.BLACK; // Root is always black
    }
    
    /**
     * Delete a value from the tree
     * @param data Value to delete
     * @return true if deleted, false if not found
     */
    public boolean delete(T data) {
        if (data == null) return false;
        
        Node nodeToDelete = search(data);
        if (nodeToDelete == NIL) return false;
        
        Node replacement;
        Node nodeToFix;
        Color originalColor = nodeToDelete.color;
        
        if (nodeToDelete.left == NIL) {
            nodeToFix = nodeToDelete.right;
            transplant(nodeToDelete, nodeToDelete.right);
        } else if (nodeToDelete.right == NIL) {
            nodeToFix = nodeToDelete.left;
            transplant(nodeToDelete, nodeToDelete.left);
        } else {
            replacement = minimum(nodeToDelete.right);
            originalColor = replacement.color;
            nodeToFix = replacement.right;
            
            if (replacement.parent == nodeToDelete) {
                nodeToFix.parent = replacement;
            } else {
                transplant(replacement, replacement.right);
                replacement.right = nodeToDelete.right;
                replacement.right.parent = replacement;
            }
            
            transplant(nodeToDelete, replacement);
            replacement.left = nodeToDelete.left;
            replacement.left.parent = replacement;
            replacement.color = nodeToDelete.color;
        }
        
        if (originalColor == Color.BLACK) {
            deleteFixup(nodeToFix);
        }
        
        size--;
        return true;
    }
    
    /**
     * Fix Red-Black tree properties after deletion
     */
    private void deleteFixup(Node node) {
        while (node != root && node.color == Color.BLACK) {
            if (node == node.parent.left) {
                Node sibling = node.parent.right;
                
                if (sibling.color == Color.RED) {
                    sibling.color = Color.BLACK;
                    node.parent.color = Color.RED;
                    leftRotate(node.parent);
                    sibling = node.parent.right;
                }
                
                if (sibling.left.color == Color.BLACK && sibling.right.color == Color.BLACK) {
                    sibling.color = Color.RED;
                    node = node.parent;
                } else {
                    if (sibling.right.color == Color.BLACK) {
                        sibling.left.color = Color.BLACK;
                        sibling.color = Color.RED;
                        rightRotate(sibling);
                        sibling = node.parent.right;
                    }
                    
                    sibling.color = node.parent.color;
                    node.parent.color = Color.BLACK;
                    sibling.right.color = Color.BLACK;
                    leftRotate(node.parent);
                    node = root;
                }
            } else {
                // Symmetric cases
                Node sibling = node.parent.left;
                
                if (sibling.color == Color.RED) {
                    sibling.color = Color.BLACK;
                    node.parent.color = Color.RED;
                    rightRotate(node.parent);
                    sibling = node.parent.left;
                }
                
                if (sibling.right.color == Color.BLACK && sibling.left.color == Color.BLACK) {
                    sibling.color = Color.RED;
                    node = node.parent;
                } else {
                    if (sibling.left.color == Color.BLACK) {
                        sibling.right.color = Color.BLACK;
                        sibling.color = Color.RED;
                        leftRotate(sibling);
                        sibling = node.parent.left;
                    }
                    
                    sibling.color = node.parent.color;
                    node.parent.color = Color.BLACK;
                    sibling.left.color = Color.BLACK;
                    rightRotate(node.parent);
                    node = root;
                }
            }
        }
        node.color = Color.BLACK;
    }
    
    /**
     * Search for a value in the tree
     */
    public boolean contains(T data) {
        return search(data) != NIL;
    }
    
    private Node search(T data) {
        Node current = root;
        while (current != NIL) {
            int cmp = data.compareTo(current.data);
            if (cmp < 0) {
                current = current.left;
            } else if (cmp > 0) {
                current = current.right;
            } else {
                return current;
            }
        }
        return NIL;
    }
    
    /**
     * Tree rotations for balancing
     */
    private void leftRotate(Node x) {
        Node y = x.right;
        x.right = y.left;
        
        if (y.left != NIL) {
            y.left.parent = x;
        }
        
        y.parent = x.parent;
        
        if (x.parent == NIL) {
            root = y;
        } else if (x == x.parent.left) {
            x.parent.left = y;
        } else {
            x.parent.right = y;
        }
        
        y.left = x;
        x.parent = y;
    }
    
    private void rightRotate(Node y) {
        Node x = y.left;
        y.left = x.right;
        
        if (x.right != NIL) {
            x.right.parent = y;
        }
        
        x.parent = y.parent;
        
        if (y.parent == NIL) {
            root = x;
        } else if (y == y.parent.right) {
            y.parent.right = x;
        } else {
            y.parent.left = x;
        }
        
        x.right = y;
        y.parent = x;
    }
    
    /**
     * Helper methods
     */
    private void transplant(Node u, Node v) {
        if (u.parent == NIL) {
            root = v;
        } else if (u == u.parent.left) {
            u.parent.left = v;
        } else {
            u.parent.right = v;
        }
        v.parent = u.parent;
    }
    
    private Node minimum(Node node) {
        while (node.left != NIL) {
            node = node.left;
        }
        return node;
    }
    
    private Node maximum(Node node) {
        while (node.right != NIL) {
            node = node.right;
        }
        return node;
    }
    
    /**
     * Tree traversals
     */
    public void inorderTraversal(Consumer<T> action) {
        inorderHelper(root, action);
    }
    
    private void inorderHelper(Node node, Consumer<T> action) {
        if (node != NIL) {
            inorderHelper(node.left, action);
            action.accept(node.data);
            inorderHelper(node.right, action);
        }
    }
    
    public void levelOrderTraversal() {
        if (root == NIL) return;
        
        Queue<Node> queue = new LinkedList<>();
        queue.offer(root);
        
        while (!queue.isEmpty()) {
            int levelSize = queue.size();
            for (int i = 0; i < levelSize; i++) {
                Node current = queue.poll();
                System.out.print(current + " ");
                
                if (current.left != NIL) queue.offer(current.left);
                if (current.right != NIL) queue.offer(current.right);
            }
            System.out.println();
        }
    }
    
    /**
     * Tree properties and validation
     */
    public int size() {
        return size;
    }
    
    public boolean isEmpty() {
        return size == 0;
    }
    
    public int height() {
        return heightHelper(root);
    }
    
    private int heightHelper(Node node) {
        if (node == NIL) return 0;
        return 1 + Math.max(heightHelper(node.left), heightHelper(node.right));
    }
    
    public int blackHeight() {
        return blackHeightHelper(root);
    }
    
    private int blackHeightHelper(Node node) {
        if (node == NIL) return 1;
        
        int leftBlackHeight = blackHeightHelper(node.left);
        if (leftBlackHeight == 0) return 0; // Invalid tree
        
        return leftBlackHeight + (node.color == Color.BLACK ? 1 : 0);
    }
    
    public boolean isValidRedBlackTree() {
        if (root != NIL && root.color != Color.BLACK) return false;
        return validateHelper(root) != -1;
    }
    
    private int validateHelper(Node node) {
        if (node == NIL) return 1;
        
        // Check red node property
        if (node.color == Color.RED) {
            if (node.left.color == Color.RED || node.right.color == Color.RED) {
                return -1; // Red node has red child
            }
        }
        
        int leftBlackHeight = validateHelper(node.left);
        int rightBlackHeight = validateHelper(node.right);
        
        if (leftBlackHeight == -1 || rightBlackHeight == -1) return -1;
        if (leftBlackHeight != rightBlackHeight) return -1; // Black height mismatch
        
        return leftBlackHeight + (node.color == Color.BLACK ? 1 : 0);
    }
    
    /**
     * Get elements in sorted order
     */
    public List<T> toSortedList() {
        List<T> result = new ArrayList<>();
        inorderTraversal(result::add);
        return result;
    }
    
    /**
     * Range queries
     */
    public List<T> rangeQuery(T min, T max) {
        List<T> result = new ArrayList<>();
        rangeQueryHelper(root, min, max, result);
        return result;
    }
    
    private void rangeQueryHelper(Node node, T min, T max, List<T> result) {
        if (node == NIL) return;
        
        if (node.data.compareTo(min) > 0) {
            rangeQueryHelper(node.left, min, max, result);
        }
        
        if (node.data.compareTo(min) >= 0 && node.data.compareTo(max) <= 0) {
            result.add(node.data);
        }
        
        if (node.data.compareTo(max) < 0) {
            rangeQueryHelper(node.right, min, max, result);
        }
    }
    
    /**
     * Test and demonstration methods
     */
    public static void main(String[] args) {
        runComprehensiveTests();
    }
    
    public static void runComprehensiveTests() {
        System.out.println("=== Advanced Red-Black Tree Tests ===\\n");
        
        // Test 1: Basic operations
        System.out.println("Test 1: Basic Operations");
        AdvancedRedBlackTree<Integer> tree = new AdvancedRedBlackTree<>();
        
        int[] values = {10, 20, 30, 15, 25, 5, 1, 35, 40};
        System.out.println("Inserting values: " + Arrays.toString(values));
        
        for (int val : values) {
            tree.insert(val);
            System.out.println("Inserted " + val + ", tree valid: " + tree.isValidRedBlackTree());
        }
        
        System.out.println("Tree size: " + tree.size());
        System.out.println("Tree height: " + tree.height());
        System.out.println("Black height: " + tree.blackHeight());
        
        System.out.println("\\nInorder traversal:");
        tree.inorderTraversal(x -> System.out.print(x + " "));
        System.out.println();
        
        System.out.println("\\nLevel order traversal:");
        tree.levelOrderTraversal();
        
        System.out.println("\\n" + "=".repeat(50) + "\\n");
        
        // Test 2: Search operations
        System.out.println("Test 2: Search Operations");
        int[] searchValues = {15, 25, 100, 1, 50};
        
        for (int val : searchValues) {
            boolean found = tree.contains(val);
            System.out.println("Search " + val + ": " + (found ? "Found" : "Not found"));
        }
        
        System.out.println("\\n" + "=".repeat(50) + "\\n");
        
        // Test 3: Range queries
        System.out.println("Test 3: Range Queries");
        List<Integer> range15to30 = tree.rangeQuery(15, 30);
        System.out.println("Elements in range [15, 30]: " + range15to30);
        
        List<Integer> range5to25 = tree.rangeQuery(5, 25);
        System.out.println("Elements in range [5, 25]: " + range5to25);
        
        System.out.println("\\n" + "=".repeat(50) + "\\n");
        
        // Test 4: Deletion operations
        System.out.println("Test 4: Deletion Operations");
        int[] deleteValues = {1, 30, 15, 100};
        
        for (int val : deleteValues) {
            boolean deleted = tree.delete(val);
            System.out.println("Delete " + val + ": " + (deleted ? "Success" : "Not found"));
            if (deleted) {
                System.out.println("  Tree valid after deletion: " + tree.isValidRedBlackTree());
                System.out.println("  New size: " + tree.size());
            }
        }
        
        System.out.println("\\nFinal inorder traversal:");
        tree.inorderTraversal(x -> System.out.print(x + " "));
        System.out.println();
        
        System.out.println("\\n" + "=".repeat(50) + "\\n");
        
        // Test 5: Performance test
        performanceTest();
        
        // Test 6: String data type
        stringDataTest();
    }
    
    private static void performanceTest() {
        System.out.println("Test 5: Performance Test");
        
        AdvancedRedBlackTree<Integer> tree = new AdvancedRedBlackTree<>();
        TreeSet<Integer> javaTreeSet = new TreeSet<>();
        
        final int n = 100000;
        Random random = new Random(42); // Fixed seed for reproducibility
        
        // Generate random values
        Set<Integer> uniqueValues = new HashSet<>();
        while (uniqueValues.size() < n) {
            uniqueValues.add(random.nextInt(n * 10));
        }
        Integer[] values = uniqueValues.toArray(new Integer[0]);
        
        // Test insertion performance
        long startTime = System.currentTimeMillis();
        for (int val : values) {
            tree.insert(val);
        }
        long rbTreeInsertTime = System.currentTimeMillis() - startTime;
        
        startTime = System.currentTimeMillis();
        for (int val : values) {
            javaTreeSet.add(val);
        }
        long treeSetInsertTime = System.currentTimeMillis() - startTime;
        
        // Test search performance
        Integer[] searchValues = Arrays.copyOf(values, Math.min(10000, values.length));
        
        startTime = System.currentTimeMillis();
        for (int val : searchValues) {
            tree.contains(val);
        }
        long rbTreeSearchTime = System.currentTimeMillis() - startTime;
        
        startTime = System.currentTimeMillis();
        for (int val : searchValues) {
            javaTreeSet.contains(val);
        }
        long treeSetSearchTime = System.currentTimeMillis() - startTime;
        
        System.out.println("Performance comparison with " + n + " elements:");
        System.out.println("Red-Black Tree:");
        System.out.println("  Insert time: " + rbTreeInsertTime + " ms");
        System.out.println("  Search time: " + rbTreeSearchTime + " ms");
        System.out.println("  Tree height: " + tree.height());
        System.out.println("  Valid tree: " + tree.isValidRedBlackTree());
        
        System.out.println("Java TreeSet:");
        System.out.println("  Insert time: " + treeSetInsertTime + " ms");
        System.out.println("  Search time: " + treeSetSearchTime + " ms");
        
        System.out.println("\\n" + "=".repeat(50) + "\\n");
    }
    
    private static void stringDataTest() {
        System.out.println("Test 6: String Data Type");
        
        AdvancedRedBlackTree<String> stringTree = new AdvancedRedBlackTree<>();
        String[] words = {"apple", "banana", "cherry", "date", "elderberry", 
                         "fig", "grape", "honeydew", "kiwi", "lemon"};
        
        System.out.println("Inserting words: " + Arrays.toString(words));
        for (String word : words) {
            stringTree.insert(word);
        }
        
        System.out.println("\\nWords in alphabetical order:");
        stringTree.inorderTraversal(word -> System.out.print(word + " "));
        System.out.println();
        
        System.out.println("\\nWords starting with 'a' to 'f':");
        List<String> range = stringTree.rangeQuery("a", "f");
        System.out.println(range);
        
        System.out.println("\\nTree properties:");
        System.out.println("Size: " + stringTree.size());
        System.out.println("Height: " + stringTree.height());
        System.out.println("Valid: " + stringTree.isValidRedBlackTree());
    }
}
