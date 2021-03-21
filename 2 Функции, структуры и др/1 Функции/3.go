/*Напишите "функцию голосования", возвращающую то значение (0 или 1), которое среди значений ее аргументов x, y, z встречается чаще.
Входные данные
Вводится 3 числа - x, y и z (x, y и z равны 0 или 1).
Выходные данные
Необходимо вернуть значение функции от x, y и z.*/
package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	var solve = vote(a, b, c)
	fmt.Println(solve)
}

func vote(x int, y int, z int) int {
	var out int
	if x+y+z > 1 {
		out = 1
	} else {
		out = 0
	}
	return out
}
