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
}
