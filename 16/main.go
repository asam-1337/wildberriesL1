package main

import (
	"fmt"
	"sort"
)

// 1 способ
func sortSlice(sl []int) {
	sort.Slice(sl, func(l, r int) bool { // под капотом используется быстрая сортировка
		return sl[l] < sl[r] // условие сортировки
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

func partition(sl []int, left, right int) int {
	mid := sl[(left+right)/2] // берем центральный элемент в качестве опорного
	l, r := left, right

	for l <= r {
		// ищем значения меньше центра
		for sl[l] < mid {
			l++
		}
		// ищем значения больше центра
		for sl[r] > mid {
			r--
		}

		if l >= r {
			break
		}

		sl[l], sl[r] = sl[r], sl[l]

		l++
		r--
	}
	return r

}

func main() {
	sl := []int{3, 4, 234, 43, 49, 23, 2}
	sortSlice(sl)
	fmt.Println(sl)
	sl = []int{3, 4, 234, 43, 49, 23, 2, 234, 236, 85395, 879, 6856754, 3242, 4434, 57, 1, 8983457, 6789, 4}
	quickSort(sl, 0, len(sl)-1)
	fmt.Println(sl)
}
