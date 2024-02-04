package main

import "testing"

func TestAdd(t *testing.T) {
	type test struct {
		data   []int
		answer int
	}

	tests := []test{
		{data: []int{10, 20, 30, 40, 50, 20}, answer: 170},
		{data: []int{10, 20, 30, 40, 50, 21}, answer: 171},
		{data: []int{10, 20, 30, 40, 50, 22}, answer: 172},
		{[]int{10, 20, 30, 40, 50, 23}, 173}, // I can also create it this way.
	}

	for _, v := range tests {
		expected := v.answer
		result := add(v.data...)
		if result != expected {
			t.Error("Expected:", v.answer, "\tGot:", result)
		}
	}

}
