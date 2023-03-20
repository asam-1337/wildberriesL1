package main

import "fmt"

func changeBit(value int64, index int, zero bool) int64 {
	if zero { // меняем на 0
		return value &^ int64(1<<index) // двигаем бит в маске, инвертируем маску и с помощью операции логического 'и' меняем бит в переменной
	} else { // меняем на 1
		return value | int64(1<<index) // двигаем бит в маске и с помощью операции логического 'или' меняем бит в переменной
	}
}

func main() {
	var value int64 = 234885384823423
	fmt.Println(changeBit(value, 10, true))
	fmt.Println(changeBit(value, 5, false))
	fmt.Println(changeBit(value, 4, true))
	fmt.Println(changeBit(value, 56, false))
	fmt.Println(changeBit(value, 63, false))
}
