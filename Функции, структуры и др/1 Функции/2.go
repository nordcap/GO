/*Напишите функцию, находящую наименьшее из четырех введённых в этой же функции чисел.
Входные данные
Вводится четыре числа.
Выходные данные
Необходимо вернуть из функции наименьшее из 4-х данных чисел*/
package main

import "fmt"

func main() {
	var m = minimumFromFour()
	fmt.Println(m)
}

func minimumFromFour() int {
	var arr [4]int

	for i := 0; i < 4; i++ {
		fmt.Scan(&arr[i])
	}

	m := arr[0]
	for _, elem := range arr {
		if elem < m {
			m = elem
		}
	}
	return m
}
