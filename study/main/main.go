package main

import (
	"fmt"
	"gitTest/study/greet"
)

func main() {

	fmt.Println("Hello From Main Package")
	fmt.Println(greet.Hello("shamim"))

	var check int = Mul(10, 2)
	fmt.Println(check)

	var check2 int = Div(10, 2)
	fmt.Println(check2)

	EvenOdd(7)
}
