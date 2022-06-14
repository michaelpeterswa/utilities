package hashing

import (
	"crypto/md5"
	"encoding/hex"
)

// StringMD5 returns the MD5 hash of a string as a hex-encoded string.
//
func StringMD5(a string) string {
	md5Bytes := md5.Sum([]byte(a))
	return hex.EncodeToString(md5Bytes[:])
}
