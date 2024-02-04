package main

// Resources at page 144

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// Simple example of error handling
	n, err := fmt.Println("Hello!")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)
	fmt.Printf("%T\n", n)
	fmt.Println("-------")

	// Error handling when creating a file
	f, err := os.Create("names.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%T\n", f)

	defer f.Close()

	r := strings.NewReader("Panos, Efraim, Aggeliki, Dimitris")
	fmt.Printf("%T\n", r)

	io.Copy(f, r)

	// Opening the created file
	f, err = os.Open("names.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%T\n", f)
	defer f.Close()

	bs, err := io.ReadAll(f)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%T\n", bs)
	fmt.Println(string(bs))
	fmt.Println("-------")

	// Error handling using "panic"
	var answer1, answer2, answer3 string

	fmt.Print("Name: ")
	_, err = fmt.Scan(&answer1)
	if err != nil {
		panic(err)
	}

	fmt.Print("Favorite fruit: ")
	_, err = fmt.Scan(&answer2)
	if err != nil {
		panic(err)
	}

	fmt.Print("Favorite programming language: ")
	_, err = fmt.Scan(&answer3)
	if err != nil {
		panic(err)
	}

	fmt.Println(answer1, answer2, answer3)
	fmt.Println("-------")

}
