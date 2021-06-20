package main

import "fmt"

type age struct {
	age int
}

func (a age) myAge() {
	fmt.Printf("My Korean age is %d years old", a.age)
}

func main() {
	myAge := age{33}
	myAge.myAge()
}
