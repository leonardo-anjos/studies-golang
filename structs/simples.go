package main

import "fmt"

type user struct {
	name string
	age int
	hight float32
}

type event struct {
	day, month, year int
}

func main() {
	// user
	fmt.Println(user{"leo", 23, 1.77})

	// event
	e := event{18, 6, 1996}
	e.day = 12
	e.month = 7
	e.year = 2019
	fmt.Println(e.day, e.month, e.year)	
}
