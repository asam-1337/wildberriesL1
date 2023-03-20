package main

import (
	"fmt"
	"strings"
)

func revertWords(s string) string {
	if len(s) <= 1 {
		return s
	}

	words := strings.Split(s, " ")
	res := "" // создаем новую строку

	for i := len(words) - 1; i > 0; i-- {
		res += words[i] + " " // конкатенируем новую строку со словами взятыми в обратном порядке
	}
	res += words[0] // добавляем последнее слово в новой строке отдельно чтобы не ставился пробел после

	return res
}

func main() {
	s := "snow dog sun"
	fmt.Println(revertWords(s))
}
