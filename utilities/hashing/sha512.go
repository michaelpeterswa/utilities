package hashing

import (
	"crypto/sha512"
	"encoding/hex"
)

// StringSHA512 returns the SHA512 hash of a string as a hex-encoded string.
func StringSHA512(a string) string {
	shaBytes := sha512.Sum512([]byte(a))
	return hex.EncodeToString(shaBytes[:])
}
