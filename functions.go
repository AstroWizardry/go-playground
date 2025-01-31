package main

import "fmt"

func add(a int, b int) int {
    return a + b
}

func addMany(a int, b int, c int) int {
	return a + b + c
}

func minus(a int, b int) int {	
	return a - b
}

func multiply(a int, b int) int {
	return a * b
}

func divide(a int, b int) int {
	return a / b
}	



func main() {
	fmt.Println(add(1, 2))
	fmt.Println(addMany(1, 2, 3))
	fmt.Println(minus(1, 2))
	fmt.Println(multiply(1, 2))
	fmt.Println(divide(1, 2))
}
