package main

import (
	"fmt"
	"time"
)

func Hello(s string) {
	fmt.Println("hello", s)
}

func main() {
	go Hello("steven")
	time.Sleep(1 * time.Second)
	fmt.Println("the end")

}
