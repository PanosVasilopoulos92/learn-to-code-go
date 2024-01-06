package main

// Resources at page 128

import (
	"fmt"
	"slices"
	"sort"
)

type person struct {
	firstname string
	age       int
}

// Another way for custom sort:
type byAge []person

func (a byAge) Len() int {
	return len(a)
}

func (a byAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a byAge) Less(i, j int) bool {
	return a[i].age < a[j].age
}

func main() {
	xi := []int{10, 4, 26, 9, 32, 1, 4, 6, 100}
	xs := []string{"Oranges", "Apples", "Lemons", "Coffee"}
	fmt.Println(xi)
	fmt.Println(slices.IsSorted(xi))
	fmt.Println()
	fmt.Println(xs)
	fmt.Println(slices.IsSorted(xs))
	fmt.Println()

	slices.Sort(xi)
	fmt.Println(xi)
	fmt.Println(slices.IsSorted(xi))
	fmt.Println()

	slices.Sort(xs)
	fmt.Println(xs)
	fmt.Println(slices.IsSorted(xs))
	fmt.Println("--------")

	// Custom sort
	p1 := person{firstname: "Panos", age: 31}
	p2 := person{firstname: "Aggeliki", age: 23}
	p3 := person{firstname: "Efraim", age: 23}
	p4 := person{firstname: "Dimitris", age: 22}

	people := []person{p1, p2, p3, p4}
	fmt.Println(people)
	fmt.Println()

	sort.SliceStable(people, func(i int, j int) bool {
		return people[i].firstname < people[j].firstname
	})
	fmt.Println(people)

	// In order to check if 'people' is sorted.
	fmt.Printf("Is sorted by firstname: %v\n", sort.SliceIsSorted(people, func(i, j int) bool {
		return people[i].firstname < people[j].firstname
	}))
	fmt.Println()

	// Another way for custom sort:
	peopleSortedByAge := byAge(people) // is of type function.
	fmt.Printf("type: %T\n", peopleSortedByAge)

	sort.Sort(peopleSortedByAge)

	fmt.Println("Sorted by age:")
	for i, v := range peopleSortedByAge {
		fmt.Printf("Position %v: %v\n", i, v)
	}
	fmt.Println()
}
