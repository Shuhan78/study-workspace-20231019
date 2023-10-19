package main

import "fmt"

func main() {
	//算術運算 Arithmetic operators
	//指定預算 Assignment operator =, +=, -=, *=, /=
	x := 3
	x += 5
	fmt.Println(x)
	//單元運算 Unary Operators ++, --
	x++
	fmt.Println(x)
	//比較運算 Comparison operators >, <, ==, >=, <=, !=
	if x > 5 {
		fmt.Println("great")
	} else {
		fmt.Println("woops")
	}

	var outcome bool = 3 == 5
	fmt.Println(outcome)
	//邏輯運算 Logical operators &&, ||, !

	var result bool = true && false
	fmt.Println(result)

	if x > 3 || result == true {
		fmt.Println("ok")
	}

	if y := 7; y < 9 && result == false {
		fmt.Println("ohoh")
	} else {
		fmt.Println("make no sense")
	}

}
