package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

// Md5Encode 小写
func Md5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	tempStr := h.Sum(nil)
	return hex.EncodeToString(tempStr)
}

// MD5Encode 大写
func MD5Encode(data string) string {
	return strings.ToUpper(Md5Encode(data))
}

// Encrypt 加密
func Encrypt(plaintext string, salt string) string {
	return Md5Encode(plaintext + salt)
}

// Decrypt 解密
func Decrypt(ciphertext string, salt string, password string) bool {
	return Encrypt(ciphertext, salt) == password
}
