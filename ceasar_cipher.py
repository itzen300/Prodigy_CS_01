def encrypt(text, shift):
    result = ""  # To store the encrypted result
    for ch in text:
        if 'A' <= ch <= 'Z':  # Encrypt uppercase letters
            result += chr((ord(ch) - ord('A') + shift) % 26 + ord('A'))
        elif 'a' <= ch <= 'z':  # Encrypt lowercase letters
            result += chr((ord(ch) - ord('a') + shift) % 26 + ord('a'))
        else:
            result += ch  # Keep non-alphabet characters unchanged
    return result

def decrypt(text, shift):
    result = ""  # To store the decrypted result
    for ch in text:
        if 'A' <= ch <= 'Z':  # Decrypt uppercase letters
            result += chr((ord(ch) - ord('A') - shift + 26) % 26 + ord('A'))
        elif 'a' <= ch <= 'z':  # Decrypt lowercase letters
            result += chr((ord(ch) - ord('a') - shift + 26) % 26 + ord('a'))
        else:
            result += ch  # Keep non-alphabet characters unchanged
    return result

if __name__ == "__main__":
    text = input("Enter text: ")
    shift = int(input("Enter shift amount: "))

    # Encrypt the text
    encrypted_text = encrypt(text, shift)
    print("Encrypted text:", encrypted_text)

    # Decrypt the text back to original
    decrypted_text = decrypt(encrypted_text, shift)
    print("Decrypted text:", decrypted_text)
