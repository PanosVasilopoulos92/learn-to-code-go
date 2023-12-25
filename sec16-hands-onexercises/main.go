package main

import (
	"fmt"
)

func main() {
	// Exercise 49: Create a map that has a key of type string and value of type []string and display everything to console:
	nameToFevThingsMap := map[string][]string{
		"Panos":    {"Item1", "Item2", "Item3"},
		"Aggeliki": {"Item1", "Item2", "Item3", "Item4", "Item5"},
		"Efraim":   {"Item1", "Item2", "Item3", "Item4", "Item5", "Item6", "Item7", "Item8"},
		"Dimitris": {"Item1", "Item2", "Item3", "Item4"},
	}

	for k, v := range nameToFevThingsMap {
		fmt.Printf("The favorite items of %s are: ", k)
		for _, e := range v {
			fmt.Printf("%s ", e)
		}
		fmt.Println()
		fmt.Printf("So, %s has %d favorite items.\n", k, len(v))
		fmt.Println()
	}

	// Or else:
	nameToFevThingsMap2 := make(map[string][]string)
	nameToFevThingsMap2["Panos"] = []string{"Item1", "Item2", "Item3"}
	nameToFevThingsMap2["Aggeliki"] = []string{"Item1", "Item2", "Item3"}
	nameToFevThingsMap2["something"] = []string{`Item1`, `"Item2"`}
	fmt.Printf("The type of 'nameToFevThingsMap2' is: %T", nameToFevThingsMap2)
	fmt.Println()

	// delete a record from a map data structure.
	fmt.Println("map 'nameToFevThingsMap2' before deleting last element: ", nameToFevThingsMap2)

	delete(nameToFevThingsMap2, "something")
	fmt.Println("map 'nameToFevThingsMap2' after deleting last element: ", nameToFevThingsMap2)
	fmt.Println("-------------------")

	// Exercise 50: Count the frequency of appearance of the following words.
	s1 := []string{
		"coconut", "pineapple", "kiwi", "watermelon", "muskmelon", "papaya", "guava", "apricot", "nectarine", "date",
		"coconut", "pineapple", "kiwi", "watermelon", "muskmelon", "papaya", "guava", "apricot", "nectarine", "date",
		"lemon", "lime", "orange", "grapefruit", "tangerine", "clementine", "pomelo", "mandarin", "satsuma", "kumquat",
		"lemon", "lime", "orange", "grapefruit", "tangerine", "clementine", "pomelo", "mandarin", "satsuma", "kumquat",
		"grape", "raspberry", "blackberry", "strawberry", "blueberry", "loganberry", "gooseberry", "boysenberry", "raspberry", "mulberry",
		"grape", "raspberry", "blackberry", "strawberry", "blueberry", "loganberry", "gooseberry", "boysenberry", "raspberry", "mulberry",
		"artichoke", "broccoli", "cauliflower", "cucumber", "zucchini", "bell pepper", "eggplant", "tomato", "carrot", "onion",
		"artichoke", "broccoli", "cauliflower", "cucumber", "zucchini", "bell pepper", "eggplant", "tomato", "carrot", "onion",
		"potato", "sweet potato", "yam", "radish", "turnip", "beetroot", "garlic", "ginger", "mushroom", "cabbage",
		"potato", "sweet potato", "yam", "radish", "turnip", "beetroot", "garlic", "ginger", "mushroom", "cabbage",
		"celery", "spinach", "lettuce", "cabbage", "kale", "collard greens", "arugula", "watercress", "endive", "mustard greens",
		"celery", "spinach", "lettuce", "cabbage", "kale", "collard greens", "arugula", "watercress", "endive", "mustard greens",
	}

	m := map[string]int{}
	for _, v := range s1 {
		word := v
		m[word]++ // Every time it finds a same key-word, it just overwrites the corresponding value.
	}

	for k, v := range m {
		fmt.Println(k, v)
	}

}
