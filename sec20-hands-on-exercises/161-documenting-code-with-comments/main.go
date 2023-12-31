package main

import "fmt"

// Sources in page 103.

/*
- $ go help doc
- $ go doc ... (in order to run an order)
Flags:
	-all
		Show all the documentation for the package.
	-c
		Respect case when matching symbols.
	-cmd
		Treat a command (package main) like a regular package.
		Otherwise package main's exported symbols are hidden
		when showing the package's top-level documentation.
	-short
		One-line representation for each symbol.
	-src
		Show the full source code for the symbol. This will
		display the full Go source of its declaration and
		definition, such as a function definition (including
		the body), type declaration or enclosing const
		block. The output may therefore include unexported
		details.
	-u
		Show documentation for unexported as well as exported
		symbols, methods, and fields.
*/

func main() {
	age := DisplayAge(31)
	fmt.Println(age)
}

// !! Comments should be written directly above the function, STARTING with the function's name. !!
// DisplayAge returns the age of a person in order to be able to be assigned as a string value.
func DisplayAge(age int) string {
	return fmt.Sprintf("My age is %d.", age)
}
