package hashing

import (
	"crypto/sha256"
	"encoding/hex"
)

// StringSHA256 returns the SHA256 hash of a string as a hex-encoded string.
//
// See https://michaelpeters.dev/posts/hashes-and-strings/ for performace notes.
func StringSHA256(a string) string {
	shaBytes := sha256.Sum256([]byte(a))
	return hex.EncodeToString(shaBytes[:])
}
