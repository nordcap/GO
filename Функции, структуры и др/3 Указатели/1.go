/*Напишите функцию, которая умножает значения на которые ссылаются два указателя и выводит получившееся произведение в консоль. Ввод данных уже написан за вас.*/
package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	x1 := &a
	x2 := &b
	test(x1, x2)

}

func test(x1 *int, x2 *int) {
	out := *x1 * *x2
	fmt.Println(out)
}
