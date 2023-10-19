package main

import "fmt"

type IPrint interface {
	MyPrint()
}

type IS1 struct {
	A, B int
	S    string
}

type IS2 struct {
	S string
}

func (i IS1) MyPrint() {
	fmt.Println(i.S)
}

func (i IS2) MyPrint() {
	fmt.Println(i.S)
}

func main() {

	var is1 IPrint
	s1 := IS1{A: 2, B: 5, S: "hola"}
	//is1.MyPrint()
	is1 = s1
	is1.MyPrint()
	//fmt.Println(is1.S)
	is1 = IS2{S: "HELLO"}
	is1.MyPrint()

}
