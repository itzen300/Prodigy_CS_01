package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Function to encrypt text using Caesar cipher
func encrypt(text string, shift int) string {
	var result strings.Builder
	for _, ch := range text {
		if ch >= 'A' && ch <= 'Z' { // Encrypt uppercase letters
			result.WriteRune(((ch-'A'+rune(shift))%26 + 'A'))
		} else if ch >= 'a' && ch <= 'z' { // Encrypt lowercase letters
			result.WriteRune(((ch-'a'+rune(shift))%26 + 'a'))
		} else {
			result.WriteRune(ch) // Keep non-alphabet characters unchanged
		}
	}
	return result.String()
}

// Function to decrypt text using Caesar cipher
func decrypt(text string, shift int) string {
	var result strings.Builder
	for _, ch := range text {
		if ch >= 'A' && ch <= 'Z' { // Decrypt uppercase letters
			result.WriteRune(((ch-'A'-rune(shift)+26)%26 + 'A'))
		} else if ch >= 'a' && ch <= 'z' { // Decrypt lowercase letters
			result.WriteRune(((ch-'a'-rune(shift)+26)%26 + 'a'))
		} else {
			result.WriteRune(ch) // Keep non-alphabet characters unchanged
		}
	}
	return result.String()
}

func main() {
	// Set up input reader
	reader := bufio.NewReader(os.Stdin)

	// Read text from user
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text) // Remove any trailing newline or spaces

	// Read shift amount from user
	var shift int
	fmt.Print("Enter shift amount: ")
	fmt.Scanf("%d", &shift)

	// Encrypt and display the result
	encryptedText := encrypt(text, shift)
	fmt.Println("Encrypted text:", encryptedText)

	// Decrypt and display the result
	decryptedText := decrypt(encryptedText, shift)
	fmt.Println("Decrypted text:", decryptedText)
}
