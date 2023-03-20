package main

import (
	"fmt"
	"math/big"
)

func Calculate(a, b *big.Int, operation string) *big.Int {

	result := new(big.Int)

	switch operation { // вызываем функцию в зависимости от операции
	case "+":
		return result.Add(a, b) // суммируем значения

	case "-":
		return result.Sub(a, b) // вычитаем одно значения из другого

	case "*":
		return result.Mul(a, b) // перемножаем значения

	case "/":
		return result.Quo(a, b) // делим одно значение на другое

	default:
		return result // другой оператор
	}
}

func main() {

	a, b := big.NewInt(1<<27), big.NewInt(1<<24) // числа более 2^20
	fmt.Println(Calculate(a, b, "+"))
	fmt.Println(Calculate(a, b, "-"))
	fmt.Println(Calculate(a, b, "*"))
	fmt.Println(Calculate(a, b, "/"))

}
