package main

import "fmt"

// исключаем повторяющиеся элементы
func main() {
	sl := []string{"cat", "cat", "dog", "cat", "tree"}
	m := make(map[string]struct{}) // по ключу будут элементы хранится, а значение будет пустой структурой, не занимающей память

	for _, v := range sl {
		m[v] = struct{}{}
	}

	fmt.Println(m)
}
