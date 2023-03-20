package main

import (
	"fmt"
)

func getType(val interface{}) string {
	switch val.(type) { // получаем тип встроенным средством go
	case bool: // свитчимся в зависимости от того какой тип
		return fmt.Sprintf("%T", val) // если нужно получить тип в виде строки
	case int:
		return fmt.Sprintf("%T", val)
	default:
		return fmt.Sprintf("%T", val)
	}
}

func main() {
	var a bool
	var b int
	var c chan string
	fmt.Println(getType(a), getType(b), getType(c))
}
