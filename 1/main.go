package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

func (h Human) Say() string {
	return fmt.Sprintf("I am %s", h.Name)
}

type Action struct {
	Human //Встраивание структуры Human
	Name  string
}

func (a *Action) Say() string {
	return fmt.Sprintf("%s and doing %s", a.Human.Say(), a.Name)
}

func main() {
	human := Human{
		Name: "Vasya",
		Age:  20,
	}

	workAction := Action{
		Human: human,
		Name:  "home work",
	}

	fmt.Println(workAction.Say())       // неявный вызов встраиваемых методов
	fmt.Println(workAction.Human.Say()) // явный вызов
	fmt.Println(workAction.Name)        // при одинаковых имен полей или методов будет вызван метод или получено поле которое ближе
	fmt.Println(workAction.Human.Name)
}
