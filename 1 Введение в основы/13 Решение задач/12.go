/*
По данному числу N распечатайте все целые значения степени двойки, не превосходящие N, в порядке возрастания.

Входные данные
Вводится натуральное число.
Выходные данные
Выведите ответ на задачу.
*/
package main

import "fmt"
import "math"

func main() {
	var n float64
	fmt.Scan(&n)

	s := 0.0
	for i := 0; s <= n; i++ {
		s = math.Pow(2, float64(i))
		if s > n {
			break
		}
		fmt.Print(s, " ")
	}
}
