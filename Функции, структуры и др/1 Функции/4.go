/*Последовательность Фибоначчи определена следующим образом: φ1=1, φ2=1, φn=φn-1+φn-2 при n>1. Начало ряда Фибоначчи выглядит следующим образом: 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, ... Напишите функцию, которая по указанному натуральному n возвращает φn.
Входные данные
Вводится одно число n.
Выходные данные
Необходимо вывести  значение φn.*/
package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var solve = fibonacci(n)
	fmt.Println(solve)

}

func fibonacci(n int) int {

	fib := []int{0, 1}
	if n == 0 || n == 1 {
		return fib[n]
	}
	for i := 2; ; i++ {
		fib = append(fib, fib[i-2]+fib[i-1])
		if n == i {
			return fib[i]
		}
	}
}
