package main

import "fmt"

func main() {

	//typed and untyped constants
	const a1 = 12       //untyped const
	const b1 = a1 * 2.8 //untyped const
	const c1 = 45.12    //untyped const
	const d1 = "yolo"
	const e1 = b1
	fmt.Println(a1, b1, c1, d1, e1)

	fmt.Println("--------------------------")
	const f1 int = 22 //typed const
	/*
		已宣告型別，無法轉換
		const g float64 = f //typed const
	*/
	fmt.Println(f1)

	var test = 25 / 2
	fmt.Printf("result is %v and the type is %T\n", test, test)
	var result = float64(25) / 2
	fmt.Println(result)

	fmt.Println("--------------------------")
	//iota

	const (
		a = 1 << iota
		b = iota
		c
		d
		e    = "hola"
		f, g = iota, 2 << iota
		h    = iota + 20
		i
		j         = -8
		k float64 = iota / 2
		_
		l
		m      = 40
		n      = iota >> 4
		o bool = false
		p      = iota
		q      = m
		r      = q
		s
		t         = iota / 2 
		u float64 = iota * 0.03
	)

	/*
		step 1: 不同 const 定義區塊互不干擾
		step 2: 忽略所有的空行和註釋行
		step 3: 沒有表達式的常數複製上一行
		step 4: 從首行開始，iota從0逐行＋1
		step 5: 替換所有iota
	*/
	fmt.Printf("a : %v, type: %T\n", a, a)
	fmt.Printf("b : %v, type: %T\n", b, b)
	fmt.Printf("c : %v, type: %T\n", c, c)
	fmt.Printf("d : %v, type: %T\n", d, d)
	fmt.Printf("e : %v, type: %T\n", e, e)
	fmt.Printf("f : %v, type: %T\n", f, f)
	fmt.Printf("g : %v, type: %T\n", g, g)
	fmt.Printf("h : %v, type: %T\n", h, h)
	fmt.Printf("i : %v, type: %T\n", i, i)
	fmt.Printf("j : %v, type: %T\n", j, j)
	fmt.Printf("k : %v, type: %T\n", k, k)
	fmt.Printf("l : %v, type: %T\n", l, l)
	fmt.Printf("m : %v, type: %T\n", m, m)
	fmt.Printf("n : %v, type: %T\n", n, n)
	fmt.Printf("o : %v, type: %T\n", o, o)
	fmt.Printf("p : %v, type: %T\n", p, p)
	fmt.Printf("q : %v, type: %T\n", q, q)
	fmt.Printf("r : %v, type: %T\n", r, r)
	fmt.Printf("s : %v, type: %T\n", s, s)
	fmt.Printf("t : %v, type: %T\n", t, t)
	fmt.Printf("u : %v, type: %T\n", u, u)

}
