package main

import (
	"fmt"
	"strings"
)

func CheckUniqueSymbols(s string) bool {
	s = strings.ToLower(s)       // переводим строку в нижний регистр
	m := make(map[rune]struct{}) // создается map для удобного отслеживания повторяющихся символов

	for _, v := range s {
		if _, ok := m[v]; ok { // проверяем попадлся ли уже такой символ
			return false // если попадался то выходим
		}

		m[v] = struct{}{} // сигнализируем что символ встретился
	}

	return true
}

func checkUniqueSymbols(s string) bool {
	return false
}

func main() {
	s1 := "KMKLMFmkmlmkmFEMKs"
	s2 := "abcd"
	fmt.Println(CheckUniqueSymbols(s1))
	fmt.Println(CheckUniqueSymbols(s2))
}
