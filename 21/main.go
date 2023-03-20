package main

import "fmt"

// Adaptee - обьект, реализация которого нас устраивает и мы хотим его использовать,
// однако он не подходит под целевой интерфейс Target
type Adaptee struct {
}

// Метод который не имплементирует нужный интерфейс
func (a *Adaptee) SpecificRequest() string {
	return "method SpecificRequest()"
}

// Adapter имлеметирует интерфейс Target
type Adapter struct {
	*Adaptee
}

// Возвращает обьект удовлетворяющий интерфейсу Target
func NewAdapter(adaptee *Adaptee) Target {
	return &Adapter{
		Adaptee: adaptee,
	}
}

// Адаптивный метод, который уже удовлетворяет нашему интерфейсу
func (a *Adapter) Request() string {
	// doing something...
	return "method Request() call: " + a.SpecificRequest()
}

type Target interface {
	Request() string
}

func main() {
	adaptee := &Adaptee{}
	adapter := NewAdapter(adaptee)

	fmt.Println(adapter.Request())
}
