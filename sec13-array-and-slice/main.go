package main

import (
	"fmt"
	"sort"
)

func main() {
	// -------------- Arrays ------------------- \\

	// declare a variable of type []int (Array)
	var arr1 [7]int
	// assign a value to index position zero
	arr1[0] = 12
	arr1[1] = 24
	fmt.Println(arr1)
	fmt.Printf("%T", arr1)
	fmt.Println()

	// array literal
	arr2 := [4]int{1, 2, 3, 4}
	fmt.Println(arr2)

	// array literal that has compiler count elements
	arr3 := [...]int{19, 8, 76, 32, 33}
	fmt.Println(arr3)

	fmt.Printf("arr2 has length: %v\n", len(arr2))
	fmt.Printf("arr1 has length: %v\n", len(arr1))
	fmt.Printf("arr3 has length: %v\n", len(arr3))
	fmt.Println()

	// -------------- Slices -------------------- \\

	s1 := []int{}
	fmt.Println(s1)

	// variadic parameter
	s1 = append(s1, 1, 2, 3, 4)
	fmt.Println(s1)

	s2 := []string{"Panos", "Aggeliki", "Efraim"}
	for i, v := range s2 {
		fmt.Printf("In index number %d, name is: %s\n", i, v)
	}
	fmt.Println()

	// copy the s2 slice to a new one: s3
	s3 := []string{}
	for i := 0; i < len(s2); i++ {
		s3 = append(s3, s2[i])
	}
	fmt.Println(s3)
	fmt.Println()

	s4 := []string{}
	s4 = append(s4, s3...) // It copies by value, so changes won't affect the original slice.
	fmt.Println("s4:", s4)

	s4[0] = "qwerto"
	fmt.Println("s4:", s4)
	fmt.Println()

	// A test to confirm that if I change the value of an index in one slice, it will not impact the other one.
	s2[0] = "abc"
	fmt.Println(s2)
	s3[2] = "Good"
	fmt.Println(s3)
	fmt.Println(s2)

	// delete element from a slice
	s5 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(s5)
	fmt.Println("-----------------")

	s5 = append(s5[:4], s5[5:]...)
	fmt.Println(s5)

	// make function
	s6 := make([]string, 0, 2)
	fmt.Println("slice 6:", s6)
	fmt.Println(len(s6))
	fmt.Println(cap(s6))
	fmt.Println()

	s6 = append(s6, "Panos", "Aggeliki", "something", "something else", "more")
	fmt.Println("slice 6 after append values:", s6)
	fmt.Println(len(s6))
	fmt.Println(cap(s6))

	// multidimensional slice
	s7 := []int{1, 2, 3, 4, 5}
	s8 := []int{9, 8, 7, 6}
	fmt.Println(s7)
	fmt.Println(s8)

	multiSlice1 := [][]int{s7, s8}
	fmt.Println("multidimensional slice:", multiSlice1)
	fmt.Println()

	// Find the median of the values inside a slice.
	fmt.Println("Find the median of the values inside a slice:")
	s9 := []float64{2, 4, 1, 9, 5, 8, 2}
	fmt.Println("slice before medianOne:", s9)
	fmt.Println("Median of slice s9:", medianOne(s9))
	fmt.Println("slice after medianOne:", s9)
	fmt.Println("-------------")
	s10 := []float64{200, 44, 11, 99, 55, 64, 12}
	fmt.Println("slice before medianTwo:", s10)
	fmt.Println("Median of slice s10:", medianTwo(s10))
	fmt.Println("slice after medianTwo:", s10)
	fmt.Println("-------------")

}

// This function will use the same underlying array so the sort function will
// affect the state of the slice.
func medianOne(s []float64) float64 {
	sort.Float64s(s)
	i := len(s) / 2
	if len(s)%2 == 1 {
		return s[i/2] // In order to get the element in the middle, if the total number of indexes in slice is an odd number.
	}
	return (s[i-1] + s[i]) / 2
}

// This function creates a new slice that will point to a new array, so the sort function will NOT
// affect the original slice.
func medianTwo(s []float64) float64 {
	newSlice := make([]float64, len(s))
	copy(newSlice, s)

	sort.Float64s(newSlice)
	i := len(newSlice) / 2
	if len(s)%2 == 1 {
		return s[i/2] // In order to get the element in the middle, if the total number of indexes in slice is an odd number.
	}
	return (s[i-1] + s[i]) / 2
}
