package maths

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

var (
	errEmpty = errors.New("empty file")
	errArgs  = errors.New("wrong input arguments")
)

func Slice() ([]float64, error) {
	var slice []float64

	arg := os.Args[1:]

	if len(arg) != 1 {
		return slice, errArgs
	}

	file := arg[0]

	f, err := os.Open(file)
	if err != nil {
		return slice, fmt.Errorf("error: %v, while opening file: %s", err, file)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		value, err := strconv.Atoi(line)
		if err != nil {
			return slice, fmt.Errorf("error while converting line: %v", err)
		}

		slice = append(slice, float64(value))
	}
	if err := scanner.Err(); err != nil {
		return slice, fmt.Errorf("error while reading file: %v", err)
	}

	if len(slice) == 0 {
		return slice, errEmpty
	}

	return slice, nil
}
