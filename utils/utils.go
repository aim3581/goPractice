package utils

import "math"

func GetArrMidInd[T any](arr []T) int {
	len := len(arr)
	mid := math.Floor(float64(len) / 2)
	if len%2 == 0 {
		return int(mid) - 1
	} else {
		return int(mid)
	}
}
