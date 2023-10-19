package main

import "fmt"

func main() {
	//map在元素給值前必須初始化
	//方法一: make
	m := make(map[string]int)
	m["n1"] = 22
	m["n2"] = 66
	m["n3"] = -5
	fmt.Println(m)
	delete(m, "n2")
	fmt.Println(m)

	value, ok := m["n2"]
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("the element is gone")
	}

	//方法二
	var m1 = map[string]int{
		"A": 200,
		"B": 100,
		"C": 50,
	}

	print(m1)

	//方法三
	m2 := map[int]string{
		1: "go",
		2: "backward",
		3: "right",
		4: "left",
	}

	for k, v := range m2 {
		fmt.Println(k, v)
	}

	//錯誤方法->宣告後無初始化
	var m3 map[string]int
	m3["go"] = 1
	print(m3)

}

func print(m map[string]int) {
	for k, v := range m {
		fmt.Println(k, v)
	}
}
