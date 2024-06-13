package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"io"
)

func EncryptAES(key, text []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	msg := pad(text)
	ciphertext := make([]byte, aes.BlockSize+len(msg))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], msg)

	return ciphertext, nil
}

func EncryptAESToHex(key, text []byte) (string, error) {
	ciphertext, err := EncryptAES(key, text)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(ciphertext), nil
}

func DecryptAES(key, ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)

	return unpad(ciphertext), nil
}

func DecryptAESFromHex(key []byte, cipherStr string) (string, error) {
	ciphertext, err := hex.DecodeString(cipherStr)
	if err != nil {
		return "", err
	}

	textBytes, err := DecryptAES(key, ciphertext)
	if err != nil {
		return "", err
	}

	return string(textBytes), nil
}
