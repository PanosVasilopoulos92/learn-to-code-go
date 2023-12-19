package main

import (
	"fmt"
	"math/rand"
)

func main() {
	val1 := 10
	val2 := 48
	var name string = "Panos"

	if val1 > 11 {
		fmt.Println("Greater than 11.")
	} else {
		fmt.Println("Less than 11.")
	}

	if name == "Panos" {
		fmt.Println("It is Panos, indeed.")
	} else if name == "Panoos" {
		fmt.Println("It is Panoos, indeed.")
	} else {
		fmt.Println("Wrong.")
	}

	// --------- Logical operators ---------------- \\

	if val1 > 12 && name == "Panos" {
		fmt.Println("Value is greater than 12 and the name is Panos.")
	}

	if val1 < 12 || val1 == 12 {
		fmt.Println("Value is less or equal to 12.")
	}

	if name != "Panos" && val1 != 19 {
		fmt.Println("The name is not Panos and the value is not the number 19.")
	}

	// --------- "statement; statement" & "comma, ok" idioms ---------------- \\

	// "statement; statement" idiom
	if z := 2 * rand.Intn(val1); z >= val1 {
		fmt.Printf("%v is greater or equal to %v\n", z, val1)
	} else {
		fmt.Printf("%v is less than %v\n", z, val1)
	}

	// The variable z has scope only in the if statement.
	if z := 4; z >= val1 {
		fmt.Printf("%v is greater or equal to %v\n", z, val1)
	} else {
		fmt.Printf("%v is less than %v\n", z, val1)
	}
	// fmt.Print(z)

	//"comma, ok" idiom -- Evaluates a value and a Map.
	// if seconds, ok := arr1; ok {
	// 	return seconds
	// }

	// --------- switch statement ---------------- \\
	switch {
	case val1 < 12:
		fmt.Println("Value is less than 12.")
	case val1 > 12:
		fmt.Println("Value is less than 12.")
	case val1 == 12:
		fmt.Println("Value is equals to 12.")
	default:
		fmt.Println("Wrong integer input.")
	}

	switch name {
	case "Panos":
		fmt.Println("Name is Panos.")
	case "Aggeliki":
		fmt.Println("Name is Aggeliki.")
	case "Efraim":
		fmt.Println("Name is Efraim.")
	case "Dimitris":
		fmt.Println("Name is Dimitris.")
	default:
		fmt.Println("Wrong string input.")
	}

	switch {
	case val2 == 48:
		fmt.Println("val2 is 48")
		fallthrough // It will go to the next case no matter what.
	case val2 == 47:
		fmt.Println("val2 is 47 or because of the fallthrough, the program entered this case and printed this string.")
	}

	// --------- select statement ---------------- \\
	// select {
	// // case v1 := <-ch1:
	// // 	fmt.Println("value from channel 1", v1)
	// // case v2 := <-ch2:
	// // 	fmt.Println("value from channel 1", v2)
	// }

	// --------- for loops ---------------- \\
	for i := 0; i < 10; i++ {
		fmt.Printf("*\t")
	}
	fmt.Println()

	// As a while loop - Example 1
	for val1 < 14 {
		fmt.Printf("Again less than 14, it is %v\n", val1)
		val1 += 1
	}
	fmt.Println()

	// As a while loop - Example 2
	for val3 := 12; val3 < 20; {
		fmt.Printf("Again less than 20, it is %v\n", val3)
		val3++
	}
	fmt.Println()

	// break
	for {
		fmt.Printf("val2 = %v", val2)
		if val2 > 52 {
			break
		}
		val2++
	}
	fmt.Println()

	// continue
	for i := 0; i < 20; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Printf("%v is an even number.\n", i)
	}
	fmt.Println()

	// Nested loop
	for i := 0; i < 5; i++ {
		fmt.Printf("--outer loop number %d\n", i+1)
		for j := 0; j < 5; j++ {
			fmt.Printf("--inner loop %d of outer loop number %d\n", j+1, i+1)
		}
	}
	fmt.Println()

	// !!!! for range loops !!!!

	// iterating a slice data structure
	xi := []int{1, 2, 3, 4, 5, 6}
	for i, v := range xi {
		fmt.Printf("Index %d has value: %d\n", i, v)
	}

	// iterating a map data structure
	// initialize a map
	m := map[string]int{
		"Panos":    31,
		"Aggeliki": 23,
	}
	for k, v := range m {
		fmt.Printf("Key %s, has value: %d\n", k, v)
	}

}
