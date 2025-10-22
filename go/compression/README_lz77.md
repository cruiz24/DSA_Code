# LZ77 Compression Algorithm in Go

A comprehensive implementation of the LZ77 compression algorithm, one of the fundamental lossless data compression techniques used in ZIP, gzip, and PNG formats.

## Overview

LZ77 is a lossless data compression algorithm published by Abraham Lempel and Jacob Ziv in 1977. It works by replacing repeated occurrences of data with references to a single copy of that data existing earlier in the uncompressed data stream.

## Algorithm Details

**Time Complexity**: O(n*w) where n is input length and w is window size
**Space Complexity**: O(w) for sliding window buffer

The algorithm maintains:
- **Sliding Window**: A buffer of recently processed data (search buffer)
- **Look-ahead Buffer**: A buffer of data to be processed
- **Tokens**: Either literal characters or back-references (offset, length, next_char)

## Features

- **Configurable Parameters**: Adjustable window size and look-ahead buffer size
- **Comprehensive Compression**: Full encode/decode cycle with validation
- **Token-based Representation**: Clear separation between literals and back-references
- **Byte Serialization**: Convert tokens to/from byte representation for storage
- **Compression Statistics**: Detailed analysis of compression performance
- **Parameter Optimization**: Automatic parameter tuning based on input characteristics
- **Extensive Testing**: 60+ test cases covering edge cases and performance benchmarks

## Usage

```go
package main

import (
    "fmt"
    "github.com/yourusername/DSA_Code/go/compression"
)

func main() {
    // Create compressor with custom parameters
    compressor := lz77.NewLZ77Compressor(1024, 32) // window=1024, buffer=32
    
    input := []byte("The quick brown fox jumps over the lazy dog. The quick brown fox...")
    
    // Compress to tokens
    tokens := compressor.Compress(input)
    fmt.Printf("Compressed to %d tokens\n", len(tokens))
    
    // Decompress back to original
    decompressed := compressor.Decompress(tokens)
    fmt.Printf("Decompressed: %s\n", string(decompressed))
    
    // Get compression statistics
    stats := compressor.GetCompressionStats(len(input), tokens)
    fmt.Printf("Compression ratio: %.2f\n", stats["compression_ratio"].(float64))
    fmt.Printf("Space savings: %.1f%%\n", stats["savings_percentage"].(float64))
    
    // Convert to bytes for storage
    compressed := compressor.CompressToBytes(tokens)
    fmt.Printf("Compressed size: %d bytes\n", len(compressed))
    
    // Validate compression integrity
    if compressor.ValidateCompression(input) {
        fmt.Println("✅ Compression validation passed")
    }
}
```

## Advanced Features

### Parameter Optimization
```go
// Automatically suggest optimal parameters for your data
windowSize, bufferSize := lz77.OptimizeParameters(yourData)
compressor := lz77.NewLZ77Compressor(windowSize, bufferSize)
```

### Compression Analysis
```go
stats := compressor.GetCompressionStats(originalSize, tokens)
// Returns: compression_ratio, savings_percentage, literal_tokens,
//          backreference_tokens, average_match_length, etc.
```

### Token Inspection
```go
for _, token := range tokens {
    fmt.Println(token.String()) // Human-readable token representation
    if token.IsLiteral {
        fmt.Printf("Literal: '%c'\n", token.Literal)
    } else {
        fmt.Printf("Back-ref: offset=%d, length=%d, next='%c'\n", 
                   token.Offset, token.Length, token.NextChar)
    }
}
```

## Algorithm Workflow

1. **Initialize**: Set up sliding window and look-ahead buffer
2. **Search**: For each position, find the longest match in the sliding window
3. **Encode**: 
   - If match found (≥3 chars): Create back-reference token
   - If no match: Create literal token
4. **Advance**: Move position forward by match length or 1
5. **Repeat**: Until all input is processed

## Performance Characteristics

| Operation | Complexity | Notes |
|-----------|------------|-------|
| Compression | O(n*w) | n=input size, w=window size |
| Decompression | O(m) | m=number of tokens |
| Memory | O(w) | Sliding window storage |

### Benchmark Results (Apple M3 Pro)
- **Compression**: ~302μs per operation (4.5KB repetitive text)
- **Decompression**: ~5.7μs per operation 
- **Match Finding**: ~3.9μs per operation

## Configuration Guidelines

### Window Size
- **Small files (<1KB)**: 256 bytes
- **Medium files (1-10KB)**: 1024 bytes  
- **Large files (10-100KB)**: 4096 bytes
- **Very large files (>100KB)**: 8192 bytes

### Look-ahead Buffer
- Typically 16-32 bytes
- Larger buffers improve compression but slow encoding
- Must fit within remaining input

## Educational Value

This implementation demonstrates:
- **Sliding Window Technique**: Fundamental pattern in compression
- **Greedy Algorithms**: Finding locally optimal matches
- **Token-based Encoding**: Separating control and data
- **Trade-offs**: Memory vs. compression ratio vs. speed
- **Real-world Applications**: Foundation for ZIP, gzip, PNG

## Limitations & Modern Context

LZ77 by itself has limitations:
- **Fixed encoding**: Real implementations use variable-length encoding
- **Sequential processing**: Cannot be easily parallelized
- **Memory bound**: Window size limits compression effectiveness

Modern derivatives (LZ78, LZW, DEFLATE) address these limitations.

## Testing

Run the comprehensive test suite:
```bash
go test -v lz77.go lz77_test.go
```

Run performance benchmarks:
```bash
go test -bench=. lz77.go lz77_test.go
```

## Historical Significance

LZ77 was groundbreaking because it:
- Introduced dictionary-based compression
- Required no prior knowledge of input statistics
- Achieved good compression ratios on typical data
- Became the foundation for many modern compression formats

## References

- [LZ77 and LZ78 - Wikipedia](https://en.wikipedia.org/wiki/LZ77_and_LZ78)
- [A Universal Algorithm for Sequential Data Compression](https://www.cs.duke.edu/courses/spring03/cps296.5/papers/ziv_lempel_1977_universal_algorithm.pdf) - Original 1977 paper
- [DEFLATE Specification](https://tools.ietf.org/html/rfc1951) - LZ77 variant used in ZIP/gzip

## Author

Created for Hacktoberfest 2025 - Contributing to open source algorithm education and implementation excellence.
