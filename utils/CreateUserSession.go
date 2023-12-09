package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
)

func secureRandomString(seed string) (string, error) {
	salt, err := generateRandomSalt()
	if err != nil {
		return "", err
	}
	seedBytes := []byte(seed)
	combinedBytes := append(seedBytes, salt...)
	hash := sha256.Sum256(combinedBytes)
	hexString := hex.EncodeToString(hash[:])
	return hexString, nil
}

func generateRandomSalt() ([]byte, error) {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, errors.New("无法生成随机盐值")
	}

	return salt, nil
}

func CreateSession(id string) (string, error) {
	secureRandomStr, err := secureRandomString(fmt.Sprintf("%s", id))
	if err != nil {
		return "", err
	}
	return secureRandomStr, nil
}
