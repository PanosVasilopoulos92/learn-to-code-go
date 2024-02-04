package main

// Resources at page 146

import (
	"fmt"

	"github.com/PanosVasilopoulos92/learn-to-code-go/sec32-and-sec33-writing-documentation/dog"
)

// Animal struct is used to create animal objects that contain a name and an age.
type animal struct {
	name string
	age  int
}

func main() {
	flufy := animal{
		name: "Flufy",
		age:  9,
	}

	flufyAgeInHumanYears := dog.DogYears(flufy.age)

	fmt.Println(flufyAgeInHumanYears)

}
