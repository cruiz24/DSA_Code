"""
Advanced Trie Data Structure Implementation in Python
Supports insert, search, delete, and prefix operations

Time Complexities:
- Insert: O(m) where m is the length of the word
- Search: O(m) where m is the length of the word  
- Delete: O(m) where m is the length of the word
- StartsWith: O(p) where p is the length of the prefix

Space Complexity: O(ALPHABET_SIZE * N * M) where N is number of words and M is average length

Author: Hacktoberfest 2025 Contributor
Date: October 2025
"""

class TrieNode:
    """Node class for Trie data structure"""
    
    def __init__(self):
        # Dictionary to store children nodes
        self.children = {}
        # Flag to mark end of word
        self.is_end_of_word = False
        # Optional: store the complete word at this node
        self.word = None

class Trie:
    """
    Advanced Trie (Prefix Tree) implementation with comprehensive functionality
    """
    
    def __init__(self):
        """Initialize the Trie with root node"""
        self.root = TrieNode()
        self.word_count = 0
    
    def insert(self, word):
        """
        Insert a word into the trie
        
        Args:
            word (str): Word to insert
            
        Returns:
            bool: True if word was newly inserted, False if already existed
        """
        if not word:
            return False
            
        current = self.root
        
        # Traverse through each character
        for char in word.lower():
            if char not in current.children:
                current.children[char] = TrieNode()
            current = current.children[char]
        
        # Check if word already existed
        if current.is_end_of_word:
            return False
            
        # Mark end of word
        current.is_end_of_word = True
        current.word = word.lower()
        self.word_count += 1
        return True
    
    def search(self, word):
        """
        Search for a word in the trie
        
        Args:
            word (str): Word to search
            
        Returns:
            bool: True if word exists, False otherwise
        """
        if not word:
            return False
            
        current = self.root
        
        # Traverse through each character
        for char in word.lower():
            if char not in current.children:
                return False
            current = current.children[char]
        
        return current.is_end_of_word
    
    def starts_with(self, prefix):
        """
        Check if there are words with given prefix
        
        Args:
            prefix (str): Prefix to check
            
        Returns:
            bool: True if prefix exists, False otherwise
        """
        if not prefix:
            return True
            
        current = self.root
        
        # Traverse through each character
        for char in prefix.lower():
            if char not in current.children:
                return False
            current = current.children[char]
        
        return True
    
    def get_words_with_prefix(self, prefix):
        """
        Get all words that start with given prefix
        
        Args:
            prefix (str): Prefix to search
            
        Returns:
            list: List of words with the prefix
        """
        words = []
        if not prefix:
            return self.get_all_words()
            
        current = self.root
        
        # Navigate to prefix end
        for char in prefix.lower():
            if char not in current.children:
                return words
            current = current.children[char]
        
        # Collect all words from this point
        self._collect_words(current, prefix.lower(), words)
        return words
    
    def _collect_words(self, node, prefix, words):
        """
        Helper method to collect words recursively
        
        Args:
            node (TrieNode): Current node
            prefix (str): Current prefix
            words (list): List to collect words
        """
        if node.is_end_of_word:
            words.append(prefix)
        
        for char, child_node in node.children.items():
            self._collect_words(child_node, prefix + char, words)
    
    def delete(self, word):
        """
        Delete a word from the trie
        
        Args:
            word (str): Word to delete
            
        Returns:
            bool: True if word was deleted, False if not found
        """
        if not word or not self.search(word):
            return False
        
        def _delete_helper(node, word, index):
            """Recursive helper for deletion"""
            if index == len(word):
                # Mark end of word as False
                node.is_end_of_word = False
                node.word = None
                # Return True if node has no children (can be deleted)
                return len(node.children) == 0
            
            char = word[index]
            child_node = node.children[char]
            
            # Recursively delete
            should_delete_child = _delete_helper(child_node, word, index + 1)
            
            if should_delete_child:
                # Delete the child node
                del node.children[char]
                # Return True if current node should be deleted
                return (not node.is_end_of_word and 
                       len(node.children) == 0)
            
            return False
        
        _delete_helper(self.root, word.lower(), 0)
        self.word_count -= 1
        return True
    
    def get_all_words(self):
        """
        Get all words in the trie
        
        Returns:
            list: List of all words
        """
        words = []
        self._collect_words(self.root, "", words)
        return words
    
    def count_words(self):
        """
        Get count of words in trie
        
        Returns:
            int: Number of words
        """
        return self.word_count
    
    def count_words_with_prefix(self, prefix):
        """
        Count words with given prefix
        
        Args:
            prefix (str): Prefix to count
            
        Returns:
            int: Number of words with prefix
        """
        return len(self.get_words_with_prefix(prefix))
    
    def longest_common_prefix(self):
        """
        Find longest common prefix of all words
        
        Returns:
            str: Longest common prefix
        """
        if self.word_count == 0:
            return ""
        
        current = self.root
        prefix = ""
        
        # Traverse while there's only one child and it's not end of word
        while (len(current.children) == 1 and 
               not current.is_end_of_word):
            char = list(current.children.keys())[0]
            prefix += char
            current = current.children[char]
        
        return prefix
    
    def auto_complete(self, prefix, max_suggestions=10):
        """
        Auto-complete suggestions for given prefix
        
        Args:
            prefix (str): Prefix for suggestions
            max_suggestions (int): Maximum number of suggestions
            
        Returns:
            list: List of suggestions
        """
        suggestions = self.get_words_with_prefix(prefix)
        return suggestions[:max_suggestions]

def run_comprehensive_tests():
    """Run comprehensive tests for Trie implementation"""
    print("=== Advanced Trie Data Structure Tests ===\n")
    
    # Create trie instance
    trie = Trie()
    
    # Test 1: Basic Operations
    print("Test 1: Basic Insert and Search Operations")
    words = ["apple", "app", "application", "apply", "banana", "band", "bandana"]
    
    for word in words:
        result = trie.insert(word)
        print(f"Insert '{word}': {'New' if result else 'Already exists'}")
    
    print(f"\nTotal words in trie: {trie.count_words()}")
    
    # Search tests
    search_words = ["app", "apple", "apply", "orange", "ban"]
    print("\nSearch Results:")
    for word in search_words:
        result = trie.search(word)
        print(f"Search '{word}': {'Found' if result else 'Not found'}")
    
    print("\n" + "="*50 + "\n")
    
    # Test 2: Prefix Operations
    print("Test 2: Prefix Operations")
    prefixes = ["app", "ban", "xyz", "a"]
    
    for prefix in prefixes:
        exists = trie.starts_with(prefix)
        words_with_prefix = trie.get_words_with_prefix(prefix)
        count = trie.count_words_with_prefix(prefix)
        
        print(f"Prefix '{prefix}':")
        print(f"  Exists: {exists}")
        print(f"  Words: {words_with_prefix}")
        print(f"  Count: {count}")
        print()
    
    print("="*50 + "\n")
    
    # Test 3: Auto-complete
    print("Test 3: Auto-complete Functionality")
    test_prefixes = ["app", "ban", "a"]
    
    for prefix in test_prefixes:
        suggestions = trie.auto_complete(prefix, max_suggestions=5)
        print(f"Auto-complete for '{prefix}': {suggestions}")
    
    print("\n" + "="*50 + "\n")
    
    # Test 4: Deletion
    print("Test 4: Deletion Operations")
    delete_words = ["app", "application", "xyz"]
    
    for word in delete_words:
        result = trie.delete(word)
        print(f"Delete '{word}': {'Success' if result else 'Not found'}")
    
    print(f"\nWords remaining: {trie.count_words()}")
    print(f"All words: {trie.get_all_words()}")
    
    print("\n" + "="*50 + "\n")
    
    # Test 5: Advanced Features
    print("Test 5: Advanced Features")
    
    # Longest common prefix
    lcp = trie.longest_common_prefix()
    print(f"Longest common prefix: '{lcp}'")
    
    # Add more words for better LCP test
    new_trie = Trie()
    lcp_words = ["interspecies", "interstellar", "interstate"]
    for word in lcp_words:
        new_trie.insert(word)
    
    lcp2 = new_trie.longest_common_prefix()
    print(f"LCP for {lcp_words}: '{lcp2}'")

def performance_test():
    """Performance test with large dataset"""
    print("\n=== Performance Test ===")
    
    import time
    import random
    import string
    
    trie = Trie()
    
    # Generate test words
    def generate_word(length):
        return ''.join(random.choices(string.ascii_lowercase, k=length))
    
    test_words = [generate_word(random.randint(3, 10)) for _ in range(10000)]
    
    # Test insertion
    start_time = time.time()
    for word in test_words:
        trie.insert(word)
    insert_time = time.time() - start_time
    
    print(f"Inserted {len(test_words)} words in {insert_time:.3f} seconds")
    print(f"Final word count: {trie.count_words()}")
    
    # Test search
    search_words = random.sample(test_words, 1000)
    start_time = time.time()
    found_count = sum(1 for word in search_words if trie.search(word))
    search_time = time.time() - start_time
    
    print(f"Searched 1000 words in {search_time:.3f} seconds")
    print(f"Found: {found_count}/1000")

if __name__ == "__main__":
    run_comprehensive_tests()
    performance_test()
