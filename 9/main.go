package main

import "fmt"

func main() {
	sl := []int{1, 2, 3, 4, 56, 123, 123, 9437, 5466, 6548, 97}
	in := make(chan int)
	out := make(chan int)

	go func() {
		for _, v := range sl { // итерируемся по массиву и записываем в канал данные
			in <- v
		}
		close(in) // закрываем канал по завершению
	}()

	go func() {
		for v := range in { // читаем данные из канала
			out <- v * v // записываем в другой канал данные возведенные в квадрат
		}
		close(out) // закрываем канал по завершению
	}()

	for v := range out { // цикл не закончится пока канал не закроется поэтому gorutines leak быть не должно
		fmt.Println(v)
	}
}
