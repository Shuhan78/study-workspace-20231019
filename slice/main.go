package main

import "fmt"

func main() {
	//array
	var number [2]int
	number[0] = 1
	number[1] = 2
	fmt.Println(len(number))
	fmt.Println(number[0])
	fmt.Println("--------------------------")

	//巡迴陣列
	grade := [4]int{70, 60, 90, 50}
	var sum int = 0
	for i := 0; i < len(grade); i++ {
		sum += grade[i]
	}
	fmt.Println(sum)
	fmt.Println("--------------------------")

	//multidimentional Array
	//create 2 arrays containing 2 integers
	var twoDim [2][2]int
	twoDim[0][0] = 2
	twoDim[0][1] = 3
	twoDim[1][0] = 2
	twoDim[1][1] = 3
	fmt.Println("twoDim: ", twoDim)
	fmt.Println("--------------------------")

	//create arrays containing 2 integers
	twoDim2 := [1][2]int{{2, 9}}
	fmt.Println("twoDim2: ", twoDim2)
	fmt.Println("--------------------------")

	//create slice from existing array
	arr := []int{23, 14, -4, -120}
	arrSlice := arr[1:3]
	fmt.Printf("arrslice: %v\n", arrSlice)
	fmt.Println("--------------------------")

	//Slice
	//方式一: 建立帶有資料的string slice，適合用於知道slice裡面的元素有哪些時
	people := []string{"steven", "nana", "mia"}
	fmt.Println(people)
	fmt.Println("--------------------------")

	//方式二: 透過make建立 空slice，適合用於對slice中特定位置進行操作時
	people1 := make([]int, 4)
	fmt.Println(people1)
	fmt.Println("--------------------------")

	//方式三: 建立空的slice
	var people2 []string
	people2 = append(people2, "hey")
	people2 = append(people2, "hey!")
	people2 = append(people2, "hey~")
	fmt.Println(len(people2))
	fmt.Println("--------------------------")

	//方式四: 大概知道需要多少元素，搭配append使用
	people3 := make([]int, 0, 2) //len=0 cap=2
	people3 = append(people3, 02)
	people3 = append(people3, 12)
	people3 = append(people3, 22)
	people3 = append(people3, 32)
	fmt.Println(people3)
	fmt.Println("--------------------------")

	a := [...]int{16, 12, 13, 14, 15}
	b := a[1:3] //從位置1取2個元素，長度為3-1 = 2
	c := a[2:3]
	d := a[:3]
	e := a[:]
	fmt.Println("a: ", a, "length: ", len(a), "capacity: ", cap(a))
	fmt.Println("b: ", b, len(b), cap(b))
	fmt.Println("c: ", c, len(c), cap(c))
	fmt.Println("d: ", d, len(d), cap(d))
	fmt.Println("e: ", e, len(e), cap(e))
	fmt.Println("--------------------------")

	//slice的操作會影響原有陣列，因為slice為pass-by-reference，而array為pass-by-value，所以array要注意pointer的問題
	b[0] = 22
	fmt.Println(b, len(b), cap(b))
	fmt.Println(a, len(a), cap(a))
	fmt.Println("--------------------------")

	aw := []int{25, 26, 24}
	fmt.Println("aw: ", aw, "cap:", cap(aw))
	fmt.Println("--------------------------")

	aw = append(aw, []int{-8, -52, 8, 52, 33, -66}...)
	fmt.Println("aw: ", aw, "cap:", cap(aw)) //cap = 10 ?????
	fmt.Println("--------------------------")

	for k, v := range aw {
		fmt.Println(k, v)
	}
	fmt.Println("cap:", cap(aw))
	fmt.Println("--------------------------")

	//只保留陣列aw前3個元素(為刪除)
	aw = append(aw[:0], aw[:3]...)
	for i, v := range aw {
		fmt.Println(i, v)
	}
	fmt.Println("aw: ", aw, "cap:", cap(a))
	fmt.Println("--------------------------")

	//取完aw[:1]的值再加上a[:3]
	aw = append(aw[:1], aw[:3]...)
	for i, v := range aw {
		fmt.Println(i, v)
	}
	fmt.Println("aw: ", aw, "cap:", cap(a))
	fmt.Println("--------------------------")

	aw = append([]int{-10}, aw...)
	for i, v := range aw {
		fmt.Println(i, v)
	}
	fmt.Println("aw: ", aw, "cap:", cap(a))
	fmt.Println("--------------------------")

	//copy
	a1 := []int{21, 11, -3, 107}
	a2 := []int{7, -2, -2}
	copy(a1, a2) //將a2複製到a1
	fmt.Println("a1:", a1)
	fmt.Println("a2:", a2)
	fmt.Println("--------------------------")

	b1 := []int{21, 11, -3, 107}
	b2 := []int{7, -2, -2}
	copy(b2, b1) //將b1複製到b2
	fmt.Println("b1:", b1)
	fmt.Println("b2:", b2)
	fmt.Println("--------------------------")

	//copy的參數必須是slice，如果array要使用copy，則需要傳遞[:]
	c1 := [4]int{21, 11, -3, 107}
	c2 := [3]int{7, -2, -2}
	copy(c2[:], c1[:]) //將c1複製到c2
	fmt.Println("c1:", c1)
	fmt.Println("c2:", c2)
	fmt.Println("--------------------------")

	//長度和容量均為0的slice並不等於nil
	qw := []int{}
	fmt.Println(qw, len(qw), cap(qw))

	if qw == nil {
		fmt.Println("garbage collection will collect it")
	} else {
		fmt.Println("clean it by myself")
	}

	//test conflict
	sli := []int{2, 4, -4}
	fmt.Println(sli)
}
