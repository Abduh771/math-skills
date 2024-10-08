package maths

func sumOfSquares(slice []float64) float64 {
	var mean float64 // sum of slice / elements
	var temp float64
	var result float64
	var length float64 = float64(len(slice))

	for i := range slice {
		mean += slice[i]
	}

	mean /= length

	for i := range slice {
		temp = slice[i] - mean
		temp *= temp
		result += temp
	}

	return result
}
