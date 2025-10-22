package compression

import (
	"container/heap"
	"fmt"
	"sort"
	"strings"
)

// HuffmanNode represents a node in the Huffman tree
type HuffmanNode struct {
	Char      rune
	Frequency int
	Left      *HuffmanNode
	Right     *HuffmanNode
}

// HuffmanHeap implements heap.Interface for HuffmanNode
type HuffmanHeap []*HuffmanNode

func (h HuffmanHeap) Len() int           { return len(h) }
func (h HuffmanHeap) Less(i, j int) bool { return h[i].Frequency < h[j].Frequency }
func (h HuffmanHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *HuffmanHeap) Push(x interface{}) {
	*h = append(*h, x.(*HuffmanNode))
}

func (h *HuffmanHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

// HuffmanEncoder encodes and decodes text using Huffman coding
type HuffmanEncoder struct {
	root        *HuffmanNode
	codes       map[rune]string
	reverseCodes map[string]rune
}

// NewHuffmanEncoder creates a new Huffman encoder
func NewHuffmanEncoder() *HuffmanEncoder {
	return &HuffmanEncoder{
		codes:        make(map[rune]string),
		reverseCodes: make(map[string]rune),
	}
}

// BuildTree builds the Huffman tree from the input text
func (he *HuffmanEncoder) BuildTree(text string) error {
	if len(text) == 0 {
		return fmt.Errorf("input text cannot be empty")
	}

	// Count character frequencies
	frequencies := make(map[rune]int)
	for _, char := range text {
		frequencies[char]++
	}

	// Handle single character case
	if len(frequencies) == 1 {
		for char := range frequencies {
			he.codes[char] = "0"
			he.reverseCodes["0"] = char
		}
		// Create a simple tree for single character
		for c := range frequencies {
			he.root = &HuffmanNode{Char: c, Frequency: frequencies[c]}
			break
		}
		return nil
	}

	// Create priority queue (min-heap)
	pq := &HuffmanHeap{}
	heap.Init(pq)

	// Add all characters to the heap
	for char, freq := range frequencies {
		heap.Push(pq, &HuffmanNode{
			Char:      char,
			Frequency: freq,
		})
	}

	// Build Huffman tree
	for pq.Len() > 1 {
		left := heap.Pop(pq).(*HuffmanNode)
		right := heap.Pop(pq).(*HuffmanNode)

		merged := &HuffmanNode{
			Char:      0, // Internal node has no character
			Frequency: left.Frequency + right.Frequency,
			Left:      left,
			Right:     right,
		}

		heap.Push(pq, merged)
	}

	he.root = heap.Pop(pq).(*HuffmanNode)

	// Generate codes
	he.generateCodes(he.root, "")

	return nil
}

// generateCodes generates Huffman codes for each character
func (he *HuffmanEncoder) generateCodes(node *HuffmanNode, code string) {
	if node == nil {
		return
	}

	// Leaf node - store the code
	if node.Left == nil && node.Right == nil {
		if code == "" {
			code = "0" // Single character case
		}
		he.codes[node.Char] = code
		he.reverseCodes[code] = node.Char
		return
	}

	// Internal node - recurse
	he.generateCodes(node.Left, code+"0")
	he.generateCodes(node.Right, code+"1")
}

// Encode compresses the input text using Huffman coding
func (he *HuffmanEncoder) Encode(text string) (string, error) {
	if he.root == nil {
		return "", fmt.Errorf("Huffman tree not built. Call BuildTree first")
	}

	var encoded strings.Builder
	for _, char := range text {
		code, exists := he.codes[char]
		if !exists {
			return "", fmt.Errorf("character '%c' not found in Huffman tree", char)
		}
		encoded.WriteString(code)
	}

	return encoded.String(), nil
}

// Decode decompresses the encoded string back to original text
func (he *HuffmanEncoder) Decode(encoded string) (string, error) {
	if he.root == nil {
		return "", fmt.Errorf("Huffman tree not built. Call BuildTree first")
	}

	if len(encoded) == 0 {
		return "", nil
	}

	var decoded strings.Builder
	current := he.root

	// Handle single character tree
	if current.Left == nil && current.Right == nil {
		for range encoded {
			decoded.WriteRune(current.Char)
		}
		return decoded.String(), nil
	}

	for _, bit := range encoded {
		if bit == '0' {
			current = current.Left
		} else if bit == '1' {
			current = current.Right
		} else {
			return "", fmt.Errorf("invalid bit '%c' in encoded string", bit)
		}

		if current == nil {
			return "", fmt.Errorf("invalid encoded string")
		}

		// Reached leaf node
		if current.Left == nil && current.Right == nil {
			decoded.WriteRune(current.Char)
			current = he.root
		}
	}

	// Check if we ended at root (complete decoding)
	if current != he.root {
		return "", fmt.Errorf("incomplete encoded string")
	}

	return decoded.String(), nil
}

// GetCodes returns the Huffman codes for all characters
func (he *HuffmanEncoder) GetCodes() map[rune]string {
	result := make(map[rune]string)
	for char, code := range he.codes {
		result[char] = code
	}
	return result
}

// GetCompressionRatio calculates the compression ratio
func (he *HuffmanEncoder) GetCompressionRatio(original, encoded string) float64 {
	if len(original) == 0 {
		return 0
	}
	originalBits := len(original) * 8 // Assuming 8 bits per character
	encodedBits := len(encoded)
	return 1.0 - float64(encodedBits)/float64(originalBits)
}

// PrintTree prints the Huffman tree structure
func (he *HuffmanEncoder) PrintTree() {
	if he.root == nil {
		fmt.Println("Tree is empty")
		return
	}
	fmt.Println("Huffman Tree Structure:")
	he.printNode(he.root, "", true)
}

// printNode recursively prints the tree structure
func (he *HuffmanEncoder) printNode(node *HuffmanNode, prefix string, isLast bool) {
	if node == nil {
		return
	}

	connector := "├── "
	if isLast {
		connector = "└── "
	}

	if node.Left == nil && node.Right == nil {
		// Leaf node
		fmt.Printf("%s%s'%c' (%d)\n", prefix, connector, node.Char, node.Frequency)
	} else {
		// Internal node
		fmt.Printf("%s%s[%d]\n", prefix, connector, node.Frequency)
	}

	childPrefix := prefix
	if isLast {
		childPrefix += "    "
	} else {
		childPrefix += "│   "
	}

	if node.Left != nil || node.Right != nil {
		he.printNode(node.Left, childPrefix, node.Right == nil)
		he.printNode(node.Right, childPrefix, true)
	}
}

// PrintCodes prints all character codes in a formatted way
func (he *HuffmanEncoder) PrintCodes() {
	if len(he.codes) == 0 {
		fmt.Println("No codes generated")
		return
	}

	fmt.Println("Huffman Codes:")
	
	// Sort characters for consistent output
	var chars []rune
	for char := range he.codes {
		chars = append(chars, char)
	}
	
	sort.Slice(chars, func(i, j int) bool {
		return chars[i] < chars[j]
	})

	for _, char := range chars {
		code := he.codes[char]
		fmt.Printf("'%c': %s\n", char, code)
	}
}

// ValidateTree checks if the Huffman tree is valid
func (he *HuffmanEncoder) ValidateTree() error {
	if he.root == nil {
		return fmt.Errorf("tree is empty")
	}

	// Check if all codes are unique
	codeSet := make(map[string]bool)
	for _, code := range he.codes {
		if codeSet[code] {
			return fmt.Errorf("duplicate code found: %s", code)
		}
		codeSet[code] = true
	}

	// Check if no code is a prefix of another (prefix property)
	codes := make([]string, 0, len(he.codes))
	for _, code := range he.codes {
		codes = append(codes, code)
	}

	for i := 0; i < len(codes); i++ {
		for j := i + 1; j < len(codes); j++ {
			if strings.HasPrefix(codes[i], codes[j]) || strings.HasPrefix(codes[j], codes[i]) {
				return fmt.Errorf("prefix property violated: '%s' and '%s'", codes[i], codes[j])
			}
		}
	}

	return nil
}

// CompressString is a convenience function that builds tree and encodes in one step
func CompressString(text string) (string, *HuffmanEncoder, error) {
	encoder := NewHuffmanEncoder()
	
	err := encoder.BuildTree(text)
	if err != nil {
		return "", nil, err
	}
	
	encoded, err := encoder.Encode(text)
	if err != nil {
		return "", nil, err
	}
	
	return encoded, encoder, nil
}

// DecompressString is a convenience function that decodes using existing encoder
func DecompressString(encoded string, encoder *HuffmanEncoder) (string, error) {
	return encoder.Decode(encoded)
}
