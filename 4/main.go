package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func startWorker(ch <-chan int, wg *sync.WaitGroup) {
	for v := range ch {
		fmt.Println(v)
		time.Sleep(200 * time.Millisecond) // sleep для наглядности
	}
	wg.Done()
}

func main() {
	var n int                           // число воркеров
	sigs := make(chan os.Signal, 1)     // канал для системных сигналов
	signal.Notify(sigs, syscall.SIGINT) // signal: interrupt, говорим, если будет сигнал, сообщи об этом в канал sigs

	fmt.Scan(&n)

	ch := make(chan int)    // канал данных
	wg := &sync.WaitGroup{} // для синхронизации и ожидания, чтобы не было утечки

	for i := 0; i < n; i++ {
		wg.Add(1)
		go startWorker(ch, wg) // все воркеры читают из одного и того же канала
	}

LOOP:
	for {
		rnd := rand.Int()
		select {
		case <-sigs: // ожидает прихода системного сигнала о завершении
			fmt.Println("close")
			close(ch) // закрывает канал и воркеры завершают свое выполнение
			break LOOP
		default:
			ch <- rnd
		}
	}
	wg.Wait()
}
