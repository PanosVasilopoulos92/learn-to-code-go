package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Firstname string
	Lastname  string
	Sayings   []string
}

func main() {
	p1 := Person{
		Firstname: "Panos",
		Lastname:  "Vasilopoulos",
		Sayings:   []string{"Hello friend", "Life is good only if you have dignity."},
	}

	fmt.Println(p1)

	bs, err := json.Marshal(p1)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(bs))

}
