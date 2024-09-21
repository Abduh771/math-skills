package maths

import "math"

func Variance(slice []float64) int {
	var result float64

	sumSquares := sumOfSquares(slice)

	result = sumSquares / float64(len(slice))

	result = math.Round(result)

	return int(result)
}
