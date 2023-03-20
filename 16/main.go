package main

import (
	"fmt"
	"sort"
)

// 1 способ
func sortSlice(sl []int) {
	sort.Slice(sl, func(i, j int) bool { // под капотом используется быстрая сортировка
		return sl[i] < sl[j] // условие сортировки
	})
}

// 2 способ
func quickSort(sl []int, left, right int) {
	if left < right {
		p := partition(sl, left, right)
		quickSort(sl, left, p)
		quickSort(sl, p+1, right)
	}
}

// необходимо найти последний элемент слева, больший чем центральный и первый элемент справа, меньший чем центральный и поменять их местами
// повторяем процедуру до тех пор, пока индекс левого элемента не станет равным или большим чем индекс правого элемента
func partition(sl []int, left, right int) int {
	mid := sl[(left+right)/2] // берем центральный элемент, в зависимости от того какой выбрать меняется производительность
	i, j := left, right

	for {
		for sl[i] < mid {
			i++
		}

		for sl[j] > mid {
			j--
		}

		if i >= j {
			break
		}

		sl[i], sl[j] = sl[j], sl[i]

		i++
		j--
	}
	return j // новый опроный элемент

}

func main() {
	sl := []int{3, 4, 234, 43, 49, 23, 2}
	sortSlice(sl)
	fmt.Println(sl)
	sl = []int{3, 4, 234, 43, 49, 23, 2}
	quickSort(sl, 0, len(sl)-1)
	fmt.Println(sl)
}
