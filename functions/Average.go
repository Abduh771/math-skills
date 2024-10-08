package maths

import (
	"math"
)

func Average(slice []float64) int {
	var result int
	var temp float64
	var length float64 = float64(len(slice))

	for i := range slice {
		temp += slice[i]
	}
	temp = temp / length

	result = int(math.Round(temp))

	return result
}
