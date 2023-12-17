package main

import (
	"fmt"

	dependency1 "github.com/PanosVasilopoulos92/learn-to-code-go-dependency-1"
)

func main() {
	dependency1.FromVersion1_2_0()
	fmt.Println()

	ex1 := dependency1.SayWelcome()
	ex2 := dependency1.SayGoodBye()
	fmt.Println(ex1, ex2)
	fmt.Println()
	fmt.Println(dependency1.SayWelcome())
	fmt.Println(dependency1.SayGoodBye())

	s3 := dependency1.Leaving()
	fmt.Println(s3)

}
