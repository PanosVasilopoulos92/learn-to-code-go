package main

import "fmt"

type person struct {
	firstname string
	age       int
}

func main() {
	p1 := person{
		firstname: "Panos",
		age:       31,
	}
	fmt.Println(p1)
	fmt.Println()

	p2 := person{
		firstname: "Aggeliki",
		age:       31,
	}
	fmt.Println(p2)
	fmt.Println()

	changePersonPtr(&p2, "Angel", 24)
	fmt.Println(p2)
	fmt.Println()

	p3 := changePerson(p1, "Efraim", 23)
	fmt.Println(p3)
	fmt.Println()

}

func changePerson(p person, s string, a int) person {
	p.firstname = s
	p.age = a
	return p
}

func changePersonPtr(p *person, s string, a int) {
	p.firstname = s
	p.age = a
}
