#include <stdio.h>
#include <string.h>

// Function to encrypt the text using Caesar cipher
void encrypt(char text[], int shift) {
    for (int i = 0; i < strlen(text); i++) {
        char ch = text[i];
        
        // Encrypt uppercase letters
        if (ch >= 'A' && ch <= 'Z') {
            text[i] = ((ch - 'A' + shift) % 26) + 'A';
        }
        // Encrypt lowercase letters
        else if (ch >= 'a' && ch <= 'z') {
            text[i] = ((ch - 'a' + shift) % 26) + 'a';
        }
    }
}

// Function to decrypt the text using Caesar cipher
void decrypt(char text[], int shift) {
    for (int i = 0; i < strlen(text); i++) {
        char ch = text[i];
        
        // Decrypt uppercase letters
        if (ch >= 'A' && ch <= 'Z') {
            text[i] = ((ch - 'A' - shift + 26) % 26) + 'A';
        }
        // Decrypt lowercase letters
        else if (ch >= 'a' && ch <= 'z') {
            text[i] = ((ch - 'a' - shift + 26) % 26) + 'a';
        }
    }
}

int main() {
    char text[100];
    int shift;
    
    printf("Enter text: ");
    fgets(text, sizeof(text), stdin);
    text[strcspn(text, "\n")] = 0;  // Remove newline character
    
    printf("Enter shift amount: ");
    scanf("%d", &shift);

    // Encrypt the text
    encrypt(text, shift);
    printf("Encrypted text: %s\n", text);

    // Decrypt the text back to original
    decrypt(text, shift);
    printf("Decrypted text: %s\n", text);

    return 0;
}
