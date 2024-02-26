package utils

import (
	randCrypto "crypto/rand"
	"math/big"
	randMath "math/rand"
	"time"
)

var defaultLetters = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandomString(n int, allowedChars ...string) string {
	var letters string
	if len(allowedChars) == 0 {
		letters = defaultLetters
	} else {
		letters = allowedChars[0]
	}

	r := randMath.New(randMath.NewSource(time.Now().UnixNano())) //nolint:gosec
	var index int
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := randCrypto.Int(randCrypto.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			index = r.Intn(len(letters))
		} else {
			index = int(num.Int64())
		}

		ret[i] = letters[index]
	}

	return string(ret)
}
