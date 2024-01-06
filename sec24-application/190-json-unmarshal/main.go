package main

// Resources at page 127

// Function signature for unmarshal:
// func Unmarshal(data []byte, v interface{}) error

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	Firstname string `json:"Firstname"`
	Lastname  string `json:"Lastname"`
	Age       int    `json:"Age"`
}

func main() {
	// JSON object
	s := `[{"Firstname":"Panos","Lastname":"Vasilopoulos","Age":31}, {"Firstname":"Efraim","Lastname":"Vasilopoulos","Age":23}]`

	// Converts 's' raw string literal to a slice of bytes.
	bs := []byte(s)
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)

	// var people []person
	people := []person{}

	err := json.Unmarshal(bs, &people)
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	fmt.Printf("Raw data: %v\n", people)
	fmt.Println()

	for i, v := range people {
		fmt.Printf("For person number %d:\n", i+1) // +1 for user friendly output.
		// fmt.Println(v)
		fmt.Println(v.Firstname, v.Lastname, v.Age)
		fmt.Println()
	}
}
