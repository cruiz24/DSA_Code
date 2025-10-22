# Huffman Coding Implementation in Go

A complete implementation of Huffman coding algorithm for lossless data compression, featuring tree construction, encoding, and decoding capabilities.

## Features

- **Complete Huffman Algorithm**: Build optimal prefix codes
- **Lossless Compression**: Perfect reconstruction of original data
- **Tree Visualization**: Print Huffman tree structure
- **Compression Analysis**: Calculate compression ratios
- **Code Validation**: Verify prefix property and uniqueness
- **Edge Case Handling**: Single character, empty input handling
- **Comprehensive Testing**: Full test coverage with benchmarks

## Time Complexity

| Operation | Time Complexity |
|-----------|-----------------|
| Build Tree | O(n log k) where k = unique characters |
| Encode | O(n) |
| Decode | O(m) where m = encoded length |
| Validation | O(k²) |

## Space Complexity: O(k) where k = unique characters

## Usage

```go
package main

import (
    "fmt"
    "compression"
)

func main() {
    text := "hello world"
    
    // Method 1: Step by step
    encoder := compression.NewHuffmanEncoder()
    
    // Build Huffman tree
    err := encoder.BuildTree(text)
    if err != nil {
        panic(err)
    }
    
    // Encode text
    encoded, err := encoder.Encode(text)
    if err != nil {
        panic(err)
    }
    
    // Decode back to original
    decoded, err := encoder.Decode(encoded)
    if err != nil {
        panic(err)
    }
    
    fmt.Printf("Original: %s\n", text)
    fmt.Printf("Encoded:  %s\n", encoded)
    fmt.Printf("Decoded:  %s\n", decoded)
    
    // Calculate compression ratio
    ratio := encoder.GetCompressionRatio(text, encoded)
    fmt.Printf("Compression ratio: %.2f%%\n", ratio*100)
    
    // Method 2: One-step compression
    encoded2, encoder2, err := compression.CompressString(text)
    if err == nil {
        decoded2, _ := compression.DecompressString(encoded2, encoder2)
        fmt.Printf("One-step result: %s\n", decoded2)
    }
    
    // Display Huffman codes
    encoder.PrintCodes()
    
    // Display tree structure
    encoder.PrintTree()
}
```

## Algorithm Details

### Huffman Tree Construction
1. Count character frequencies
2. Create leaf nodes for each character
3. Build min-heap with nodes ordered by frequency
4. Repeatedly merge two nodes with lowest frequencies
5. Continue until only one node remains (root)

### Code Generation
- Traverse tree from root to leaves
- Assign '0' for left branches, '1' for right branches
- Path from root to leaf = character's Huffman code

### Prefix Property
- No code is prefix of another
- Enables unambiguous decoding
- Automatically satisfied by tree construction

## Applications

- **File Compression**: ZIP, GZIP implementations
- **Image Compression**: JPEG (modified Huffman)
- **Network Protocols**: HTTP/2 header compression
- **Data Storage**: Database compression
- **Streaming**: Real-time data compression

## Compression Examples

```go
// Text with repeated characters compresses well
text1 := "aaaaaabbbbccdd"
// Expected good compression ratio

// Uniformly distributed text compresses poorly  
text2 := "abcdefghijklmn"
// Expected poor compression ratio

// Real-world text usually compresses moderately well
text3 := "the quick brown fox jumps over the lazy dog"
```

## Advanced Features

### Tree Validation
```go
err := encoder.ValidateTree()
// Checks prefix property and code uniqueness
```

### Compression Analysis
```go
ratio := encoder.GetCompressionRatio(original, encoded)
// Returns compression ratio (0.0 to 1.0)
```

### Code Inspection
```go
codes := encoder.GetCodes()
// Returns map of character -> Huffman code
```

### Tree Structure Display
```go
encoder.PrintTree()
// Displays tree in hierarchical format
```

## Edge Cases Handled

- **Empty String**: Proper error handling
- **Single Character**: Special encoding with "0"
- **Two Characters**: Optimal 1-bit codes
- **Unicode Support**: Full rune support for international text

## Testing

Comprehensive test suite:

```bash
go test -v                    # All tests
go test -bench=.             # Benchmarks
go test -cover               # Coverage analysis
```

Test cases include:
- Basic encode/decode cycles
- Edge cases (empty, single char)
- Long text compression
- Invalid input handling
- Performance benchmarks

## Optimization Notes

- **Heap Operations**: Using Go's container/heap for efficiency
- **String Building**: Using strings.Builder for performance  
- **Memory Management**: Minimal allocations during encoding/decoding
- **Tree Traversal**: Efficient recursive algorithms

## Comparison with Other Compression

| Algorithm | Type | Ratio | Speed | Complexity |
|-----------|------|-------|-------|------------|
| Huffman | Statistical | Good | Fast | Medium |
| LZ77 | Dictionary | Better | Medium | High |
| Arithmetic | Statistical | Best | Slow | High |
| Run-Length | Simple | Poor | Very Fast | Low |

## Educational Value

This implementation demonstrates:
- Greedy algorithm design
- Binary trees and heaps
- Prefix codes and encoding theory
- Information theory concepts
- Practical compression techniques
- Algorithm optimization strategies

## Extensions

Possible enhancements:
- Adaptive Huffman coding
- Canonical Huffman codes
- Integration with other compression methods
- Parallel tree construction
- Memory-mapped file handling
