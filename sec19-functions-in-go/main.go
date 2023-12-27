package main

import "fmt"

// func (r receiver) identifier(p parameter/s) (return/s) {code}

type dog struct {
	name string
}

type animal struct {
	dog
	has4Legs bool
}

// interfaces. If a value of one type implements a method that is inside an inteface, it automaticly also belongs to the type of the inteface.
type livingBeing interface {
	bark()
}

func saySomething(lb livingBeing) {
	lb.bark()
}

func main() {
	foo()
	displayFirstname("Panos")
	lastname := displayLastname("Vasilopoulos")
	fmt.Println(lastname)
	name, age := displayNameAndAge("Aggeliki", 23)
	fmt.Println(name, age)
	fmt.Println("---------")

	resultOfAddition := sum(1, 20, 33, 10)
	fmt.Println("Result =", resultOfAddition)
	fmt.Println()

	fmt.Println("----Defer statement-----")
	defer displayFirstBook()
	displaySecondBook()

	fmt.Println()
	dog1 := dog{
		name: "Lucky",
	}

	dog1.bark()
	fmt.Println()
	fmt.Println("----Polymorphism-----")

	// Polymorphism
	animal1 := animal{
		dog: dog{
			name: "Cerberus",
		},
		has4Legs: true,
	}

	animal1.bark()
	fmt.Println()

	saySomething(animal1)
	saySomething(dog1)
	saySomething(animal1)
	fmt.Println()

}

func foo() {
	fmt.Println("I am from foo.")
}

func displayFirstname(s string) {
	fmt.Println("My name is", s)
}

func displayLastname(s string) string {
	return fmt.Sprintln("My lastname is", s)
}

func displayNameAndAge(name string, age int) (string, int) {
	return fmt.Sprint(name, " has age of"), age
}

// variadic parameter. It treats 'ii' as a slice of type int. It MUST BE the LAST PARAMETER in order to work.
func sum(ii ...int) int {
	fmt.Println(ii)

	result := 0
	for _, v := range ii {
		result += v
	}
	return result
}

// defer statement
func displayFirstBook() {
	fmt.Println("The seagull.")
}

func displaySecondBook() {
	fmt.Println("The hungry bear.")
}

// attach a method to a type
func (d dog) bark() {
	fmt.Println("woof woof!")
}

// interfaces and polymorphism \\
/*
	An interface in Go defines a set of method signatures.
	Polymorphism is the ability of a value of a certain type to also be of another type.
	In Go, values can be of more than one type.
*/
func (a animal) bark() {
	fmt.Println("woof woof!")
}
