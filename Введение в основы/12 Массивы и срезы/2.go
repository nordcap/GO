/*
Напишите программу, принимающая на вход число N (N \geq 4)N(N≥4), а затем NN целых чисел-элементов среза.
На вывод нужно подать 4-ый (3 по индексу) элемент данного среза.
*/
package main

import "fmt"

func main() {
	var a, N int
	fmt.Scan(&N)

	slice := make([]int, 0, 0)

	for i := 0; i < N; i++ {
		fmt.Scan(&a)
		slice = append(slice, a)
	}

	fmt.Println(slice[3])

}
