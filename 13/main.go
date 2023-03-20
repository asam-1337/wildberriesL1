package main

import "fmt"

func main() {
	// меняет два числа доступными функциями языка
	a, b := 1, 2
	a, b = b, a
	fmt.Println("a:", a)
	fmt.Println("b:", b)

	// меняет два числа математическими операциями
	c, d := 3, 4
	c += d
	d = c - d
	c = c - d

	fmt.Println("c:", c)
	fmt.Println("d:", d)
}
