package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var sum int64
	wg := sync.WaitGroup{}
	sl := []int64{2, 4, 6, 8, 10}
	for _, i := range sl {
		go func(i int64) {
			atomic.AddInt64(&sum, i*i) // атомарно увеличиваем сумму квадратов
			wg.Done()
		}(i)
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(sum)
}
