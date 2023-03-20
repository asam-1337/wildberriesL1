package main

import (
	"fmt"
	"sort"
)

func binarySearch(sl []int, el int) int {
	left, right := 0, len(sl)-1 // начальные значения границ поиска

	for left <= right {
		mid := (left + right) / 2 // находим середину
		if sl[mid] > el {         // если средний элемент больше искомого, ищем в левой части от среднего
			right = mid - 1
		} else if sl[mid] < el { // если меньше ищем в правой части от среднего элемента
			left = mid + 1
		} else { // если равен возвращаем индекс искомого элемента
			return mid
		}
	}
	return -1 // возвращаем -1 если не нашли
}

func main() {
	sl := []int{3, 4, 234, 43, 49, 23, 2, 5, 3432, 34, 1, 7775, 45, 90}
	sort.Slice(sl, func(i, j int) bool { // сортируем слайс перед тем как использовать бинарный поиск
		return sl[i] < sl[j]
	})
	fmt.Println(binarySearch(sl, 5))
	fmt.Println(binarySearch(sl, 1111))
}
