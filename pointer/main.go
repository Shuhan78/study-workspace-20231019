package main

import "fmt"

func multiple(Pointer *float64) {
	var sum float64
	sum = *Pointer * 5.2
	fmt.Println("multiple value", sum)
}

func main() {
	//1.建立存放資料的變數
	var x int = 22
	fmt.Println("原來的資料: ", x)

	//2.取得記憶體的位址
	fmt.Println("變數x的記憶體位址:", &x)

	//3.建立存放 x 的記憶體位址的指標變數，*資料型態
	var xPointer *int = &x
	fmt.Println(xPointer)

	//4.反解指標變數，*指標變數名稱
	var xDeference = *xPointer
	fmt.Println(xDeference)

	/*fmt.Println("請輸入一段話")
	  var msg string
	  fmt.Scanln(&msg) //&msg就是指傳遞輸入字串的指標變數
	  fmt.Println(msg)
	*/

	var n float64 = 12
	//方法一
	// var nPointer *float64 = &n
	// multiple(nPointer)

	//方法二
	multiple(&n)
	fmt.Println("main function value: ", n)

	a := [...]int{1, 2, 3, 4, 5}
	b := &a
	fmt.Println(a[0], a[1])
	fmt.Println(b[0], b[1])
}
