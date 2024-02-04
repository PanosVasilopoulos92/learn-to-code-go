package main

import (
	"fmt"
	"log"
	"math"
)

func main() {
	result, err := sqrt(-4)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(result)
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		// return 0, errors.New("square root of negative number")
		return 0, fmt.Errorf("square root of negative number")
	}
	result := math.Sqrt(f)
	return result, nil
}
