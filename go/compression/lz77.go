// time complexity: O(n*w) where n is input length and w is window size
// space complexity: O(w) for sliding window buffer

package lz77

import (
	"fmt"
)

// LZ77Token represents a token in LZ77 compression
// Either a literal character or a back-reference (offset, length, next_char)
type LZ77Token struct {
	IsLiteral bool   // true if literal character, false if back-reference
	Literal   byte   // literal character (when IsLiteral = true)
	Offset    int    // offset for back-reference (when IsLiteral = false)
	Length    int    // length of match (when IsLiteral = false)
	NextChar  byte   // next character after match (when IsLiteral = false)
}

// LZ77Compressor implements the LZ77 compression algorithm
type LZ77Compressor struct {
	WindowSize int // size of sliding window (search buffer)
	BufferSize int // size of look-ahead buffer
}

// NewLZ77Compressor creates a new LZ77 compressor with specified parameters
func NewLZ77Compressor(windowSize, bufferSize int) *LZ77Compressor {
	if windowSize <= 0 {
		windowSize = 4096 // default window size
	}
	if bufferSize <= 0 {
		bufferSize = 18 // default buffer size
	}
	
	return &LZ77Compressor{
		WindowSize: windowSize,
		BufferSize: bufferSize,
	}
}

// Compress compresses the input data using LZ77 algorithm
// Returns a slice of LZ77 tokens representing the compressed data
func (lz *LZ77Compressor) Compress(input []byte) []LZ77Token {
	if len(input) == 0 {
		return []LZ77Token{}
	}

	var tokens []LZ77Token
	i := 0

	for i < len(input) {
		// Find the longest match in the sliding window
		bestOffset, bestLength := lz.findLongestMatch(input, i)

		if bestLength > 0 {
			// Found a match - create back-reference token
			nextChar := byte(0)
			if i+bestLength < len(input) {
				nextChar = input[i+bestLength]
			}

			token := LZ77Token{
				IsLiteral: false,
				Offset:    bestOffset,
				Length:    bestLength,
				NextChar:  nextChar,
			}
			tokens = append(tokens, token)
			
			// Move position forward by match length + 1 (for next char)
			if i+bestLength < len(input) {
				i += bestLength + 1
			} else {
				i += bestLength
			}
		} else {
			// No match found - create literal token
			token := LZ77Token{
				IsLiteral: true,
				Literal:   input[i],
			}
			tokens = append(tokens, token)
			i++
		}
	}

	return tokens
}

// Decompress decompresses LZ77 tokens back to original data
func (lz *LZ77Compressor) Decompress(tokens []LZ77Token) []byte {
	var result []byte

	for _, token := range tokens {
		if token.IsLiteral {
			// Literal character
			result = append(result, token.Literal)
		} else {
			// Back-reference
			start := len(result) - token.Offset
			if start < 0 {
				// Invalid offset - skip this token
				continue
			}

			// Copy the matched sequence
			for j := 0; j < token.Length; j++ {
				if start+j < len(result) {
					result = append(result, result[start+j])
				}
			}

			// Add the next character if it exists
			if token.NextChar != 0 {
				result = append(result, token.NextChar)
			}
		}
	}

	return result
}

// findLongestMatch finds the longest match in the sliding window
// Returns offset and length of the best match
func (lz *LZ77Compressor) findLongestMatch(input []byte, pos int) (bestOffset, bestLength int) {
	bestOffset = 0
	bestLength = 0

	// Define search window boundaries
	windowStart := 0
	if pos > lz.WindowSize {
		windowStart = pos - lz.WindowSize
	}

	// Define look-ahead buffer boundaries
	bufferEnd := pos + lz.BufferSize
	if bufferEnd > len(input) {
		bufferEnd = len(input)
	}

	// Search for matches in the sliding window
	for searchPos := windowStart; searchPos < pos; searchPos++ {
		matchLength := 0

		// Calculate maximum possible match length
		maxLength := pos - searchPos
		if maxLength > bufferEnd-pos {
			maxLength = bufferEnd - pos
		}

		// Find match length at current search position
		for matchLength < maxLength && 
			pos+matchLength < len(input) && 
			input[searchPos+matchLength] == input[pos+matchLength] {
			matchLength++
		}

		// Update best match if this one is longer
		if matchLength > bestLength && matchLength >= 3 { // minimum match length
			bestLength = matchLength
			bestOffset = pos - searchPos
		}
	}

	return bestOffset, bestLength
}

// CompressToBytes converts LZ77 tokens to a byte representation
// This is a simple encoding - real implementations would use more sophisticated encoding
func (lz *LZ77Compressor) CompressToBytes(tokens []LZ77Token) []byte {
	var result []byte

	for _, token := range tokens {
		if token.IsLiteral {
			// Encode literal: 0x00 followed by the character
			result = append(result, 0x00, token.Literal)
		} else {
			// Encode back-reference: 0x01 followed by offset (2 bytes), length (1 byte), next char
			result = append(result, 0x01)
			result = append(result, byte(token.Offset>>8), byte(token.Offset&0xFF)) // offset as 2 bytes
			result = append(result, byte(token.Length))                             // length as 1 byte
			result = append(result, token.NextChar)                                 // next character
		}
	}

	return result
}

// DecompressFromBytes converts byte representation back to LZ77 tokens
func (lz *LZ77Compressor) DecompressFromBytes(data []byte) []LZ77Token {
	var tokens []LZ77Token
	i := 0

	for i < len(data) {
		if i >= len(data) {
			break
		}

		if data[i] == 0x00 {
			// Literal token
			if i+1 < len(data) {
				token := LZ77Token{
					IsLiteral: true,
					Literal:   data[i+1],
				}
				tokens = append(tokens, token)
				i += 2
			} else {
				break
			}
		} else if data[i] == 0x01 {
			// Back-reference token
			if i+5 <= len(data) {
				offset := int(data[i+1])<<8 | int(data[i+2])
				length := int(data[i+3])
				nextChar := data[i+4]

				token := LZ77Token{
					IsLiteral: false,
					Offset:    offset,
					Length:    length,
					NextChar:  nextChar,
				}
				tokens = append(tokens, token)
				i += 5
			} else {
				break
			}
		} else {
			// Unknown token type - skip
			i++
		}
	}

	return tokens
}

// GetCompressionStats returns statistics about the compression
func (lz *LZ77Compressor) GetCompressionStats(originalSize int, tokens []LZ77Token) map[string]interface{} {
	literalCount := 0
	backRefCount := 0
	totalMatchLength := 0

	for _, token := range tokens {
		if token.IsLiteral {
			literalCount++
		} else {
			backRefCount++
			totalMatchLength += token.Length
		}
	}

	// Calculate compressed size (simplified estimation)
	compressedSize := literalCount*2 + backRefCount*5 // 2 bytes per literal, 5 bytes per back-ref

	compressionRatio := 0.0
	if originalSize > 0 {
		compressionRatio = float64(compressedSize) / float64(originalSize)
	}

	return map[string]interface{}{
		"original_size":       originalSize,
		"compressed_size":     compressedSize,
		"compression_ratio":   compressionRatio,
		"savings_percentage":  (1.0 - compressionRatio) * 100,
		"literal_tokens":      literalCount,
		"backreference_tokens": backRefCount,
		"total_tokens":        len(tokens),
		"average_match_length": func() float64 {
			if backRefCount > 0 {
				return float64(totalMatchLength) / float64(backRefCount)
			}
			return 0.0
		}(),
	}
}

// String returns a string representation of an LZ77 token for debugging
func (token LZ77Token) String() string {
	if token.IsLiteral {
		return fmt.Sprintf("L('%c')", token.Literal)
	}
	return fmt.Sprintf("R(offset:%d, length:%d, next:'%c')", 
		token.Offset, token.Length, token.NextChar)
}

// ValidateCompression validates that compression and decompression work correctly
func (lz *LZ77Compressor) ValidateCompression(original []byte) bool {
	tokens := lz.Compress(original)
	decompressed := lz.Decompress(tokens)
	
	if len(original) != len(decompressed) {
		return false
	}
	
	for i := 0; i < len(original); i++ {
		if original[i] != decompressed[i] {
			return false
		}
	}
	
	return true
}

// OptimizeParameters suggests optimal window and buffer sizes for given input
func OptimizeParameters(input []byte) (windowSize, bufferSize int) {
	inputSize := len(input)
	
	// Suggest window size based on input size
	switch {
	case inputSize < 1024:
		windowSize = 256
	case inputSize < 10240:
		windowSize = 1024
	case inputSize < 102400:
		windowSize = 4096
	default:
		windowSize = 8192
	}
	
	// Buffer size is typically much smaller than window size
	bufferSize = windowSize / 256
	if bufferSize < 16 {
		bufferSize = 16
	}
	if bufferSize > 32 {
		bufferSize = 32
	}
	
	return windowSize, bufferSize
}
