package symmetric

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
)

func Des3ECBEncrypt(key []byte, text []byte) ([]byte, error) {
	// Create the cipher.Block interface for Triple DES encryption
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}

	// Generate the initialization vector (IV)
	iv := make([]byte, des.BlockSize)
	// You should use a random IV for each encryption operation.
	// For simplicity, we're using a zero-filled IV here.
	// Never reuse an IV with the same key.

	// Create the cipher.BlockMode interface for Triple DES encryption in CBC mode
	mode := cipher.NewCBCEncrypter(block, iv)

	// Pad plaintext to a multiple of the block size using PKCS7 padding
	paddedPlaintext := pad(text, block.BlockSize())

	// Create a buffer to hold the encrypted data
	ciphertext := make([]byte, len(paddedPlaintext))

	// Encrypt the padded plaintext
	mode.CryptBlocks(ciphertext, paddedPlaintext)

	return ciphertext, nil
}

func Des3ECBDecrypt(key []byte, text []byte) ([]byte, error) {
	// Create the cipher.Block interface for Triple DES encryption
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}

	// Generate the initialization vector (IV)
	iv := make([]byte, des.BlockSize)
	// You should use a random IV for each encryption operation.
	// For simplicity, we're using a zero-filled IV here.
	// Never reuse an IV with the same key.

	// Create the cipher.BlockMode interface for Triple DES decryption in CBC mode
	mode := cipher.NewCBCDecrypter(block, iv)

	// Create a buffer to hold the decrypted data
	decrypted := make([]byte, len(text))

	// Decrypt the ciphertext
	mode.CryptBlocks(decrypted, text)

	// Unpad the decrypted plaintext
	return unpad(decrypted), nil
}

// pad plaintext to a multiple of the block size using PKCS7 padding
func pad(plaintext []byte, blockSize int) []byte {
	padding := blockSize - (len(plaintext) % blockSize)
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(plaintext, padText...)
}

// unpad plaintext by removing PKCS7 padding
func unpad(plaintext []byte) []byte {
	padding := plaintext[len(plaintext)-1]
	return plaintext[:len(plaintext)-int(padding)]
}
