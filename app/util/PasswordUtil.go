package util

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

var salt = "0123456789abcdef"

type PasswordUtil struct{}

func (c *PasswordUtil) Encrypt(plainStr string) (*string, error) {
	key := []byte(salt)

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	iv := make([]byte, aes.BlockSize)
	stream := cipher.NewCTR(block, iv)

	cipherStr := make([]byte, len(plainStr))
	stream.XORKeyStream(cipherStr, []byte(plainStr))

	encrypted := base64.StdEncoding.EncodeToString(cipherStr)

	return &encrypted, nil
}

func (c *PasswordUtil) Decrypt(encrypted string) (*string, error) {
	key := []byte(salt)

	cipherStr, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	iv := make([]byte, aes.BlockSize)
	stream := cipher.NewCTR(block, iv)

	result := make([]byte, len(cipherStr))
	stream.XORKeyStream(result, cipherStr)

	plainStr := string(result)

	return &plainStr, nil
}
