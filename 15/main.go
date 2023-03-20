package main

import (
	"fmt"
	"strings"
)

var justString string

func createHugeString(n int) string {
	sb := strings.Builder{}
	for i := 0; i < n; i++ { // создает новую строку
		sb.WriteString("字")
	}
	return sb.String()
}

func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100] // т.к некоторые символы занимают больше 1 байта то можем получить строку длиной меньше 100 символов
} // нужно использовать срез рун

func correctFunc() {
	v := createHugeString(1 << 10)
	justString = string([]rune(v)[:100])
}

func main() {
	correctFunc()
	fmt.Println(justString)
}
