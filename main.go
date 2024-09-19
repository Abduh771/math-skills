package main

import (
	"fmt"
	maths "functions/maths"
	"log"
)

func main() {
	data, err := maths.Slice()
	if err != nil {
		log.Fatalf("Error while reading file: %v", err)
	}

	average := maths.Average(data)
	median := maths.Median(data)
	variance := maths.Variance(data)
	deviation := maths.Deviation(data)

	// Print the results
	fmt.Printf("Average: %d\n", average)
	fmt.Printf("Median: %d\n", median)
	fmt.Printf("Variance: %d\n", variance)
	fmt.Printf("Standard Deviation: %d\n", deviation)
}
