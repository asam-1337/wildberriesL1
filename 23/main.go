package main

import "fmt"

func deleteElem(sl []int, index int) []int {
	if len(sl) < 1 { // выходим если пустой слайс
		return sl
	}
	return append(sl[:index], sl[index+1:]...) // берем часть элементов до целевого элемента и добавляем часть справа от элемента
}

func main() {
	sl := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sl = deleteElem(sl, 4)
	fmt.Println(sl)

}
