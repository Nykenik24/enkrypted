package util

import (
	"errors"
	"math/rand/v2"
)

// i = max, j = min
func RandInt(i, j int) int {
	if i < j {
		panic(errors.New("max less than min in randint"))
	}
	return rand.IntN(i-j) + j
}

func Choice[T any](S []T) T {
	return S[RandInt(len(S), 0)]
}
