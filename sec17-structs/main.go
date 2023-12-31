package main

import (
	"fmt"
)

// When creating a struct it is best to write first the fields with the largest amount of bytes(float64,int64,int32...).
type person struct {
	firstname string
	lastname  string
	age       int
}

type job struct {
	person
	jobTitle           string
	companyName        string
	programingLanguage string
	yearsOfExperience  int
}

type personAge int // creates a user-defined type, that is underlying the type 'int'

func main() {
	person1 := person{
		firstname: "Panos",
		lastname:  "Vasilopoulos",
		age:       31,
	}

	person2 := person{
		firstname: "Efraim",
		lastname:  "Vasilopoulos",
		age:       23,
	}

	fmt.Printf("person1 is %s %s.\n", person1.firstname, person1.lastname)
	fmt.Printf("person2 is %s %s.\n", person2.firstname, person2.lastname)

	// Embeded structs

	// person1Resume := job{
	// 	person: person{
	// 		firstname: "Panos",
	// 		lastname:  "Vasilopoulos",
	// 		age:       31,
	// 	},
	// 	jobTitle:           "Programmer",
	// 	companyName:        "XGate Solutions",
	// 	programingLanguage: "Go",
	// 	yearsOfExperience:  2,
	// }
	// fmt.Printf("%s %s is a %s with %d years of experience in Software Development. He now works at '%s'.",
	// 	person1Resume.firstname, person1Resume.lastname, person1Resume.jobTitle, person1Resume.yearsOfExperience, person1Resume.companyName)

	person1Resume := job{
		person:             person1,
		jobTitle:           "Programmer",
		companyName:        "XGate Solutions",
		programingLanguage: "Go",
		yearsOfExperience:  2,
	}
	fmt.Printf("%s %s is a %s with %d years of experience in Software Development. He now works at '%s'.\n",
		person1Resume.firstname, person1Resume.lastname, person1Resume.jobTitle, person1Resume.yearsOfExperience, person1Resume.companyName)

	fmt.Println(person1Resume.age)
	fmt.Println()

	// Anonymus structs
	carOfPerson1 := struct {
		brand string
		color string
		model int
	}{
		brand: "Alfa Romeo",
		model: 147,
		color: "ciel",
	}
	fmt.Printf("%s drives a %s %s %d.\n", person1Resume.firstname, carOfPerson1.color, carOfPerson1.brand, carOfPerson1.model)
	fmt.Println()
	fmt.Println("--------------")

	// Composition
	var a personAge = 31 // Initialize a variable of type 'personAge'.
	fmt.Printf("Variable 'a' is of type: '%T' and it's valuse is: %d", a, a)

}
