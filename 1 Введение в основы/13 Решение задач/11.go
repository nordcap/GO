/*
По данному числу n закончите фразу "На лугу пасется..." одним из возможных продолжений: "n коров", "n корова", "n коровы", правильно склоняя слово "корова".
Входные данные
Дано число n (0<n<100).
Выходные данные
Программа должна вывести введенное число n и одно из слов (на латинице): korov, korova или korovy, например, 1 korova, 2 korovy, 5 korov. Между числом и словом должен стоять ровно один пробел.
*/
package main

import "fmt"

func main() {
	var n int
	var cow string
	fmt.Scan(&n)

	remind := n % 10
	if n >= 11 && n <= 14 {
		cow = "korov"
	} else if remind == 1 {
		cow = "korova"
	} else if remind >= 2 && remind <= 4 {
		cow = "korovy"
	} else {
		cow = "korov"
	}
	fmt.Print(n, " ", cow)

}
