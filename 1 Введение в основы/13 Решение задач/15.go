/*Из натурального числа удалить заданную цифру.
Входные данные
Вводятся натуральное число и цифра, которую нужно удалить.
Выходные данные
Вывести число без заданных цифр.*/
package main

import "fmt"

func main() {
	var a, n int
	fmt.Scan(&a, &n)

	arr := make([]int, 0, 0)
	for a != 0 {
		m := a % 10
		if m != n {
			arr = append(arr, m)
		}
		a = a / 10
	}

	for i := len(arr) - 1; i >= 0; i-- {
		fmt.Print(arr[i])
	}

}
