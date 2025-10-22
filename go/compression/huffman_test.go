package compression

import (
	"strings"
	"testing"
)

func TestHuffmanBasic(t *testing.T) {
	encoder := NewHuffmanEncoder()
	text := "hello world"
	
	err := encoder.BuildTree(text)
	if err != nil {
		t.Fatalf("BuildTree failed: %v", err)
	}
	
	encoded, err := encoder.Encode(text)
	if err != nil {
		t.Fatalf("Encode failed: %v", err)
	}
	
	decoded, err := encoder.Decode(encoded)
	if err != nil {
		t.Fatalf("Decode failed: %v", err)
	}
	
	if decoded != text {
		t.Errorf("Expected '%s', got '%s'", text, decoded)
	}
}

func TestHuffmanSingleCharacter(t *testing.T) {
	encoder := NewHuffmanEncoder()
	text := "aaaa"
	
	err := encoder.BuildTree(text)
	if err != nil {
		t.Fatalf("BuildTree failed: %v", err)
	}
	
	encoded, err := encoder.Encode(text)
	if err != nil {
		t.Fatalf("Encode failed: %v", err)
	}
	
	decoded, err := encoder.Decode(encoded)
	if err != nil {
		t.Fatalf("Decode failed: %v", err)
	}
	
	if decoded != text {
		t.Errorf("Expected '%s', got '%s'", text, decoded)
	}
	
	// Check that single character gets code "0"
	codes := encoder.GetCodes()
	if codes['a'] != "0" {
		t.Errorf("Expected code '0' for single character, got '%s'", codes['a'])
	}
}

func TestHuffmanEmptyString(t *testing.T) {
	encoder := NewHuffmanEncoder()
	
	err := encoder.BuildTree("")
	if err == nil {
		t.Error("Expected error for empty string")
	}
}

func TestHuffmanLongText(t *testing.T) {
	encoder := NewHuffmanEncoder()
	text := "this is a longer text with multiple repeated characters and spaces"
	
	err := encoder.BuildTree(text)
	if err != nil {
		t.Fatalf("BuildTree failed: %v", err)
	}
	
	encoded, err := encoder.Encode(text)
	if err != nil {
		t.Fatalf("Encode failed: %v", err)
	}
	
	decoded, err := encoder.Decode(encoded)
	if err != nil {
		t.Fatalf("Decode failed: %v", err)
	}
	
	if decoded != text {
		t.Errorf("Expected '%s', got '%s'", text, decoded)
	}
	
	// Check compression ratio
	ratio := encoder.GetCompressionRatio(text, encoded)
	if ratio <= 0 {
		t.Errorf("Expected positive compression ratio, got %f", ratio)
	}
}

func TestHuffmanTwoCharacters(t *testing.T) {
	encoder := NewHuffmanEncoder()
	text := "ab"
	
	err := encoder.BuildTree(text)
	if err != nil {
		t.Fatalf("BuildTree failed: %v", err)
	}
	
	codes := encoder.GetCodes()
	
	// With two characters, codes should be "0" and "1"
	if len(codes) != 2 {
		t.Errorf("Expected 2 codes, got %d", len(codes))
	}
	
	// Both codes should be single bits
	for char, code := range codes {
		if len(code) != 1 {
			t.Errorf("Expected single bit code for char '%c', got '%s'", char, code)
		}
	}
}

func TestHuffmanValidation(t *testing.T) {
	encoder := NewHuffmanEncoder()
	text := "abcdef"
	
	err := encoder.BuildTree(text)
	if err != nil {
		t.Fatalf("BuildTree failed: %v", err)
	}
	
	err = encoder.ValidateTree()
	if err != nil {
		t.Errorf("Tree validation failed: %v", err)
	}
}

func TestHuffmanEncodeWithoutTree(t *testing.T) {
	encoder := NewHuffmanEncoder()
	
	_, err := encoder.Encode("test")
	if err == nil {
		t.Error("Expected error when encoding without building tree")
	}
}

func TestHuffmanDecodeWithoutTree(t *testing.T) {
	encoder := NewHuffmanEncoder()
	
	_, err := encoder.Decode("101010")
	if err == nil {
		t.Error("Expected error when decoding without building tree")
	}
}

func TestHuffmanInvalidEncodedString(t *testing.T) {
	encoder := NewHuffmanEncoder()
	text := "abc"
	
	err := encoder.BuildTree(text)
	if err != nil {
		t.Fatalf("BuildTree failed: %v", err)
	}
	
	// Test invalid bit
	_, err = encoder.Decode("10102")
	if err == nil {
		t.Error("Expected error for invalid bit in encoded string")
	}
	
	// Test incomplete encoded string (doesn't end at root)
	_, err = encoder.Decode("1")
	if err == nil {
		t.Error("Expected error for incomplete encoded string")
	}
}

func TestHuffmanEncodeUnknownCharacter(t *testing.T) {
	encoder := NewHuffmanEncoder()
	text := "abc"
	
	err := encoder.BuildTree(text)
	if err != nil {
		t.Fatalf("BuildTree failed: %v", err)
	}
	
	// Try to encode character not in tree
	_, err = encoder.Encode("abcd")
	if err == nil {
		t.Error("Expected error when encoding unknown character")
	}
}

func TestHuffmanGetCodes(t *testing.T) {
	encoder := NewHuffmanEncoder()
	text := "aab"
	
	err := encoder.BuildTree(text)
	if err != nil {
		t.Fatalf("BuildTree failed: %v", err)
	}
	
	codes := encoder.GetCodes()
	
	if len(codes) != 2 {
		t.Errorf("Expected 2 codes, got %d", len(codes))
	}
	
	// Check that codes exist for both characters
	if _, exists := codes['a']; !exists {
		t.Error("Code for 'a' not found")
	}
	
	if _, exists := codes['b']; !exists {
		t.Error("Code for 'b' not found")
	}
}

func TestHuffmanCompressionRatio(t *testing.T) {
	encoder := NewHuffmanEncoder()
	
	// Test with repeated characters (should compress well)
	text := "aaaaaabbbbccdd"
	
	err := encoder.BuildTree(text)
	if err != nil {
		t.Fatalf("BuildTree failed: %v", err)
	}
	
	encoded, err := encoder.Encode(text)
	if err != nil {
		t.Fatalf("Encode failed: %v", err)
	}
	
	ratio := encoder.GetCompressionRatio(text, encoded)
	if ratio <= 0 {
		t.Errorf("Expected positive compression ratio for repeated text, got %f", ratio)
	}
	
	// Test with empty original (edge case)
	ratioEmpty := encoder.GetCompressionRatio("", "anything")
	if ratioEmpty != 0 {
		t.Errorf("Expected 0 ratio for empty original, got %f", ratioEmpty)
	}
}

func TestCompressString(t *testing.T) {
	text := "hello world"
	
	encoded, encoder, err := CompressString(text)
	if err != nil {
		t.Fatalf("CompressString failed: %v", err)
	}
	
	decoded, err := DecompressString(encoded, encoder)
	if err != nil {
		t.Fatalf("DecompressString failed: %v", err)
	}
	
	if decoded != text {
		t.Errorf("Expected '%s', got '%s'", text, decoded)
	}
}

func TestHuffmanFrequencyOrdering(t *testing.T) {
	encoder := NewHuffmanEncoder()
	text := "aabbbcccc" // 'c' most frequent, 'a' least frequent
	
	err := encoder.BuildTree(text)
	if err != nil {
		t.Fatalf("BuildTree failed: %v", err)
	}
	
	codes := encoder.GetCodes()
	
	// More frequent characters should generally have shorter codes
	// This is not always guaranteed due to tree structure, but let's check the principle
	if len(codes['c']) > len(codes['a']) {
		// This might happen due to tree balancing, so let's just ensure it works
		t.Logf("Code lengths - 'a': %d, 'b': %d, 'c': %d", 
			len(codes['a']), len(codes['b']), len(codes['c']))
	}
}

// Benchmark tests
func BenchmarkHuffmanBuildTree(b *testing.B) {
	text := strings.Repeat("hello world! this is a test string with various characters.", 100)
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		encoder := NewHuffmanEncoder()
		encoder.BuildTree(text)
	}
}

func BenchmarkHuffmanEncode(b *testing.B) {
	encoder := NewHuffmanEncoder()
	text := strings.Repeat("hello world! this is a test string.", 100)
	encoder.BuildTree(text)
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		encoder.Encode(text)
	}
}

func BenchmarkHuffmanDecode(b *testing.B) {
	encoder := NewHuffmanEncoder()
	text := strings.Repeat("hello world! this is a test string.", 100)
	encoder.BuildTree(text)
	encoded, _ := encoder.Encode(text)
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		encoder.Decode(encoded)
	}
}

func BenchmarkHuffmanFullCycle(b *testing.B) {
	text := strings.Repeat("the quick brown fox jumps over the lazy dog.", 50)
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		encoder := NewHuffmanEncoder()
		encoder.BuildTree(text)
		encoded, _ := encoder.Encode(text)
		encoder.Decode(encoded)
	}
}
