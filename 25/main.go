package main

import (
	"fmt"
	"time"
)

func Sleep1(d time.Duration) {
	<-time.After(d) // горутина блокирует свое выполнение пока в канал не придет значение
}

func sleep2(d time.Duration) {
	end := time.Now().Add(d).UnixNano()

	for { // находимся в цикле, пока число наносекунд прошедшее с 01.01.1970 конечного времени
		if end <= time.Now().UnixNano() { // больше чем в данный момент
			return
		}
	}
}

func main() {
	fmt.Println("start")
	Sleep1(5 * time.Second)
	fmt.Println("sleep 5s")
	sleep2(5 * time.Second)
	fmt.Println("sleep 5s")
}
