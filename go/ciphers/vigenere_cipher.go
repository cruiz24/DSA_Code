// time complexity: O(n)
// space complexity: O(n)

package vigenere

// Encrypt encrypts plaintext using the Vigenère cipher with the given key
// The Vigenère cipher is a polyalphabetic substitution cipher that uses
// a repeating keyword to encrypt the message character by character
func Encrypt(plaintext, key string) string {
	if len(key) == 0 {
		return plaintext
	}

	// Normalize key to uppercase and remove non-alphabetic characters
	normalizedKey := normalizeKey(key)
	if len(normalizedKey) == 0 {
		return plaintext
	}

	var result []byte
	keyIndex := 0

	for _, char := range []byte(plaintext) {
		if 'A' <= char && char <= 'Z' {
			// Uppercase letter
			shift := int(normalizedKey[keyIndex%len(normalizedKey)] - 'A')
			encryptedChar := 'A' + (char-'A'+byte(shift))%26
			result = append(result, encryptedChar)
			keyIndex++
		} else if 'a' <= char && char <= 'z' {
			// Lowercase letter
			shift := int(normalizedKey[keyIndex%len(normalizedKey)] - 'A')
			encryptedChar := 'a' + (char-'a'+byte(shift))%26
			result = append(result, encryptedChar)
			keyIndex++
		} else {
			// Non-alphabetic character (preserve as-is)
			result = append(result, char)
		}
	}

	return string(result)
}

// Decrypt decrypts ciphertext using the Vigenère cipher with the given key
// It reverses the encryption process by subtracting the key shifts
func Decrypt(ciphertext, key string) string {
	if len(key) == 0 {
		return ciphertext
	}

	// Normalize key to uppercase and remove non-alphabetic characters
	normalizedKey := normalizeKey(key)
	if len(normalizedKey) == 0 {
		return ciphertext
	}

	var result []byte
	keyIndex := 0

	for _, char := range []byte(ciphertext) {
		if 'A' <= char && char <= 'Z' {
			// Uppercase letter
			shift := int(normalizedKey[keyIndex%len(normalizedKey)] - 'A')
			decryptedChar := 'A' + (char-'A'-byte(shift)+26)%26
			result = append(result, decryptedChar)
			keyIndex++
		} else if 'a' <= char && char <= 'z' {
			// Lowercase letter
			shift := int(normalizedKey[keyIndex%len(normalizedKey)] - 'A')
			decryptedChar := 'a' + (char-'a'-byte(shift)+26)%26
			result = append(result, decryptedChar)
			keyIndex++
		} else {
			// Non-alphabetic character (preserve as-is)
			result = append(result, char)
		}
	}

	return string(result)
}

// normalizeKey converts the key to uppercase and removes non-alphabetic characters
func normalizeKey(key string) string {
	var normalized []byte
	for _, char := range []byte(key) {
		if 'A' <= char && char <= 'Z' {
			normalized = append(normalized, char)
		} else if 'a' <= char && char <= 'z' {
			normalized = append(normalized, char-32) // Convert to uppercase
		}
		// Skip non-alphabetic characters
	}
	return string(normalized)
}

// AnalyzeFrequency analyzes the frequency of characters in the text
// This can be useful for cryptanalysis of Vigenère ciphers
func AnalyzeFrequency(text string) map[rune]int {
	frequency := make(map[rune]int)
	for _, char := range text {
		if ('A' <= char && char <= 'Z') || ('a' <= char && char <= 'z') {
			// Convert to uppercase for consistent analysis
			if 'a' <= char && char <= 'z' {
				char = char - 32
			}
			frequency[char]++
		}
	}
	return frequency
}

// EstimateKeyLength attempts to estimate the key length using the Kasiski examination method
// This is a simplified version for educational purposes
func EstimateKeyLength(ciphertext string, maxKeyLength int) []int {
	if maxKeyLength <= 0 {
		maxKeyLength = 20
	}

	var possibleLengths []int
	
	// Look for repeated sequences of at least 3 characters
	sequences := findRepeatedSequences(ciphertext, 3)
	
	// Calculate distances between repeated sequences
	distances := make(map[int]int)
	for _, positions := range sequences {
		if len(positions) > 1 {
			for i := 0; i < len(positions)-1; i++ {
				for j := i + 1; j < len(positions); j++ {
					distance := positions[j] - positions[i]
					// Find factors of the distance
					for k := 2; k <= maxKeyLength && k <= distance; k++ {
						if distance%k == 0 {
							distances[k]++
						}
					}
				}
			}
		}
	}

	// Sort by frequency of occurrence
	type keyLengthFreq struct {
		length int
		freq   int
	}
	var candidates []keyLengthFreq
	for length, freq := range distances {
		candidates = append(candidates, keyLengthFreq{length, freq})
	}

	// Sort by frequency (highest first)
	for i := 0; i < len(candidates)-1; i++ {
		for j := i + 1; j < len(candidates); j++ {
			if candidates[i].freq < candidates[j].freq {
				candidates[i], candidates[j] = candidates[j], candidates[i]
			}
		}
	}

	// Return the most likely key lengths
	for _, candidate := range candidates {
		possibleLengths = append(possibleLengths, candidate.length)
		if len(possibleLengths) >= 5 { // Return top 5 candidates
			break
		}
	}

	return possibleLengths
}

// findRepeatedSequences finds all repeated sequences of a given minimum length
func findRepeatedSequences(text string, minLength int) map[string][]int {
	sequences := make(map[string][]int)
	
	for i := 0; i <= len(text)-minLength; i++ {
		for length := minLength; length <= len(text)-i && length <= 10; length++ {
			sequence := text[i : i+length]
			if isAlphabetic(sequence) {
				sequences[sequence] = append(sequences[sequence], i)
			}
		}
	}

	// Filter to keep only sequences that appear more than once
	filtered := make(map[string][]int)
	for seq, positions := range sequences {
		if len(positions) > 1 {
			filtered[seq] = positions
		}
	}

	return filtered
}

// isAlphabetic checks if a string contains only alphabetic characters
func isAlphabetic(s string) bool {
	for _, char := range s {
		if !('A' <= char && char <= 'Z') && !('a' <= char && char <= 'z') {
			return false
		}
	}
	return true
}
