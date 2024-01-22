package main

// Resources at page 129

import (
	"encoding/json"
	"fmt"
	"log"
)

type user struct {
	First string
	Age   int
}

func main() {
	u1 := user{
		First: "James",
		Age:   32,
	}

	u2 := user{
		First: "MoneyPenny",
		Age:   27,
	}

	u3 := user{
		First: "M",
		Age:   54,
	}

	users := []user{u1, u2, u3}
	fmt.Println(users)
	fmt.Println()

	bs, err := json.Marshal(users)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	fmt.Println(string(bs)) // Must be cast to string in order to be displayed appropriately.
}
