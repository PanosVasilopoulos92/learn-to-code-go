package main

import "fmt"

// Resources at page 136

type person struct {
	firstname string
}

func (p *person) speak(s string, s2 string) {
	fmt.Println(s, s2)
}

type human interface {
	speak(s string, s2 string)
}

func saySomething(h human, firstname string) {
	great := "Hello from"
	h.speak(great, firstname)
}

func main() {
	p1 := person{
		firstname: "Panos",
	}

	p2 := person{
		firstname: "Efraim",
	}
	saySomething(&p1, p1.firstname)

	saySomething(&p2, p2.firstname)
}
