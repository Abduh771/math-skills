package maths

import "math"

func Deviation(slice []float64) int {
	var result float64

	ss := Sqrt(slice)

	result = ss / float64(len(slice))

	result = math.Sqrt(result)

	result = math.Round(result)

	return int(result)
}
