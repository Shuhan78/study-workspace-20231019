package main

import (
	"fmt"
)

func main() {
	//zero value
	var (
		name    string
		id      int
		weight  float64
		married bool
	)

	fmt.Printf("name: %s\n id: %d\n weight: %f\n married: %t\n", name, id, weight, married)

	//short declaration (only inside the function!!!)
	a := 22
	fmt.Printf("a : %d\n", a)

	//string
	//Normal String (Can not contain newlines, and can have escape characters like `\n`, `\t` etc)
	speech1 := "bourn is a very perosnable man. He is a greatest guy in the world who create a biscuts.\n"

	//Raw String (Can span multiple lines. Escape characters are not interpreted)
	speech2 := `bourn is a very perosnable man. 
	He is a greatest guy in the world who create a biscuts.\n`
	fmt.Printf("speech1: %s\nspeech2: %s\n", speech1, speech2)

	//rune, byte
	var (
		b rune = 'あ'
		c rune = 'X'
		d byte = 'A'
	)

	fmt.Printf("%c = %U, %c = %U, %c = %d\n", b, c, b, c, d, d)

	//complex
	//Both real and imaginary parts must be of the same floating-point type
	var real float64 = 2.78
	var real2 float64 = 9.88
	var real3 = complex(real, real2)
	fmt.Println(real3)

	var r1 = 2 + 5i
	var r2 = 4 + 9.6i

	fmt.Println(r1 + r2)

	//type conversion (NO implicit type conversion in Go.)
	t1 := 'A'
	t2 := int(t1)
	t3 := rune(t2)

	fmt.Println(t2)
	fmt.Println(t3)

}
