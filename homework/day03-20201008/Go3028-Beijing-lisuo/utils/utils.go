package utils

import (
	"crypto/rand"
	"math/big"
)

func GenId() (res int64) {
	// gen a random number in [0, 999999999999)
	result, _ := rand.Int(rand.Reader, big.NewInt(999999999999))
	return result.Int64()
}
