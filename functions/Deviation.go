package maths

import "math"

func Deviation(slice []float64) int {
	var res float64

	ss := Sqrt(slice)

	res = ss / float64(len(slice))

	res = math.Sqrt(res)

	res = math.Round(res)

	return int(res)
}
