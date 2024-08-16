package helpers

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"math/big"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz0123456789"

// GenerateRandomString generates a random string of specified length
func GenerateUniqueString() string {
	length := 16
	// Get a timestamp for uniqueness
	timestamp := time.Now().UnixNano()
	timestampString := fmt.Sprintf("%s%d", charset, timestamp)
	randomString := make([]byte, length)
	for i := range randomString {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(length)))
		if err != nil {
			panic(err)
		}
		randomString[i] = timestampString[num.Int64()]
	}
	return hex.EncodeToString(randomString)
}
