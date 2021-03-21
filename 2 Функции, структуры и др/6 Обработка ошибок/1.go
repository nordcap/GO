/*Вы должны вызвать функцию divide которая делит два числа, но она возвращает не только результат деления, но и информацию об ошибке. В случае какой-либо ошибки вы должны вывести "ошибка", если ошибки нет - результат функции. Функция divide(a int, b int) (int, error) получает на вход два числа которые нужно поделить и возвращает результат (int) и ошибку (error). Пакет main уже объявлен, а нужные пакеты уже импортированы!
Не забудьте считать два числа которые необходимо поделить (передать в функцию)!*/
package main

import (
	"errors"
	"fmt"
)

func main() {
	var a, b int
	err := errors.New("ошибка")
	_, e := fmt.Scan(&a, &b)
	if e != nil {
		fmt.Println(err)
	} else {
		solve, e := divide(a, b)
		if e != nil {
			fmt.Println(err)
		} else {
			fmt.Println(solve)
		}
	}
}

func divide(a int, b int) (int, error) {
	return a / b, errors.New("ошибка")

}
