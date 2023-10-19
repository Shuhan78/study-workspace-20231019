package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {

	//adding third party languages, use go get to download the third party packages from remote repositories
	fmt.Println(quote.Hello())
	fmt.Println(quote.Glass())
	fmt.Println(quote.Go())
	fmt.Println(quote.Opt())

	//creating and managing custom packages

}
