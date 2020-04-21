package crypto

import (
	"crypto/sha256"
	"crypto/sha512" // https://golang.org/pkg/crypto/
	"encoding/hex"
)

func GetSha256(input string) string {
	hasher := sha256.New()
	defer hasher.Reset()
	hasher.Write([]byte(input))
	return hex.EncodeToString(hasher.Sum(nil))
}

func GetSha512(input string) string {
	hasher := sha512.New()
	defer hasher.Reset()
	hasher.Write([]byte(input))
	return hex.EncodeToString(hasher.Sum(nil))
}
