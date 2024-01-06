package main

// Resources at page 127

// Function signature for json.marshal:
// func Marshal(v interface{}) ([]byte, error)

import (
	"encoding/json"
	"fmt"
	"log"
)

// In order for JSON to access the value of the fields inside a struct, the fields must be exported(capital first letter).
// The name of the struct does not matter if it begins with capital or lower case letter.
type person struct {
	Firstname string
	Lastname  string
	Age       int
}

func main() {
	p1 := person{
		Firstname: "Panos",
		Lastname:  "Vasilopoulos",
		Age:       31,
	}
	p2 := person{
		Firstname: "Efraim",
		Lastname:  "Vasilopoulos",
		Age:       23,
	}

	persons := []person{
		p1,
		p2,
	}

	fmt.Println(persons)

	bs, err := json.Marshal(persons)
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	fmt.Println(string(bs))
}
