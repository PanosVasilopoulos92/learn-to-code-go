package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Go type definition for JSON object
type person struct {
	First string `jason:"First"`
	Age   int    `json:"Age"`
}

func main() {
	s := `[{"First":"James","Age":32},{"First":"MoneyPenny","Age":27},{"First":"M","Age":54}]`

	// We must convert 's' into a slice of bytes.
	bs := []byte(s)

	// We must declare a variable of the propriate type in order to unmarshal to.
	agents := []person{}
	fmt.Printf("%T\n", agents) // a slice of type person.

	err := json.Unmarshal(bs, &agents)
	if err != nil {
		log.Fatalf("Error while unmarshal: %s", err)
	}
	fmt.Println(agents)
	fmt.Println()

	for i, v := range agents {
		fmt.Println("Agent", i+1) // +1 for user friendly display.
		fmt.Printf("%v\n", v)
		fmt.Println()
	}
}
