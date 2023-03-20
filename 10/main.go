package main

import "fmt"

func main() {
	sl := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	m := make(map[int][]float32) // инициализируем map

	for _, v := range sl { // итерируемся по слайсу температур
		key := int(v) / 10 * 10    // получаем ключ округлением числа до десятых
		m[key] = append(m[key], v) // по ключу собираем группу пополняя слайс темератур
	} // паники не возникает поскольку массив инициализирован

	fmt.Println(m)
}
