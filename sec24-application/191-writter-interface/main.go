package main

// Resources at page 128

import (
	"fmt"
	"io"
	"os"
)

// Function signature of 'Decode':
// func (dec *Decoder) Decode(v interface{}) error

func main() {
	fmt.Println("Hello, friend.")
	fmt.Fprintln(os.Stdout, "Hello, friend.")
	io.WriteString(os.Stdout, "Hello, friend.")
	fmt.Println()

}
