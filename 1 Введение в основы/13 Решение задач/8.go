/*Количество минимумов
Найдите количество минимальных элементов в последовательности.
Входные данные
Вводится натуральное число N, а затем N чисел.
Выходные данные
Выведите количество минимальных элементов.*/
package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	slice := make([]int, n, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&slice[i])
	}
	min := slice[0]
	count_min := 0
	for _, elem := range slice {
		if elem < min {
			min = elem
			count_min = 1
		} else if elem == min {
			count_min++
		}
	}

	fmt.Println(count_min)
}
