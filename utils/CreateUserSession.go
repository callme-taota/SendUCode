package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func secureRandomString(seed string) (string, error) {
	// 将种子字符串转换为字芀f数组
	seedBytes := []byte(seed)

	// 将种子和随机字节数组合并
	combinedBytes := append(seedBytes)

	// 使用SHA-256哈希函数生成哈希值
	hash := sha256.Sum256(combinedBytes)

	// 将哈希值转换为十六进制字符串
	hexString := hex.EncodeToString(hash[:])

	return hexString, nil
}

func CreateSession(id string) (string, error) {
	secureRandomStr, err := secureRandomString(fmt.Sprintf("%s", id))
	if err != nil {
		return "", err
	}
	return secureRandomStr, nil
}
