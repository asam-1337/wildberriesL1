package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	sl := []int{2, 4, 6, 8, 10}
	newSl := make([]int, len(sl))
	for i, v := range sl {
		go func(i int, v int) { // запускаем на каждый элемент горутину которая, выведет значение конкурентно
			fmt.Println(v * v)
			newSl[i] = v * v // можно без синхронизации поскольку каждая горутина записывает только по одному индексу
			wg.Done()        // сообщаем что горутина выполнена

		}(i, v)
		wg.Add(1) // говорим о старте новой горутины инкрементируя счетчик
	}
	wg.Wait() // ждем завершения работы горутин
	fmt.Println(newSl)
}
