package utils

import "math/rand"

func NewSeededRNG(seed int64) *rand.Rand {
	return rand.New(rand.NewSource(seed))
}
