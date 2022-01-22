package main

func Add(a int, b int) int {
	return a + b
}

func Sub(a int, b int) int {
	return a - b
}

func Mul(a int, b int) int {
	return a * b
}

func Div(a int, b int) int {

	var res int

	if b > 0 {
		res = a / b
	} else {
		res = -1
	}

	return res
}
