package main

import (
	"fmt"
	"log"
	"os"
)

func init() {
	nf, err := os.Create("log.txt")
	if err != nil {
		fmt.Printf("error occurred while creating log file.")
	}
	log.SetOutput(nf)
}

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		// fmt.Printf("error occurred: '%s'\n", err)
		log.Printf("error occurred: '%s'\n", err)
		// log.Fatalln(err)
		// panic(err)
	}

	fmt.Println("Program finished...")
}
