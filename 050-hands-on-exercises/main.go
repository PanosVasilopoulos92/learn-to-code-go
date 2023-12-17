package main

import (
	"fmt"
)

var val4 int32 = 18
var val5 uint32 = 18_000

const val6 int32 = 31

func main() {
	val1 := "Value 1"
	var val2 string = "Value 2"
	var val3 float64 = 12.32
	fmt.Printf("\"%v\", \"%v\", \"%v\" are values initialized in block scope.\n", val1, val2, val3)
	fmt.Printf("\"%v\", \"%v\" are values initialized in package level-scope.\n", val4, val5)
	// val6 := 18 // The const will change but a warning will apear where const was first initialized.
	val5 := 99
	fmt.Println(val6)
	fmt.Println(val5)
}
