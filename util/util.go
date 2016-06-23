package util

import (
	srand "crypto/rand"
	"math/rand"
	"time"
)

const dictionary = "_0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func CreateRandomString() string {
	b := make([]byte, 32)
	l := len(dictionary)
	_, err := srand.Read(b)
	if err != nil {
		// fail back to insecure rand
		rand.Seed(time.Now().UnixNano())
		for i := range b {
			b[i] = dictionary[rand.Int() % l]
		}
	} else {
		for i, v := range b {
			b[i] = dictionary[v % byte(l)]
		}
	}

	return string(b)
}