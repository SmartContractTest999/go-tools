package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/rand"
	"io"
)

// GenerateAesKey 生成一个AES秘钥
func GenerateAesKey() ([]byte, error) {
	key := make([]byte, 32) // AES-256 需要32字节长的密钥
	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		return nil, err
	}
	return key, nil
}

// pad 对数据进行填充，使其长度为块大小的倍数
func pad(src []byte) []byte {
	padding := aes.BlockSize - len(src)%aes.BlockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

// unpad 移除填充数据
func unpad(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}
