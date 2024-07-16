package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
)

func main() {
	key := []byte("aaaaaaaaaaaaaaaa")
	plaintext := []byte("This is a secret message!")

	ciphertext, err := encryptAES(plaintext, key)
	if err != nil {
		fmt.Println("Encryption error:", err)
		return
	}
	fmt.Println("Ciphertext:", hex.EncodeToString(ciphertext))

	decryptedText, err := decryptAES(ciphertext, key)
	if err != nil {
		fmt.Println("Decryption error:", err)
		return
	}
	fmt.Println("Decrypted text:", string(decryptedText))
}

// 使用 AES-128 加密
func encryptAES(plaintext, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	ciphertext := make([]byte, len(plaintext))
	stream := cipher.NewCTR(block, make([]byte, aes.BlockSize))
	stream.XORKeyStream(ciphertext, plaintext)
	return ciphertext, nil
}

// 使用 AES-128 解密
func decryptAES(ciphertext, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	plaintext := make([]byte, len(ciphertext))
	stream := cipher.NewCTR(block, make([]byte, aes.BlockSize))
	stream.XORKeyStream(plaintext, ciphertext)
	return plaintext, nil
}
