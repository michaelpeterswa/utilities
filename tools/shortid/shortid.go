package shortid

import (
	"crypto/rand"
	"math/big"
)

const (
	Base58CharacterSet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
)

type ShortID struct {
	alphabet string
}

// NewShortID returns a new ShortID instance.
func NewShortID(alphabet string) *ShortID {
	return &ShortID{
		alphabet: alphabet,
	}
}

// A cryptographically secure, random, base58-encoded string.
func (s *ShortID) Generate(length int) (string, error) {
	id := ""
	for i := 0; i < length; i++ {
		nBig, err := rand.Int(rand.Reader, big.NewInt(int64(len(s.alphabet))))
		if err != nil {
			return "", err
		}
		id += string(s.alphabet[nBig.Int64()])
	}
	return id, nil
}
