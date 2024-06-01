package shortener

import (
	"crypto/sha256"
	"math/big"
	"fmt"
	"github.com/itchyny/base58-go"
)

func sha256Of(input string) []byte {
	algo := sha256.New()
	algo.Write([]byte(input))
	return algo.Sum(nil)
}

// Interesting, base58 is more readable than base64
func base58Encoded(bytes []byte) string {
	encoding := base58.BitcoinEncoding
	encoded, err := encoding.Encode(bytes)
	if err != nil {
		panic(err)
	}
	return string(encoded)
}

func GenerateShortUrl(longURL string, userID string) string {
	hash  := sha256Of(longURL + userID)
	// why Uint64? 
	// base58: expecting a non-negative number
	num   := new(big.Int).SetBytes(hash).Uint64()
	short := base58Encoded([]byte(fmt.Sprintf("%d", num)))
	return short[:8]
}

