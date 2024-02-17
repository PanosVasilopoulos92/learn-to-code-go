package main

import (
	"errors"
	"fmt"
	"log"
	"math"
)

func main() {
	// _, err := Sqrt(-10)
	result, err := Sqrt(10)
	if err != nil {
		fmt.Println(result)
		log.Fatalln(err)
	}
	fmt.Printf("%.3g", result)
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("norgate math: square root of negative number") // Will create a custom error message.
	}
	return math.Sqrt(f), nil
}
