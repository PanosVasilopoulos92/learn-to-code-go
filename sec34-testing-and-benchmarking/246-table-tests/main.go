package main

func main() {
	add(10, 20, 30)
}

func add(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}
