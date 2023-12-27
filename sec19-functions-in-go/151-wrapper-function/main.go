package main

import (
	"fmt"
	"log"
	"os"
)

// Sources located at page 95

func main() {
	xb, err := readFile("poem.txt")
	if err != nil {
		log.Fatalf("Error : %s", err)
	}
	// fmt.Println(xb)
	fmt.Println(string(xb))
}

func readFile(filename string) ([]byte, error) {
	xb, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error in readFile func: %s", err)
	}
	return xb, nil
}
