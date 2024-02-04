package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.Create("logs.txt")
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	log.SetOutput(f) // Sets the file where logs will be printed.

	f2, err := os.Open("no-file.txt")
	if err != nil {
		// fmt.Println(err)
		// log.Println(err)
		// log.Fatalln(err)
		log.Panicln(err)
		// log.SetOutput(f)
		// panic(err)
	}
	defer f2.Close()
}
