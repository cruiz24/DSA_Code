package data_structures

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
	"testing"
)

func TestNewTrie(t *testing.T) {
	trie := NewTrie()
	if trie.Root == nil {
		t.Error("Root should not be nil")
	}
	if trie.Size != 0 {
		t.Errorf("Expected size 0, got %d", trie.Size)
	}
	if !trie.IsEmpty() {
		t.Error("New trie should be empty")
	}
}

func TestInsertAndSearch(t *testing.T) {
	trie := NewTrie()
	words := []string{"cat", "car", "card", "care", "careful"}
	
	// Insert words
	for _, word := range words {
		trie.Insert(word)
	}
	
	// Test search
	for _, word := range words {
		if !trie.Search(word) {
			t.Errorf("Word '%s' should exist in trie", word)
		}
	}
	
	// Test non-existent words
	nonExistent := []string{"ca", "cards", "carefully"}
	for _, word := range nonExistent {
		if trie.Search(word) {
			t.Errorf("Word '%s' should not exist in trie", word)
		}
	}
	
	if trie.Size != len(words) {
		t.Errorf("Expected size %d, got %d", len(words), trie.Size)
	}
}

func TestStartsWith(t *testing.T) {
	trie := NewTrie()
	words := []string{"cat", "car", "card"}
	
	for _, word := range words {
		trie.Insert(word)
	}
	
	// Test existing prefixes
	prefixes := []string{"c", "ca", "car"}
	for _, prefix := range prefixes {
		if !trie.StartsWith(prefix) {
			t.Errorf("Prefix '%s' should exist", prefix)
		}
	}
	
	// Test non-existing prefixes
	nonPrefixes := []string{"d", "x", "care"}
	for _, prefix := range nonPrefixes {
		if trie.StartsWith(prefix) {
			t.Errorf("Prefix '%s' should not exist", prefix)
		}
	}
}

func TestInsertWithValue(t *testing.T) {
	trie := NewTrie()
	
	trie.InsertWithValue("apple", 100)
	trie.InsertWithValue("banana", 200)
	
	// Test retrieval
	if val, exists := trie.GetValue("apple"); !exists || val != 100 {
		t.Errorf("Expected value 100 for 'apple', got %v", val)
	}
	
	if val, exists := trie.GetValue("banana"); !exists || val != 200 {
		t.Errorf("Expected value 200 for 'banana', got %v", val)
	}
	
	// Test non-existent key
	if _, exists := trie.GetValue("orange"); exists {
		t.Error("Value should not exist for 'orange'")
	}
}

func TestDelete(t *testing.T) {
	trie := NewTrie()
	words := []string{"cat", "cats", "dog", "dodge"}
	
	for _, word := range words {
		trie.Insert(word)
	}
	
	originalSize := trie.Size
	
	// Delete existing word
	if !trie.Delete("cat") {
		t.Error("Delete should return true for existing word")
	}
	
	if trie.Search("cat") {
		t.Error("Deleted word should not exist")
	}
	
	if trie.Size != originalSize-1 {
		t.Errorf("Size should decrease after deletion")
	}
	
	// Ensure other words still exist
	if !trie.Search("cats") {
		t.Error("'cats' should still exist after deleting 'cat'")
	}
	
	// Delete non-existing word
	if trie.Delete("bird") {
		t.Error("Delete should return false for non-existing word")
	}
}

func TestFindWordsWithPrefix(t *testing.T) {
	trie := NewTrie()
	words := []string{"cat", "car", "card", "care", "careful", "dog"}
	
	for _, word := range words {
		trie.Insert(word)
	}
	
	// Test prefix "car"
	carWords := trie.FindWordsWithPrefix("car")
	expected := []string{"car", "card", "care", "careful"}
	sort.Strings(expected)
	
	if !reflect.DeepEqual(carWords, expected) {
		t.Errorf("Expected %v, got %v", expected, carWords)
	}
	
	// Test empty prefix (should return all words)
	allWords := trie.FindWordsWithPrefix("")
	if len(allWords) != len(words) {
		t.Errorf("Empty prefix should return all words")
	}
	
	// Test non-existing prefix
	noWords := trie.FindWordsWithPrefix("xyz")
	if len(noWords) != 0 {
		t.Error("Non-existing prefix should return empty slice")
	}
}

func TestAutoComplete(t *testing.T) {
	trie := NewTrie()
	words := []string{"cat", "car", "card", "care", "careful", "cargo"}
	
	for _, word := range words {
		trie.Insert(word)
	}
	
	// Test auto-complete with limit
	suggestions := trie.AutoComplete("car", 3)
	if len(suggestions) != 3 {
		t.Errorf("Expected 3 suggestions, got %d", len(suggestions))
	}
	
	// All suggestions should start with "car"
	for _, suggestion := range suggestions {
		if !strings.HasPrefix(suggestion, "car") {
			t.Errorf("Suggestion '%s' should start with 'car'", suggestion)
		}
	}
}

func TestLongestCommonPrefix(t *testing.T) {
	// Test with common prefix
	trie1 := NewTrie()
	words1 := []string{"flower", "flow", "flight"}
	for _, word := range words1 {
		trie1.Insert(word)
	}
	
	lcp1 := trie1.LongestCommonPrefix()
	if lcp1 != "fl" {
		t.Errorf("Expected 'fl', got '%s'", lcp1)
	}
	
	// Test with no common prefix
	trie2 := NewTrie()
	words2 := []string{"dog", "racecar", "car"}
	for _, word := range words2 {
		trie2.Insert(word)
	}
	
	lcp2 := trie2.LongestCommonPrefix()
	if lcp2 != "" {
		t.Errorf("Expected empty string, got '%s'", lcp2)
	}
	
	// Test empty trie
	trie3 := NewTrie()
	lcp3 := trie3.LongestCommonPrefix()
	if lcp3 != "" {
		t.Error("Empty trie should have empty LCP")
	}
}

func TestGetCount(t *testing.T) {
	trie := NewTrie()
	
	// Insert same word multiple times
	trie.Insert("hello")
	trie.Insert("hello")
	trie.Insert("hello")
	
	if count := trie.GetCount("hello"); count != 3 {
		t.Errorf("Expected count 3, got %d", count)
	}
	
	if count := trie.GetCount("world"); count != 0 {
		t.Errorf("Expected count 0 for non-existing word, got %d", count)
	}
	
	// Size should still be 1 (unique words)
	if trie.Size != 1 {
		t.Errorf("Expected size 1, got %d", trie.Size)
	}
}

func TestFindShortestUniquePrefix(t *testing.T) {
	trie := NewTrie()
	words := []string{"cat", "car", "card"}
	
	for _, word := range words {
		trie.Insert(word)
	}
	
	prefixes := trie.FindShortestUniquePrefix()
	
	// "cat" should have unique prefix "cat" (since "ca" is shared)
	if prefixes["cat"] != "cat" {
		t.Errorf("Expected 'cat' for cat, got '%s'", prefixes["cat"])
	}
	
	// "car" should have unique prefix "car" 
	if prefixes["car"] != "car" {
		t.Errorf("Expected 'car' for car, got '%s'", prefixes["car"])
	}
}

func TestCountWordsWithPrefix(t *testing.T) {
	trie := NewTrie()
	words := []string{"cat", "car", "card", "care", "dog"}
	
	for _, word := range words {
		trie.Insert(word)
	}
	
	// Count words with prefix "car"
	count := trie.CountWordsWithPrefix("car")
	if count != 3 { // car, card, care
		t.Errorf("Expected 3 words with prefix 'car', got %d", count)
	}
	
	// Count all words (empty prefix)
	totalCount := trie.CountWordsWithPrefix("")
	if totalCount != len(words) {
		t.Errorf("Expected %d total words, got %d", len(words), totalCount)
	}
}

func TestFindLongestWord(t *testing.T) {
	trie := NewTrie()
	words := []string{"cat", "car", "card", "care", "careful"}
	
	for _, word := range words {
		trie.Insert(word)
	}
	
	longest := trie.FindLongestWord()
	if longest != "careful" {
		t.Errorf("Expected 'careful', got '%s'", longest)
	}
	
	// Test empty trie
	emptyTrie := NewTrie()
	if emptyTrie.FindLongestWord() != "" {
		t.Error("Empty trie should return empty string")
	}
}

func TestFindWordsOfLength(t *testing.T) {
	trie := NewTrie()
	words := []string{"cat", "car", "card", "care"}
	
	for _, word := range words {
		trie.Insert(word)
	}
	
	// Find words of length 3
	length3 := trie.FindWordsOfLength(3)
	expected := []string{"car", "cat"}
	sort.Strings(expected)
	
	if !reflect.DeepEqual(length3, expected) {
		t.Errorf("Expected %v, got %v", expected, length3)
	}
	
	// Find words of length 4
	length4 := trie.FindWordsOfLength(4)
	expectedLength4 := []string{"card", "care"}
	sort.Strings(expectedLength4)
	
	if !reflect.DeepEqual(length4, expectedLength4) {
		t.Errorf("Expected %v, got %v", expectedLength4, length4)
	}
}

func TestClear(t *testing.T) {
	trie := NewTrie()
	words := []string{"cat", "car", "dog"}
	
	for _, word := range words {
		trie.Insert(word)
	}
	
	trie.Clear()
	
	if !trie.IsEmpty() {
		t.Error("Trie should be empty after clear")
	}
	
	if trie.Size != 0 {
		t.Errorf("Size should be 0 after clear, got %d", trie.Size)
	}
	
	// Test that words are actually gone
	for _, word := range words {
		if trie.Search(word) {
			t.Errorf("Word '%s' should not exist after clear", word)
		}
	}
}

// Benchmark tests
func BenchmarkTrieInsert(b *testing.B) {
	trie := NewTrie()
	words := generateWords(1000)
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		trie.Insert(words[i%len(words)])
	}
}

func BenchmarkTrieSearch(b *testing.B) {
	trie := NewTrie()
	words := generateWords(1000)
	
	for _, word := range words {
		trie.Insert(word)
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		trie.Search(words[i%len(words)])
	}
}

func BenchmarkTrieAutoComplete(b *testing.B) {
	trie := NewTrie()
	words := generateWords(1000)
	
	for _, word := range words {
		trie.Insert(word)
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		trie.AutoComplete("test", 10)
	}
}

// Helper function to generate test words
func generateWords(n int) []string {
	words := make([]string, n)
	prefixes := []string{"test", "example", "demo", "sample", "data"}
	
	for i := 0; i < n; i++ {
		prefix := prefixes[i%len(prefixes)]
		words[i] = fmt.Sprintf("%s%d", prefix, i)
	}
	
	return words
}
