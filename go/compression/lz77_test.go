package lz77

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestNewLZ77Compressor(t *testing.T) {
	tests := []struct {
		name           string
		windowSize     int
		bufferSize     int
		expectedWindow int
		expectedBuffer int
	}{
		{
			"Valid parameters",
			1024,
			16,
			1024,
			16,
		},
		{
			"Zero window size uses default",
			0,
			16,
			4096,
			16,
		},
		{
			"Zero buffer size uses default",
			1024,
			0,
			1024,
			18,
		},
		{
			"Negative parameters use defaults",
			-100,
			-50,
			4096,
			18,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			compressor := NewLZ77Compressor(tt.windowSize, tt.bufferSize)
			if compressor.WindowSize != tt.expectedWindow {
				t.Errorf("Expected window size %d, got %d", tt.expectedWindow, compressor.WindowSize)
			}
			if compressor.BufferSize != tt.expectedBuffer {
				t.Errorf("Expected buffer size %d, got %d", tt.expectedBuffer, compressor.BufferSize)
			}
		})
	}
}

func TestCompress(t *testing.T) {
	compressor := NewLZ77Compressor(256, 16)

	tests := []struct {
		name     string
		input    string
		validate bool // whether to validate the output can be decompressed correctly
	}{
		{
			"Empty string",
			"",
			true,
		},
		{
			"Single character",
			"A",
			true,
		},
		{
			"No repetition",
			"ABCDEFG",
			true,
		},
		{
			"Simple repetition",
			"AAAA",
			true,
		},
		{
			"Complex repetition",
			"ABCABCABC",
			true,
		},
		{
			"Long repeated pattern",
			"ABCDEFGABCDEFGABCDEFG",
			true,
		},
		{
			"Text with spaces",
			"THE QUICK BROWN FOX JUMPS OVER THE LAZY DOG",
			true,
		},
		{
			"Highly repetitive text",
			strings.Repeat("PATTERN", 10),
			true,
		},
		{
			"Mixed repetition and unique",
			"ABCXYZABCXYZ123456",
			true,
		},
		{
			"Long string with overlapping matches",
			"AABAABAABAABAABAABAAB",
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := []byte(tt.input)
			tokens := compressor.Compress(input)
			
			if len(input) == 0 {
				if len(tokens) != 0 {
					t.Errorf("Expected empty tokens for empty input, got %d tokens", len(tokens))
				}
				return
			}

			// Validate that we can decompress back to original
			if tt.validate {
				decompressed := compressor.Decompress(tokens)
				if !bytes.Equal(input, decompressed) {
					t.Errorf("Decompressed data doesn't match original.\nOriginal: %q\nDecompressed: %q", 
						string(input), string(decompressed))
				}
			}
		})
	}
}

func TestDecompress(t *testing.T) {
	compressor := NewLZ77Compressor(256, 16)

	tests := []struct {
		name     string
		tokens   []LZ77Token
		expected string
	}{
		{
			"Empty tokens",
			[]LZ77Token{},
			"",
		},
		{
			"Single literal",
			[]LZ77Token{{IsLiteral: true, Literal: 'A'}},
			"A",
		},
		{
			"Multiple literals",
			[]LZ77Token{
				{IsLiteral: true, Literal: 'A'},
				{IsLiteral: true, Literal: 'B'},
				{IsLiteral: true, Literal: 'C'},
			},
			"ABC",
		},
		{
			"Literal then back-reference",
			[]LZ77Token{
				{IsLiteral: true, Literal: 'A'},
				{IsLiteral: true, Literal: 'B'},
				{IsLiteral: false, Offset: 2, Length: 1, NextChar: 'C'},
			},
			"ABAC",
		},
		{
			"Complex sequence",
			[]LZ77Token{
				{IsLiteral: true, Literal: 'A'},
				{IsLiteral: true, Literal: 'B'},
				{IsLiteral: true, Literal: 'C'},
				{IsLiteral: false, Offset: 3, Length: 2, NextChar: 'D'},
			},
			"ABCABD",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := compressor.Decompress(tt.tokens)
			if string(result) != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, string(result))
			}
		})
	}
}

func TestCompressDecompressRoundTrip(t *testing.T) {
	compressor := NewLZ77Compressor(1024, 32)

	testCases := []string{
		"",
		"A",
		"HELLO WORLD",
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		"AAAAAAAAAAAAAAAAAAAAAAAA",
		"ABCABCABCABCABCABC",
		"THE QUICK BROWN FOX JUMPS OVER THE LAZY DOG",
		"Lorem ipsum dolor sit amet, consectetur adipiscing elit. " +
			"Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
		strings.Repeat("COMPRESSION", 20),
		"MixedCaseAndNumbers123456789",
		"!@#$%^&*()_+{}[]|\\:;\"'<>?,./",
	}

	for _, testCase := range testCases {
		t.Run("Round trip: "+testCase[:min(20, len(testCase))], func(t *testing.T) {
			original := []byte(testCase)
			tokens := compressor.Compress(original)
			decompressed := compressor.Decompress(tokens)

			if !bytes.Equal(original, decompressed) {
				t.Errorf("Round trip failed.\nOriginal: %q\nDecompressed: %q", 
					string(original), string(decompressed))
			}
		})
	}
}

func TestFindLongestMatch(t *testing.T) {
	compressor := NewLZ77Compressor(256, 16)

	tests := []struct {
		name         string
		input        string
		pos          int
		expectedOff  int
		expectedLen  int
	}{
		{
			"No match at beginning",
			"ABCDEF",
			0,
			0,
			0,
		},
		{
			"Simple match",
			"ABCABC",
			3,
			3,
			3,
		},
		{
			"Partial match below minimum length",
			"ABCABX",
			3,
			0, // No match because minimum length is 3
			0,
		},
		{
			"No match when position is 0",
			"AAAAAA",
			0,
			0,
			0,
		},
		{
			"Long match",
			"PATTERNPATTERN",
			7,
			7,
			7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := []byte(tt.input)
			offset, length := compressor.findLongestMatch(input, tt.pos)
			if offset != tt.expectedOff || length != tt.expectedLen {
				t.Errorf("Expected offset=%d, length=%d; got offset=%d, length=%d",
					tt.expectedOff, tt.expectedLen, offset, length)
			}
		})
	}
}

func TestCompressToBytes(t *testing.T) {
	compressor := NewLZ77Compressor(256, 16)

	tokens := []LZ77Token{
		{IsLiteral: true, Literal: 'A'},
		{IsLiteral: false, Offset: 1, Length: 2, NextChar: 'B'},
	}

	result := compressor.CompressToBytes(tokens)
	
	// Check the structure: 
	// Literal 'A': [0x00, 'A']
	// Back-ref: [0x01, 0x00, 0x01, 0x02, 'B'] (offset=1, length=2, next='B')
	expected := []byte{0x00, 'A', 0x01, 0x00, 0x01, 0x02, 'B'}
	
	if !bytes.Equal(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestDecompressFromBytes(t *testing.T) {
	compressor := NewLZ77Compressor(256, 16)

	// Input bytes representing: Literal 'A', Back-ref (offset=1, length=2, next='B')
	input := []byte{0x00, 'A', 0x01, 0x00, 0x01, 0x02, 'B'}
	
	tokens := compressor.DecompressFromBytes(input)
	
	if len(tokens) != 2 {
		t.Errorf("Expected 2 tokens, got %d", len(tokens))
		return
	}
	
	// Check first token (literal)
	if !tokens[0].IsLiteral || tokens[0].Literal != 'A' {
		t.Errorf("First token should be literal 'A', got %+v", tokens[0])
	}
	
	// Check second token (back-reference)
	if tokens[1].IsLiteral || tokens[1].Offset != 1 || tokens[1].Length != 2 || tokens[1].NextChar != 'B' {
		t.Errorf("Second token should be back-ref (1,2,'B'), got %+v", tokens[1])
	}
}

func TestGetCompressionStats(t *testing.T) {
	compressor := NewLZ77Compressor(256, 16)
	
	tokens := []LZ77Token{
		{IsLiteral: true, Literal: 'A'},
		{IsLiteral: true, Literal: 'B'},
		{IsLiteral: false, Offset: 2, Length: 2, NextChar: 'C'},
	}
	
	stats := compressor.GetCompressionStats(5, tokens)
	
	expectedStats := map[string]interface{}{
		"original_size":       5,
		"compressed_size":     9, // 2 literals * 2 bytes + 1 back-ref * 5 bytes
		"literal_tokens":      2,
		"backreference_tokens": 1,
		"total_tokens":        3,
	}
	
	for key, expected := range expectedStats {
		if stats[key] != expected {
			t.Errorf("Expected %s=%v, got %v", key, expected, stats[key])
		}
	}
	
	// Check calculated values
	if ratio := stats["compression_ratio"].(float64); ratio != 1.8 {
		t.Errorf("Expected compression ratio 1.8, got %f", ratio)
	}
}

func TestValidateCompression(t *testing.T) {
	compressor := NewLZ77Compressor(512, 32)
	
	testCases := [][]byte{
		[]byte(""),
		[]byte("A"),
		[]byte("HELLO WORLD"),
		[]byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ"),
		[]byte(strings.Repeat("PATTERN", 5)),
		[]byte("The quick brown fox jumps over the lazy dog"),
	}
	
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("Validation test %d", i), func(t *testing.T) {
			if !compressor.ValidateCompression(testCase) {
				t.Errorf("Validation failed for input: %q", string(testCase))
			}
		})
	}
}

func TestOptimizeParameters(t *testing.T) {
	tests := []struct {
		name        string
		inputSize   int
		expectWin   int
		expectBuf   int
	}{
		{
			"Small input",
			500,
			256,
			16,
		},
		{
			"Medium input",
			5000,
			1024,
			16,
		},
		{
			"Large input",
			50000,
			4096,
			16,
		},
		{
			"Very large input",
			500000,
			8192,
			32,
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := make([]byte, tt.inputSize)
			windowSize, bufferSize := OptimizeParameters(input)
			
			if windowSize != tt.expectWin {
				t.Errorf("Expected window size %d, got %d", tt.expectWin, windowSize)
			}
			if bufferSize != tt.expectBuf {
				t.Errorf("Expected buffer size %d, got %d", tt.expectBuf, bufferSize)
			}
		})
	}
}

func TestLZ77TokenString(t *testing.T) {
	tests := []struct {
		name     string
		token    LZ77Token
		expected string
	}{
		{
			"Literal token",
			LZ77Token{IsLiteral: true, Literal: 'A'},
			"L('A')",
		},
		{
			"Back-reference token",
			LZ77Token{IsLiteral: false, Offset: 5, Length: 3, NextChar: 'X'},
			"R(offset:5, length:3, next:'X')",
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.token.String()
			if result != tt.expected {
				t.Errorf("Expected %q, got %q", tt.expected, result)
			}
		})
	}
}

// Helper function for min calculation
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Benchmark tests
func BenchmarkCompress(b *testing.B) {
	compressor := NewLZ77Compressor(1024, 32)
	input := []byte(strings.Repeat("The quick brown fox jumps over the lazy dog. ", 100))
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		compressor.Compress(input)
	}
}

func BenchmarkDecompress(b *testing.B) {
	compressor := NewLZ77Compressor(1024, 32)
	input := []byte(strings.Repeat("The quick brown fox jumps over the lazy dog. ", 100))
	tokens := compressor.Compress(input)
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		compressor.Decompress(tokens)
	}
}

func BenchmarkFindLongestMatch(b *testing.B) {
	compressor := NewLZ77Compressor(1024, 32)
	input := []byte(strings.Repeat("PATTERN", 200))
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		compressor.findLongestMatch(input, len(input)/2)
	}
}
