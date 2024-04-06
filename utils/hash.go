package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

// Calculate the hash of the content
func GenerateContentHash(content string) string {
	hash := sha256.Sum256([]byte(content))
	return hex.EncodeToString(hash[:])
}
