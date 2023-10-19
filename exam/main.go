package main

import "fmt"

func hello(s string) {
	fmt.Println("HELLO", s, "how r u ")
}

func sumUp(x, y int) int {
	result := x + y
	return result
}

func main() {
	hello("steven")
	sumUp(2, 3)
}
