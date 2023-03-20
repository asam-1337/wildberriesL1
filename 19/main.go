package main

import "fmt"

func revertString(s string) string {
	if len(s) <= 1 { // если длина меньше или равна 1 то ничего не делаем
		return s
	}

	runes := make([]rune, len(s))
	for i, v := range s { // идем в обратном порядке, тем самым переворачивая строку
		runes[len(s)-1-i] = v
	}

	return string(runes) // возвращаем готовую строку
}

func main() {
	s := "kfkakd fdsaj mfdsm f"
	fmt.Println(revertString(s))
}
