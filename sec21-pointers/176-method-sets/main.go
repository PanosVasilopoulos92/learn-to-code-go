package main

import "fmt"

// Resources in page 121
/*
 The following code is for learning and understanding purposes, not for production code.
 For maintaining consistency in our code, the receivers of one type should stick with either POINTER or VALUE semantics - not both.
*/

type dog struct {
	name string
}

func (d dog) walk() {
	fmt.Println("My name is", d.name, "and I' m walking.")
}

func (d *dog) run() {
	d.name = "Roger" // When a variable use this function, the name of the dog will change to this value.
	fmt.Println("My name is", d.name, "and I' m running!")
}

type youngin interface {
	walk()
	run()
}

func youngRun(y youngin) {
	y.run()
	y.walk()
}

func main() {
	d1 := dog{name: "Lucky"}
	d2 := &dog{name: "Lucy"} // value of type pointer.
	fmt.Println(d1.name)
	fmt.Println(d2.name)

	d1.walk()
	d1.run()
	/*
		!! youngRun(d1) wont work because 'd1' is NOT of type pointer. If an interface implements
		a function that has as receiver a pointer type, every variable that want to use that function
		MUST BE of type pointer. !!
	*/
	// youngRun(d1) // will work because d2 is NOT of type pointer.
	fmt.Println()

	d2.walk()
	d2.run()
	fmt.Println()

	// Will execute every function exist in the 'youngin' interface.
	youngRun(d2)
}
