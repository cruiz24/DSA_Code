package vigenere

import (
	"reflect"
	"testing"
)

func TestEncrypt(t *testing.T) {
	var vigenereTestData = []struct {
		description string
		plaintext   string
		key         string
		expected    string
	}{
		{
			"Basic Vigenère encryption with single character key",
			"A",
			"B",
			"B",
		},
		{
			"Basic Vigenère encryption with simple word",
			"HELLO",
			"KEY",
			"RIJVS",
		},
		{
			"Vigenère encryption with repeating key",
			"ATTACKATDAWN",
			"LEMON",
			"LXFOPVEFRNHR",
		},
		{
			"Vigenère encryption with mixed case",
			"Hello World",
			"Key",
			"Rijvs Uyvjn",
		},
		{
			"Vigenère encryption preserving spaces and punctuation",
			"Hello, World!",
			"SECRET",
			"Zincs, Pgvnu!",
		},
		{
			"Vigenère encryption with numbers and special characters",
			"Test123!@#",
			"ABC",
			"Tfut123!@#",
		},
		{
			"Empty plaintext",
			"",
			"KEY",
			"",
		},
		{
			"Empty key returns original text",
			"HELLO",
			"",
			"HELLO",
		},
		{
			"Key with non-alphabetic characters (normalized)",
			"HELLO",
			"K3E@Y!",
			"RIJVS",
		},
		{
			"Long text with short key",
			"THEQUICKBROWNFOXJUMPSOVERTHELAZYDOG",
			"CODE",
			"VVHUWWFODFRAPTRBLIPTUCYITHKINOCCFCJ",
		},
		{
			"Case insensitive key",
			"hello",
			"key",
			"rijvs",
		},
		{
			"Wrap around alphabet",
			"ZZZ",
			"ABC",
			"ZAB",
		},
	}

	for _, data := range vigenereTestData {
		t.Run(data.description, func(t *testing.T) {
			result := Encrypt(data.plaintext, data.key)
			if result != data.expected {
				t.Errorf("Expected %s, got %s for input '%s' with key '%s'",
					data.expected, result, data.plaintext, data.key)
			}
		})
	}
}

func TestDecrypt(t *testing.T) {
	var vigenereDecryptTestData = []struct {
		description string
		ciphertext  string
		key         string
		expected    string
	}{
		{
			"Basic Vigenère decryption",
			"RIJVS",
			"KEY",
			"HELLO",
		},
		{
			"Vigenère decryption with repeating key",
			"LXFOPVEFRNHR",
			"LEMON",
			"ATTACKATDAWN",
		},
		{
			"Vigenère decryption with mixed case",
			"Rijvs Uyvjn",
			"Key",
			"Hello World",
		},
		{
			"Vigenère decryption preserving punctuation",
			"Zincs, Pgvnu!",
			"SECRET",
			"Hello, World!",
		},
		{
			"Empty ciphertext",
			"",
			"KEY",
			"",
		},
		{
			"Empty key returns original text",
			"RIJVS",
			"",
			"RIJVS",
		},
	}

	for _, data := range vigenereDecryptTestData {
		t.Run(data.description, func(t *testing.T) {
			result := Decrypt(data.ciphertext, data.key)
			if result != data.expected {
				t.Errorf("Expected %s, got %s for ciphertext '%s' with key '%s'",
					data.expected, result, data.ciphertext, data.key)
			}
		})
	}
}

func TestEncryptDecryptRoundTrip(t *testing.T) {
	testCases := []struct {
		plaintext string
		key       string
	}{
		{"Hello World", "SECRET"},
		{"The quick brown fox jumps over the lazy dog", "CRYPTOGRAPHY"},
		{"123ABC!@#def", "KEY"},
		{"", "KEY"},
		{"TEXT", ""},
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZ", "ALPHABET"},
		{"abcdefghijklmnopqrstuvwxyz", "alphabet"},
	}

	for _, tc := range testCases {
		t.Run("Round trip test", func(t *testing.T) {
			encrypted := Encrypt(tc.plaintext, tc.key)
			decrypted := Decrypt(encrypted, tc.key)
			if decrypted != tc.plaintext {
				t.Errorf("Round trip failed: original='%s', key='%s', encrypted='%s', decrypted='%s'",
					tc.plaintext, tc.key, encrypted, decrypted)
			}
		})
	}
}

func TestNormalizeKey(t *testing.T) {
	testCases := []struct {
		description string
		input       string
		expected    string
	}{
		{
			"All uppercase letters",
			"HELLO",
			"HELLO",
		},
		{
			"All lowercase letters",
			"world",
			"WORLD",
		},
		{
			"Mixed case",
			"HeLLo",
			"HELLO",
		},
		{
			"With numbers",
			"KEY123",
			"KEY",
		},
		{
			"With special characters",
			"K!E@Y#",
			"KEY",
		},
		{
			"Mixed everything",
			"Ke3y!@#",
			"KEY",
		},
		{
			"Empty string",
			"",
			"",
		},
		{
			"Only non-alphabetic",
			"123!@#",
			"",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			result := normalizeKey(tc.input)
			if result != tc.expected {
				t.Errorf("Expected '%s', got '%s' for input '%s'",
					tc.expected, result, tc.input)
			}
		})
	}
}

func TestAnalyzeFrequency(t *testing.T) {
	testCases := []struct {
		description string
		text        string
		expected    map[rune]int
	}{
		{
			"Simple text",
			"HELLO",
			map[rune]int{'H': 1, 'E': 1, 'L': 2, 'O': 1},
		},
		{
			"Mixed case",
			"Hello",
			map[rune]int{'H': 1, 'E': 1, 'L': 2, 'O': 1},
		},
		{
			"With spaces and punctuation",
			"Hello, World!",
			map[rune]int{'H': 1, 'E': 1, 'L': 3, 'O': 2, 'W': 1, 'R': 1, 'D': 1},
		},
		{
			"Empty string",
			"",
			map[rune]int{},
		},
		{
			"Only non-alphabetic",
			"123!@#",
			map[rune]int{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			result := AnalyzeFrequency(tc.text)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, got %v for text '%s'",
					tc.expected, result, tc.text)
			}
		})
	}
}

func TestEstimateKeyLength(t *testing.T) {
	// Test with a known Vigenère cipher
	ciphertext := "LXFOPVEFRNHRLXFOPVEFRNHRLXFOPVEFRNHR" // "ATTACKATDAWNATTACKATDAWNATTACKATDAWN" with key "LEMON"
	result := EstimateKeyLength(ciphertext, 10)
	
	// The key length should be detected (5 for "LEMON")
	// Note: This is a simplified test - in reality, cryptanalysis is more complex
	if len(result) == 0 {
		t.Error("Expected some key length estimates, got none")
	}
	
	// Test edge cases
	emptyResult := EstimateKeyLength("", 10)
	if len(emptyResult) != 0 {
		t.Error("Expected no key length estimates for empty string")
	}
	
	shortResult := EstimateKeyLength("ABC", 10)
	if len(shortResult) != 0 {
		t.Error("Expected no key length estimates for very short string")
	}
}

func TestFindRepeatedSequences(t *testing.T) {
	text := "ABCABCABC"
	result := findRepeatedSequences(text, 3)
	
	// Should find "ABC" repeated
	if sequences, exists := result["ABC"]; !exists || len(sequences) < 2 {
		t.Error("Expected to find repeated sequence 'ABC'")
	}
	
	// Test with no repeated sequences
	noRepeats := "ABCDEFGHI"
	resultNoRepeats := findRepeatedSequences(noRepeats, 3)
	if len(resultNoRepeats) != 0 {
		t.Error("Expected no repeated sequences in unique string")
	}
}

func TestIsAlphabetic(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{"ABC", true},
		{"abc", true},
		{"AbC", true},
		{"ABC123", false},
		{"AB C", false},
		{"", true}, // Empty string is considered alphabetic
		{"ABC!", false},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := isAlphabetic(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %v for input '%s', got %v",
					tc.expected, tc.input, result)
			}
		})
	}
}

// Benchmark tests
func BenchmarkEncrypt(b *testing.B) {
	plaintext := "The quick brown fox jumps over the lazy dog. This is a test of the Vigenère cipher implementation."
	key := "CRYPTOGRAPHYKEY"
	
	for i := 0; i < b.N; i++ {
		Encrypt(plaintext, key)
	}
}

func BenchmarkDecrypt(b *testing.B) {
	ciphertext := "Dpa jozgx vrszh nux pqzma yrag jra pqbm dwg. Bqom ml u bale yv bqa Vigenère gmlrar mzjoazahbuemyh."
	key := "CRYPTOGRAPHYKEY"
	
	for i := 0; i < b.N; i++ {
		Decrypt(ciphertext, key)
	}
}

func BenchmarkAnalyzeFrequency(b *testing.B) {
	text := "This is a very long text that contains many characters for frequency analysis testing purposes."
	
	for i := 0; i < b.N; i++ {
		AnalyzeFrequency(text)
	}
}
