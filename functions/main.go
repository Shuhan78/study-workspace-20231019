package main

import (
	"fmt"
	"math"
)

func testString(msg string) {
	fmt.Println(msg)
}

func add(n int) {
	sum := (1 + n) * n / 2
	fmt.Println(sum)
}
func addUp(n1, n2 int) int {
	var sumUp int = n1 + n2
	fmt.Println(sumUp)

	return 10

}

func show(msg string) string {
	if msg == "hello" {
		return "嗨嗨"
	}
	return msg
}

func evaluate() (int, string) {
	return 5, "test"
}

func main() {

	//基礎函式語法
	testString("安安")
	fmt.Println("--------------------------")

	//計算 1+2+3+....max 的函示包裝
	add(50)
	fmt.Println("--------------------------")

	//基本return 運作

	fmt.Println(show("hello"))
	fmt.Println(show("hola"))
	fmt.Println("--------------------------")

	//單一回傳值
	x := addUp(58, 22)
	fmt.Println(x)
	fmt.Println("--------------------------")

	//多個回傳值
	fmt.Println(evaluate())

	var s int
	var t string
	s, t = evaluate()
	fmt.Println(s, t)
	fmt.Println("--------------------------")

	//functions with multiple return values & blank identifier _ (to ignore the variable)
	change, changeOfStockPrice, _ := changeInStockPrice(19, 20.1)

	if change < 0 {
		fmt.Printf("The stock price is decreased by $%.2f, which is %.2f%% of previous stock price", math.Abs(change), math.Abs(changeOfStockPrice))
	} else {
		fmt.Printf("The stock price is increased by $%.2f, which is %.2f%% of previous stock price", change, changeOfStockPrice)
	}

	fmt.Println("\n--------------------------")

	//functions return an error
	change2, changeOfStockPrice2, error := changeInStockPrice(0, 20.1)

	if error != nil {
		fmt.Println(error)
	} else {
		if change2 < 0 {
			fmt.Printf("The stock price is decreased by $%.2f, which is %.2f%% of previous stock price", math.Abs(change2), math.Abs(changeOfStockPrice2))
		} else {
			fmt.Printf("The stock price is increased by $%.2f, which is %.2f%% of previous stock price", change2, changeOfStockPrice2)
		}
	}

}

func changeInStockPrice(prevPrice, currentPrice float64) (change, changeOfStockPrice float64, err error) {
	if prevPrice == 0 {
		err := fmt.Errorf("previous price can not be zero")
		return 0, 0, err
	}
	change = currentPrice - prevPrice
	changeOfStockPrice = (change / prevPrice) * 100
	return
}
