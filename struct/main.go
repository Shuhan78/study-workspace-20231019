package main

import "fmt"

// struct的名稱及其成員屬性大小寫會影響accessibility
type Aircraft struct {
	Name           string
	Function       string
	Weight, Length float64
}

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	p      Person
	Salary int
}

type Student struct {
	Person
	School string
}

func main() {

	//宣告方法一
	var c1 Aircraft
	c1 = Aircraft{"Boeing", "commercial", 2000.4, 125}
	fmt.Println(c1)

	//宣告方法二
	c2 := Aircraft{Name: "Airbus", Function: "commercial", Weight: 3000.58, Length: 140.89}
	fmt.Println(c2)

	//struct為pass-by-value，可以建立pointer類型
	//var c3 *Aircraft
	c4 := new(Aircraft)
	c4.Name = "Lockheed"

	p1 := Person{Name: "steven", Age: 34}
	p2 := AddAge(p1)
	fmt.Println(p1)
	fmt.Println(p2)

	addAgePlusOne(&p1)
	fmt.Println(p1)

	//利用new()函數來初始化pointer，這時該變數已非儲存指標的位址，而為反解指標變數
	pp := new(Person)
	addAgePlusOne(pp)
	fmt.Println(pp)

	//struct combination
	e := Employee{p: Person{"steven", 34}, Salary: 38000}
	fmt.Println(e.p.Name)

	var s Student
	s.Name = "Kate"
	s.Age = 35
	s.School = "賴厝"
	fmt.Println(s)
}

func AddAge(p Person) Person {
	p.Age += 1
	return p
}

func addAgePlusOne(pp *Person) {
	pp.Age += 1
}
