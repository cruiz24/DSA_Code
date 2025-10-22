# Vigenère Cipher Implementation in Go

This implementation provides a robust and well-tested Vigenère cipher with additional cryptanalysis features.

## Overview

The Vigenère cipher is a method of encrypting alphabetic text by using a simple form of polyalphabetic substitution. It was first described by Giovan Battista Bellaso in 1553, but later misattributed to Blaise de Vigenère in the 19th century.

## Features

- **Encryption/Decryption**: Full Vigenère cipher implementation
- **Case Preservation**: Maintains original case of input text
- **Non-alphabetic Preservation**: Preserves spaces, numbers, and punctuation
- **Key Normalization**: Automatically handles mixed-case keys and filters non-alphabetic characters
- **Frequency Analysis**: Character frequency analysis for cryptanalysis
- **Key Length Estimation**: Kasiski examination for estimating key length
- **Comprehensive Testing**: Extensive test coverage including edge cases and benchmarks

## Algorithm Details

**Time Complexity**: O(n) where n is the length of the input text
**Space Complexity**: O(n) for the output string

The algorithm works by:
1. Normalizing the key (uppercase, alphabetic characters only)
2. For each alphabetic character in the input:
   - Calculate the shift based on the corresponding key character
   - Apply the shift with modular arithmetic (wrapping around the alphabet)
   - Preserve the original case
3. Non-alphabetic characters are preserved unchanged

## Usage

```go
package main

import (
    "fmt"
    "github.com/yourusername/DSA_Code/go/ciphers"
)

func main() {
    plaintext := "Hello, World!"
    key := "SECRET"
    
    // Encrypt
    encrypted := vigenere.Encrypt(plaintext, key)
    fmt.Printf("Encrypted: %s\n", encrypted) // Output: "Zincs, Pgvnu!"
    
    // Decrypt
    decrypted := vigenere.Decrypt(encrypted, key)
    fmt.Printf("Decrypted: %s\n", decrypted) // Output: "Hello, World!"
    
    // Frequency analysis
    freq := vigenere.AnalyzeFrequency("HELLO WORLD")
    fmt.Printf("Frequency: %v\n", freq) // Character frequency map
    
    // Estimate key length (cryptanalysis)
    keyLengths := vigenere.EstimateKeyLength(encrypted, 10)
    fmt.Printf("Possible key lengths: %v\n", keyLengths)
}
```

## Security Note

The Vigenère cipher, while historically significant, is not secure by modern standards and should only be used for educational purposes. It's vulnerable to:
- Frequency analysis when the key is short relative to the message
- Kasiski examination for key length determination
- Index of coincidence attacks

For real-world applications, use modern encryption algorithms like AES.

## Testing

Run the complete test suite:
```bash
go test -v vigenere_cipher.go vigenere_cipher_test.go
```

Run benchmarks:
```bash
go test -bench=. vigenere_cipher.go vigenere_cipher_test.go
```

## Educational Value

This implementation is ideal for:
- Learning about classical cryptography
- Understanding polyalphabetic substitution ciphers
- Practicing cryptanalysis techniques
- Demonstrating the evolution from classical to modern cryptography

## Author

Created for Hacktoberfest 2025 - Contributing to open source cryptography education.

## References

- [Vigenère Cipher - Wikipedia](https://en.wikipedia.org/wiki/Vigen%C3%A8re_cipher)
- [Cryptanalysis of the Vigenère Cipher](https://en.wikipedia.org/wiki/Cryptanalysis_of_the_Vigen%C3%A8re_cipher)
- [Kasiski Examination](https://en.wikipedia.org/wiki/Kasiski_examination)
