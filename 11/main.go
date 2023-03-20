package main

import "fmt"

func Intersection(sl1, sl2 []int) []int {
	m := make(map[int]struct{})

	for _, v := range sl1 { // заносим в мапу все элементы из первого множества
		m[v] = struct{}{}
	}

	res := make([]int, 0)
	for _, v := range sl2 { // итерируемся по 2 множеству и проверяем есть ли такие элементы в 1 множестве
		_, ok := m[v]
		if ok {
			res = append(res, v)
		}
	}

	return res
}

func main() {
	sl1 := []int{1, 12, 343, 456, 42, 7, 78, 9, 4}
	sl2 := []int{1, 143, 343, 45642, 7, 78, 9, 5}

	fmt.Println(Intersection(sl1, sl2))
	fmt.Println(Intersection(sl1, sl1))
}
