package main

// Sources located at page 86
import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

type person struct {
	firstname string
}

func (p person) writeOut(w io.Writer) error {
	_, err := w.Write([]byte(p.firstname))
	return err
}

func main() {
	b := bytes.NewBufferString("Hello! ") // Creates a new buffer.
	fmt.Println(b.String())

	b.WriteString("Gophers") // Adds this string to the existing buffer.
	fmt.Println(b.String())

	b.Reset()
	fmt.Println(b.String())
	// fmt.Println(b.Available())

	b.WriteString("Hi! My name is Panos.")
	fmt.Println(b.String())
	b.Reset()

	b.Write([]byte("I am learning Go language.😉"))
	fmt.Println(b.String())
	fmt.Println("--------")
	fmt.Println()

	// Writing to either a file or a buffer
	p := person{
		firstname: "Panos",
	}

	f, err := os.Create("output2.txt")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	defer f.Close()

	var b2 bytes.Buffer

	p.writeOut(f)
	p.writeOut(&b2)          // Points to b2
	fmt.Println(b2.String()) // In order to print to the console what b2 contains.

}
