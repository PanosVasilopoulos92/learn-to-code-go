package main

func main() {
	add(1, 2, 3, 4, 5)
}

func add(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}
