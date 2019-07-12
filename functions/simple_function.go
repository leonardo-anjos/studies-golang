package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func sub(x int, y int) int {
	return x - y
}

func div(x int, y int) int {
	return x / y
}

func mult(x int, y int) int {
	return x * y
}

func main() {
	fmt.Println("sum =", add(100, 5))
	fmt.Println("sub =", sub(100, 5))
	fmt.Println("div =", div(100, 5))
	fmt.Println("mult =", mult(100, 5))
}