package maths

import (
	"math"
	"sort"
)

func Median(slice []float64) int {
	var temp float64
	var i int = len(slice) / 2 // get the pivot of the slice of numbers

	sort.Float64s(slice)

	if len(slice)%2 == 0 { // is number of slice is even?
		temp = (slice[i-1] + slice[i]) / 2
		return int(math.Round(temp))
	}

	return int(slice[i])
}
