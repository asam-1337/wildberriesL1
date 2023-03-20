package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var n time.Duration
	wg := &sync.WaitGroup{}
	fmt.Scan(&n)
	ch := make(chan int)
	ctx, cancel := context.WithTimeout(context.Background(), n*time.Second)
	defer cancel()

	wg.Add(1)
	go func() {
	loop:
		for {
			select {
			case <-ctx.Done(): // ожидаем истечения времени
				close(ch)
				break loop
			default:
				ch <- rand.Int() // записываем в канал
			}
		}
		wg.Done()
	}()

	for v := range ch { // читаем из канала
		fmt.Println(v)
	}
	wg.Wait()
}
