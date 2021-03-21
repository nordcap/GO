/*
Поменяйте местами значения переменных на которые ссылаются указатели. После этого переменные нужно вывести.
ВАЖНО: Считайте что пакет main уже объявлен, а также функция main() уже есть.
 */
package main

import "fmt"

func main()  {
	var a, b int
	fmt.Scan(&a, &b)
	x1 := &a
	x2 := &b
	test2(x1, x2)

}

func test2(x1 *int, x2 *int) {
	*x1, *x2 = *x2, *x1
	fmt.Println(*x1,*x2)
}

