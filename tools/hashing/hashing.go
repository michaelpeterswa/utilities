package hashing

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
)

// StringMD5 returns the MD5 hash of a string as a hex-encoded string.
func StringMD5(a string) string {
	md5Bytes := md5.Sum([]byte(a))
	return hex.EncodeToString(md5Bytes[:])
}

// StringSHA256 returns the SHA256 hash of a string as a hex-encoded string.
//
// See https://michaelpeters.dev/posts/hashes-and-strings/ for performace notes.
func StringSHA256(a string) string {
	shaBytes := sha256.Sum256([]byte(a))
	return hex.EncodeToString(shaBytes[:])
}

// StringSHA512 returns the SHA512 hash of a string as a hex-encoded string.
func StringSHA512(a string) string {
	shaBytes := sha512.Sum512([]byte(a))
	return hex.EncodeToString(shaBytes[:])
}
