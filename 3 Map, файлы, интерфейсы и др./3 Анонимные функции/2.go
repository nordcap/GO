/*Используем анонимные функции на практике.
Внутри функции main (объявлять ее не нужно) вы должны объявить функцию вида func(uint) uint, которую внутри функции main в дальнейшем можно будет вызвать по имени fn.
Функция на вход получает целое положительное число (uint), возвращает число того же типа в котором отсутствуют нечетные цифры или цифра 0, если же получившееся число равно 0, то возвращается число 100. Цифры в возвращаемом числе имеют тот же порядок, что и в исходном числе.*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	fn := func(i uint) uint {
		if i == 0 {
			return 100
		} else {
			str := strconv.FormatUint(uint64(i), 10)

			rs := []rune(str)
			rsout := []rune("")
			for _, value := range rs {
				if value == '0' ||
					value == '1' ||
					value == '3' ||
					value == '5' ||
					value == '7' ||
					value == '9' {
					continue
				} else {
					rsout = append(rsout, value)
				}
			}

			out, err := strconv.ParseUint(string(rsout), 10, 64)
			if err != nil {
				return 100
			}
			return uint(out)
		}
	}
	fmt.Println(fn(727178))
	fmt.Println(fn(777777))
	fmt.Println(fn(0))
}
