package main

import (
	"errors"
	"fmt"
	"log"
	"math"
)

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v, %v, %v", se.lat, se.long, se.err)
}

func main() {
	result, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(result)
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		e := errors.New("negative number input")
		return 0, sqrtError{"123", "345", e}
	}
	return math.Sqrt(f), nil
}
