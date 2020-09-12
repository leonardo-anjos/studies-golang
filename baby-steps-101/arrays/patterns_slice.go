package main

import "fmt"

func main() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)
	
	q = q[1:4]
	fmt.Println(q)
	
	q = q[:2] 
	fmt.Println(q)
	
	q = q[1:] 
	fmt.Println(q) 
	
}