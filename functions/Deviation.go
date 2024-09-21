package maths

import "math"

func Deviation(slice []float64) int {
	var result float64

	sumSquares := sumOfSquares(slice)

	result = sumSquares / float64(len(slice))

	result = math.Sqrt(result)

	result = math.Round(result)

	return int(result)
}
