/*Напишите функцию sumInt, принимающую переменное количество аргументов типа int, и возвращающую количество полученных функцией аргументов и их сумму. Пакет "fmt" уже импортирован, функция и пакет main объявлены.
Пример вызова вашей функции:
a, b := sumInt(1, 0)
fmt.Println(a, b)*/
package main

import "fmt"

func main() {
	a, b := sumInt(1, 6, 9)
	fmt.Println(a, b)

}

func sumInt(a ...int) (int, int) {

	var sum = 0
	for _, elem := range a {
		sum += elem
	}
	return len(a), sum
}
