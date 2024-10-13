package main

import "fmt"

func main() {
	var num int

	fmt.Println("Enter the value: ")
	fmt.Scan(&num)
	EvenOrOdd(num)
	Foreach()

}

func EvenOrOdd(a int) {
	if a%2 == 0 && a != 0 {
		fmt.Print("task_1\n", "even\n")
	} else if a%2 == 1 && a != 0 {
		fmt.Print("task_1\n", "odd\n")
	} else {
		fmt.Println("0?-_- taska2")
	}
	WhatIsNum(a)
}

func WhatIsNum(a int) {
	fmt.Println("task_2")
	if a > 0 {
		fmt.Print("\nPositive")
	} else if a < 0 {
		fmt.Print("\nNegative")
	} else {
		fmt.Print("\nZero")
	}
}

func Foreach() {
	fmt.Print("\ntask_3\n")
	for i := 1; i <= 10; i++ {
		fmt.Print(i, " ")
	}
	var str string
	fmt.Print("\nEnter text: \n")
	fmt.Scan(&str)
	StringLen(str)
}

func StringLen(str string) {
	fmt.Print("task_4\n", "Lenght ", "'", str, "'", " is - ", len(str))
	forRect()
}

type Rectangle struct {
	a float64
	b float64
}

func forRect() {
	fmt.Print("\n\ntask_5\n", "enter 2 side of rectangle: \n")

	var a1 float64
	var a2 float64

	fmt.Scan(&a1, &a2)
	rectangle := Rectangle{a: a1, b: a2}
	fmt.Print("Rectangle area is - ", rectangle.a*rectangle.b)
	AverageValue(int(a1), int(a2))
}

func AverageValue(a int, b int) {
	multVar := (a + b) / 2
	fmt.Print("\ntask_6\n", "average value: ", multVar)
}
