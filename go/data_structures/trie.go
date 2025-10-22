package data_structures

import (
	"sort"
	"strings"
)

// TrieNode represents a node in the trie
type TrieNode struct {
	Children map[rune]*TrieNode
	IsEnd    bool
	Value    interface{} // Store associated value
	Count    int         // Frequency of this word
}

// Trie represents a prefix tree data structure
type Trie struct {
	Root *TrieNode
	Size int
}

// NewTrie creates a new trie
func NewTrie() *Trie {
	return &Trie{
		Root: &TrieNode{
			Children: make(map[rune]*TrieNode),
		},
	}
}

// Insert adds a word to the trie
func (t *Trie) Insert(word string) {
	t.InsertWithValue(word, nil)
}

// InsertWithValue adds a word with associated value
func (t *Trie) InsertWithValue(word string, value interface{}) {
	node := t.Root
	
	for _, char := range word {
		if node.Children[char] == nil {
			node.Children[char] = &TrieNode{
				Children: make(map[rune]*TrieNode),
			}
		}
		node = node.Children[char]
	}
	
	if !node.IsEnd {
		t.Size++
	}
	node.IsEnd = true
	node.Value = value
	node.Count++
}

// Search checks if a word exists in the trie
func (t *Trie) Search(word string) bool {
	node := t.findNode(word)
	return node != nil && node.IsEnd
}

// StartsWith checks if any word starts with the given prefix
func (t *Trie) StartsWith(prefix string) bool {
	return t.findNode(prefix) != nil
}

// Delete removes a word from the trie
func (t *Trie) Delete(word string) bool {
	return t.deleteHelper(t.Root, word, 0)
}

// GetValue retrieves the value associated with a word
func (t *Trie) GetValue(word string) (interface{}, bool) {
	node := t.findNode(word)
	if node != nil && node.IsEnd {
		return node.Value, true
	}
	return nil, false
}

// GetCount returns the frequency count of a word
func (t *Trie) GetCount(word string) int {
	node := t.findNode(word)
	if node != nil && node.IsEnd {
		return node.Count
	}
	return 0
}

// FindWordsWithPrefix returns all words with the given prefix
func (t *Trie) FindWordsWithPrefix(prefix string) []string {
	var result []string
	node := t.findNode(prefix)
	
	if node != nil {
		t.collectWords(node, prefix, &result)
	}
	
	sort.Strings(result)
	return result
}

// AutoComplete returns suggestions for auto-completion
func (t *Trie) AutoComplete(prefix string, maxResults int) []string {
	words := t.FindWordsWithPrefix(prefix)
	if len(words) > maxResults {
		return words[:maxResults]
	}
	return words
}

// LongestCommonPrefix finds the longest common prefix of all words
func (t *Trie) LongestCommonPrefix() string {
	if t.Size == 0 {
		return ""
	}
	
	var prefix strings.Builder
	node := t.Root
	
	for len(node.Children) == 1 && !node.IsEnd {
		for char, child := range node.Children {
			prefix.WriteRune(char)
			node = child
			break
		}
	}
	
	return prefix.String()
}

// GetAllWords returns all words in the trie
func (t *Trie) GetAllWords() []string {
	return t.FindWordsWithPrefix("")
}

// WordCount returns the total number of words in the trie
func (t *Trie) WordCount() int {
	return t.Size
}

// IsEmpty checks if the trie is empty
func (t *Trie) IsEmpty() bool {
	return t.Size == 0
}

// Clear removes all words from the trie
func (t *Trie) Clear() {
	t.Root = &TrieNode{
		Children: make(map[rune]*TrieNode),
	}
	t.Size = 0
}

// FindShortestUniquePrefix finds the shortest unique prefix for each word
func (t *Trie) FindShortestUniquePrefix() map[string]string {
	result := make(map[string]string)
	words := t.GetAllWords()
	
	for _, word := range words {
		prefix := t.findShortestUniquePrefix(word)
		result[word] = prefix
	}
	
	return result
}

// CountWordsWithPrefix counts words starting with prefix
func (t *Trie) CountWordsWithPrefix(prefix string) int {
	node := t.findNode(prefix)
	if node == nil {
		return 0
	}
	
	return t.countWords(node)
}

// FindLongestWord returns the longest word in the trie
func (t *Trie) FindLongestWord() string {
	longest := ""
	words := t.GetAllWords()
	
	for _, word := range words {
		if len(word) > len(longest) {
			longest = word
		}
	}
	
	return longest
}

// FindWordsOfLength returns all words of specific length
func (t *Trie) FindWordsOfLength(length int) []string {
	var result []string
	t.findWordsOfLength(t.Root, "", length, &result)
	sort.Strings(result)
	return result
}

// Helper methods

func (t *Trie) findNode(word string) *TrieNode {
	node := t.Root
	
	for _, char := range word {
		if node.Children[char] == nil {
			return nil
		}
		node = node.Children[char]
	}
	
	return node
}

func (t *Trie) deleteHelper(node *TrieNode, word string, index int) bool {
	if index == len(word) {
		if !node.IsEnd {
			return false
		}
		node.IsEnd = false
		node.Count = 0
		t.Size--
		return len(node.Children) == 0
	}
	
	char := rune(word[index])
	childNode := node.Children[char]
	
	if childNode == nil {
		return false
	}
	
	shouldDelete := t.deleteHelper(childNode, word, index+1)
	
	if shouldDelete {
		delete(node.Children, char)
		return !node.IsEnd && len(node.Children) == 0
	}
	
	return false
}

func (t *Trie) collectWords(node *TrieNode, prefix string, result *[]string) {
	if node.IsEnd {
		*result = append(*result, prefix)
	}
	
	for char, child := range node.Children {
		t.collectWords(child, prefix+string(char), result)
	}
}

func (t *Trie) findShortestUniquePrefix(word string) string {
	node := t.Root
	prefix := ""
	
	for i, char := range word {
		prefix += string(char)
		node = node.Children[char]
		
		// If this is a complete word or has only one child, it's unique
		if node.IsEnd || len(node.Children) <= 1 {
			if i == len(word)-1 || len(node.Children) == 0 {
				return prefix
			}
		}
	}
	
	return word
}

func (t *Trie) countWords(node *TrieNode) int {
	count := 0
	
	if node.IsEnd {
		count = 1
	}
	
	for _, child := range node.Children {
		count += t.countWords(child)
	}
	
	return count
}

func (t *Trie) findWordsOfLength(node *TrieNode, current string, targetLength int, result *[]string) {
	if len(current) == targetLength {
		if node.IsEnd {
			*result = append(*result, current)
		}
		return
	}
	
	if len(current) < targetLength {
		for char, child := range node.Children {
			t.findWordsOfLength(child, current+string(char), targetLength, result)
		}
	}
}
