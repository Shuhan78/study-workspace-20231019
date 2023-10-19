package main

import (
	"fmt"
	"image/color"
)

type Rectangle struct {
	wide, height float64
}

type ColorRec struct {
	Rectangle
	Color color.RGBA
}

func main() {

	rec := Rectangle{wide: 2.5, height: 5.9}
	fmt.Println(area(rec.height, rec.wide))
	fmt.Println(rec.area())

	//繼承效果
	var cr ColorRec
	cr.wide = 2
	cr.height = 9
	fmt.Println(cr.area())

}

// 定義矩形求面積的方法
func area(w, h float64) float64 {
	return w * h
}

// 定義矩形類型的方法
func (r Rectangle) area() float64 { //r 為Rectangle的接收器，作用為該方法的調用者，接收器名稱都自己定義，通堂是對應類型的首字母
	return r.wide * r.height
}
