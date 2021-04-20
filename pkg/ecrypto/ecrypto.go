package ecrypto

import (
	"crypto/md5"
	"crypto/sha512"
	"encoding/hex"
	"math/rand"
)

func MD5(key string) string {
	data := md5.Sum([]byte(key))
	return string(data[:])
}

func Sha512(key string) string {
	data := sha512.Sum512([]byte(key))
	return string(data[:])
}

func GenerateRandomString(n int) string {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "6b0d767330ea1b312e65a4965db85ae1"
	}
	return hex.EncodeToString(bytes)
}

func GenerateRandomHex(n int) string {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "6b0d767330ea1b312e65a4965db85ae1" // n = 16
	}
	return string(bytes[:])
}

func GeneratePassword(password string, salt string) string {
	combination := salt + password
	data := sha512.Sum512([]byte(combination))
	return string(data[:])
}
