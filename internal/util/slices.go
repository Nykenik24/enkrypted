package util

import "fmt"

func RemoveByIndex[S any](slice []S, i int) ([]S, error) {
	if i < 0 || i >= len(slice) {
		return nil, fmt.Errorf("index %d out of range", i)
	}
	return append(slice[:i], slice[i+1:]...), nil
}
