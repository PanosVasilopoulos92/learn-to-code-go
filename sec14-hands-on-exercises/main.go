package main

import (
	"fmt"
)

func main() {
	// Creates an Array using Composite Literal:
	arr1 := [5]int{1, 2, 3, 4, 5}
	for i, v := range arr1 {
		fmt.Printf("Index %d has value: %d of type: %T\n", i, v, v)
	}
	fmt.Println("--------------")

	// Creates a slice using make function
	s1 := make([]int, 0, 10)
	s1 = append(s1, 42, 23, 44, 45, 46, 47, 48, 49, 50, 51)
	for i, v := range s1 {
		fmt.Printf("Index %d has value: %d of type: %T\n", i, v, v)
	}

	// Slicing a slice:
	s2 := make([]int, 5)
	copy(s2, s1[:5])
	fmt.Println(s2)

	s3 := make([]int, 5)
	copy(s3, s1[5:])
	fmt.Println(s3)

	s4 := make([]int, 5)
	copy(s4, s1[2:7])
	fmt.Println(s4)

	// Exercise 45:
	s5 := []int{42, 23, 44, 45, 46, 47, 48, 49, 50, 51}
	s6 := []int{56, 57, 58, 59, 60}
	s5 = append(s5, 52)
	fmt.Println(s5)
	fmt.Println()
	s5 = append(s5, 53, 54, 55)
	fmt.Println(s5)
	fmt.Println()
	// appends the values of slice s6 to slice s5.
	s5 = append(s5, s6...)
	fmt.Println(s5)
	fmt.Println()

	// Exercise 46: Delete an element from a slice:
	s7 := []int{21, 32, 43, 54, 65, 76, 87, 98}
	fmt.Println(s7)
	// Delete element in index 3
	s7 = append(s7[:3], s7[4:]...)
	fmt.Println("s7 after deletion of a value:", s7)
	fmt.Println()

	// Exercise 47:
	s8 := make([]string, 0, 10)
	s8 = append(s8, "Panos", "Aggeliki", "Efraim", "Maria", "everyone else")
	fmt.Println(len(s8))
	fmt.Println(cap(s8))
	fmt.Println(s8)
	fmt.Println()

	// Exercise 47:
	s9 := []string{"Oranges", "Apples", "Bananas", "Nectarines"}
	multiSlice2 := [][]string{s8, s9}
	// fmt.Println("multiSlice2:", multiSlice2)
	for i, v := range multiSlice2 {
		fmt.Println(i, v)
	}
	fmt.Println()

	// Testing multidimentional slice
	multiSlice1 := [2][3]string{}
	multiSlice1[0][0] = "Panos"
	multiSlice1[0][1] = "Aggeliki"
	multiSlice1[0][2] = "Efraim"
	multiSlice1[1][0] = "ok"
	multiSlice1[1][1] = "Ok"
	multiSlice1[1][2] = "OK"
	for i, v := range multiSlice1 {
		fmt.Println(i, v)
		// fmt.Printf("%c", v)
		for j, v := range multiSlice1[i] {
			fmt.Println(j, v)
		}
	}

}
