package main

import (
	"fmt"
)

func main() {
	nameToAgeMap := map[string]int{
		"Panos":    31,
		"Aggeliki": 23,
		"Efraim":   23,
		"Dimitris": 21,
	}
	for k, v := range nameToAgeMap {
		fmt.Printf("%s is %d years old.\n", k, v)
	}
	fmt.Println()

	// Or i can access a value like below:
	fmt.Printf("Panos is %d years old.\n", nameToAgeMap["Panos"])

	// Will not panic, just prints the default value of a type if key was not found.
	fmt.Printf("Panos is %d years old.\n", nameToAgeMap["Panos12"])
	fmt.Println()

	fmt.Printf("%v\n", nameToAgeMap)
	fmt.Printf("%#v\n", nameToAgeMap)
	fmt.Println()

	// Creates a map with the use of make function:
	fruitsToAmmountMap := make(map[string]int)
	fmt.Println("Number of keys in the map:", len(fruitsToAmmountMap))
	// Adds element to the map
	fruitsToAmmountMap["Oranges"] = 10
	fruitsToAmmountMap["Apples"] = 12
	fruitsToAmmountMap["Bananas"] = 5
	fmt.Println("Number of keys in the map:", len(fruitsToAmmountMap))
	fmt.Println(fruitsToAmmountMap)

	// How I delete an element from a map
	delete(fruitsToAmmountMap, "Bananas")
	delete(fruitsToAmmountMap, "Apleeeees") // will not panic
	fmt.Printf("Map after deletion: %#v\n", fruitsToAmmountMap)
	fmt.Println()

	// The right way to check if a key exists in the map, using "comma, ok" idiom.
	_, ok := nameToAgeMap["Aggeliki"]
	if ok {
		fmt.Println("Name 'Aggeliki' exist in the map.")
	} else {
		fmt.Println("Name 'Aggeliki' does not exist in the map.")
	}
	fmt.Println()

	// The proper way for the above:
	if _, ok := nameToAgeMap["Aggeliki"]; !ok {
		fmt.Println("Name 'Aggeliki' does not exist in the map.")
	} else {
		fmt.Println("Name 'Aggeliki' exist in the map.")
	}

}
