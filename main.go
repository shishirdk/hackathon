package main

import "fmt"

func main() {
	a := addTwoNumber(10, 6)
	fmt.Println("First GoLang Code abd total is :", a)

}

func addTwoNumber(a, b int) int {
	one := a
	two := b
	total := one + two
	return total
}
