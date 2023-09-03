package helpers

import (
	"math/rand"
	"time"
)

func RandomNumber(n int) int {
	// randomize the seed to make it more random
	rand.New(rand.NewSource(time.Now().UnixNano()))
	return rand.Intn(n)
}
