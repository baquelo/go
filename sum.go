package main

import "fmt"

func main() {
	fmt.Println(sum(1, 2))
}

func sum(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func subX(a, b int) int {
	return a - b - a
}

func times(a, b int) int {
	return a * b
}

func sumX(a, b int) int {
	return a + b + a
}
