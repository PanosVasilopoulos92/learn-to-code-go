package main

// Sources located at page 82

import (
	"log"
	"os"
)

// type person struct{
// 	firstname string
// }

// func (p person) writeOut(w io.Writer) error{
// 	_, err := w.Write([]byte(p.firstname))
// 	return err
// }

func main() {
	f, err := os.Create("output.txt") // Creates a new file
	if err != nil {
		log.Fatalf("error %s in file '%s'", err, "output.txt")
	}
	defer f.Close()

	s := []byte("Hello gophers!") // The parenthesis is used to convert the string into a slice of bytes.

	_, err = f.Write(s) // Will write to the file the value of 's'.
	if err != nil {
		log.Fatalf("error %s", err)
	}

}
