package main

import "fmt"

func main() {

	// if
	num := 20
	if num%2 == 0 {
		fmt.Printf("%d is a multiple of 2\n", num)
	}

	fmt.Println("--------------------------")

	var err error
	if err := fmt.Errorf("err 1"); err != nil {
		fmt.Println(err)
		//err = fmt.Errorf("err 2") 此變數只存在if邏輯裡
	}
	fmt.Println(err) //印出的值為var err

	fmt.Println("--------------------------")

	// if...else if with short statement
	if BMI := 20; BMI > 30 {
		fmt.Println("overweight")
	} else if BMI < 30 && BMI >= 20 {
		fmt.Println("slightly overweight")
	} else if BMI < 20 && BMI >= 13 {
		fmt.Println("normal")
	} else if BMI < 13 && BMI >= 5 {
		fmt.Println("fit")
	} else {
		fmt.Println("thin")
	}

	fmt.Println("--------------------------")

	//switch
	var score int = 89
	switch {
	case score >= 90:
		fmt.Println("A")
	case score >= 80:
		fmt.Println("B")
	case score >= 70:
		fmt.Println("C")
	case score >= 60:
		fmt.Println("D")
	default:
		fmt.Println("F")
	}

	fmt.Println(grader(59))

	fmt.Println("--------------------------")

	switch weather := "windy"; weather {
	case "sunny", "windy":
		fmt.Println("GO HIKING")
	default:
		fmt.Println("STAY A HOME")
	}

	fmt.Println("--------------------------")

	//break and continue
	//持續讓使用者輸入數字，計算總合，直到使用者輸入0為止

	var keyIn int
	sum := 0

	for {
		fmt.Print("請輸入數字: ")
		fmt.Scanln(&keyIn)
		if keyIn == 0 {
			break
		}
		sum += keyIn
		fmt.Println(sum)

	}

}

func grader(score uint) string {
	switch {
	case score >= 90:
		return "A"
	case score >= 80:
		return "B"
	case score >= 70:
		return "C"
	case score >= 60:
		return "D"
	default:
		return "F"
	}
}
