package main

import "fmt"

func main() {
	//Regular loop
	for i := 1; i <= 50; i++ {
		if i == 50 {
			fmt.Print(i)
		} else {
			fmt.Print(i, "+")
		}
	}

	fmt.Println("\n--------------------------")

	num := 1
	for ; num <= 15; num++ {
		if num%2 == 0 {
			fmt.Printf("%v ", num)
		}
	}

	fmt.Println("\n--------------------------")

	var j int = 1
	for ; j <= 20; j++ {
		if j%2 != 0 {
			fmt.Printf("%v ", j)
		}
	}

	fmt.Println("\n--------------------------")

	//While loop
	x := 3
	for x > 0 {
		fmt.Print(x, " ")
		x--
	}

	fmt.Println("\n--------------------------")
	//Infinite loop
	/*for true {
	      fmt.Println("haiya")
	  }
	*/

	//Range loop
	s := [...]string{"hola", "哈囉", "hello", "Привіт"}

	for i, v := range s {
		fmt.Printf("s[%d]:%s\n", i, v)
	}

	fmt.Println("--------------------------")

	for _, v := range s {
		fmt.Printf("%s\n", v)
	}
	// for i := 0; i < len(s); i++ {
	//  fmt.Printf("s[%d]:%s\n", i, s[i])
	// }

}
