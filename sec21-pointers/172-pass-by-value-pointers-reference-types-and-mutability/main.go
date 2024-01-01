package main

// Sources in page 114

import "fmt"

func main() {
	a := 12
	fmt.Println(a)
	intDelta(&a)
	fmt.Println(a)
	fmt.Println()

	xi := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(xi)
	sliceDelta(xi)
	fmt.Println(xi)
	fmt.Println()

	m := map[string]int{
		"Panos":    31,
		"Aggeliki": 23,
	}
	fmt.Println(m)
	mapDelta(m, "Aggeliki")
	fmt.Println(m)
	fmt.Println()

}

// The passed argument will have to be a pointer when the type of parameter is '*<name of type>.'
func intDelta(n *int) {
	*n = 44
}

func sliceDelta(ii []int) {
	ii[0] = 10
}

func mapDelta(md map[string]int, name string) {
	md[name]++ // Adds one year to selected person. If imput name does not exist in the map, it will automaticaly CREATE it.
}
