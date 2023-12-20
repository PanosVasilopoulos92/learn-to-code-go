package main

import (
	"fmt"
	"math/rand"
)

// Exercise 21: Print a string with func init() (before the func main() runs)
func init() {
	fmt.Printf("This is where initialization for my program occurs.\n\n")
}

func main() {

	// Exercise 18: Run a SHA-256 checksum on a file
	// shasum -a 256 ./<name of file> // If we are in the same folder.
	// shasum -a 256 <path to file>/<name of file>

	// Exercise 19:
	x := rand.Intn(261) // [0,n) -> Include 0, exclude n.
	fmt.Println(x)

	if x >= 0 && x <= 100 {
		fmt.Println("x is between 0 and 100.")
	} else if x > 100 && x <= 200 {
		fmt.Println("x is between 101 and 200.")
	} else if x > 200 && x <= 250 {
		fmt.Println("x is between 201 and 250.")
	} else {
		fmt.Println("x bigger than 250.")
	}
	fmt.Println()

	// Exercise 20: Same as above using a switch statement.
	y := rand.Intn(261) // [0,n) -> Include 0, exclude n.
	fmt.Println(y)

	switch {
	case y >= 0 && y <= 100:
		fmt.Println("y is between 0 and 100.")
	case y > 100 && y <= 200:
		fmt.Println("y is between 101 and 200.")
	case y > 200 && y <= 250:
		fmt.Println("y is between 201 and 250.")
	default:
		fmt.Println("y bigger than 250.")
	}
	fmt.Println()

	// Exercise 24: create a loop.
	for i := 0; i < 100; i++ {
		fmt.Printf("i = %d\n", i)
	}
	fmt.Println()

	for counter := 0; counter < 10; {
		fmt.Printf("Loop number %d\n", counter+1) // +1 for user friendly output.

		z := rand.Intn(261) // [0,n) -> Include 0, exclude n.
		fmt.Println(z)

		switch {
		case z >= 0 && z <= 100:
			fmt.Println("z is between 0 and 100.")
		case z > 100 && z <= 200:
			fmt.Println("z is between 101 and 200.")
		case z > 200 && z <= 250:
			fmt.Println("z is between 201 and 250.")
		default:
			fmt.Println("z bigger than 250.")
		}

		counter++
	}
	fmt.Println()

	fmt.Println("Infinite for loop.")
	counter := 0

	for {
		fmt.Printf("Loop number %d\n", counter+1) // +1 for user friendly output.

		e := rand.Intn(361) // [0,n) -> Include 0, exclude n.
		fmt.Println(e)

		switch {
		case e >= 0 && e <= 100:
			fmt.Println("e is between 0 and 100.")
		case e > 100 && e <= 200:
			fmt.Println("e is between 101 and 200.")
		case e > 200 && e <= 250:
			fmt.Println("e is between 201 and 250.")
		default:
			fmt.Println("e bigger than 250, so this loop does not count!")
			continue
		}

		counter += 1
		if counter == 11 {
			break
		}
	}
	fmt.Println()

	// Exercise 30: Iterate through a slice data structure with for loop.
	xi := []int{12, 14, 16, 18, 23}

	for i, v := range xi {
		fmt.Printf("Index %d has the value: %d\n", i, v)
	}
	fmt.Println()

	// Exercise 30: Iterate through a map data structure with for loop.
	m := map[string]string{
		"password": "12qw14",
		"username": "panos",
	}

	for k, v := range m {
		fmt.Printf("Key '%s' has value: '%s'\n", k, v)
	}
	fmt.Println()

	// Exercise 32: Check if a map has a specific value using comma, ok idiom.
	// assigns a value from the map data structure to a variable.
	username := m["username"]
	fmt.Println("username is " + username)
	fmt.Printf("username is of type %T", username)
	fmt.Println()

	// comma, ok idiom
	if _, ok := m["username"]; ok {
		fmt.Println("'panos' is in the map as a username.")
	} else {
		fmt.Println("'panos' is not in the map as a username.")
	}

}
