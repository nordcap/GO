/*
Самое большое число, кратное 7
Найдите самое большее число на отрезке от a до b, кратное 7 .
Входные данные
Вводится два целых числа a и b (a≤b).
Выходные данные
Найдите самое большее число на отрезке от a до b (отрезок включает в себя числа a и b), кратное 7 , или выведите "NO" - если таковых нет.
*/
package main

import "fmt"

func main() {
	var a, b, solve, count int
	fmt.Scan(&a, &b)
	for i := a; i <= b; i++ {
		if i%7 == 0 {
			solve = i
			count++
		}
	}
	if count > 0 {
		fmt.Println(solve)
	} else {
		fmt.Println("NO")
	}

}
