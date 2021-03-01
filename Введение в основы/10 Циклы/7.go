package main

import "fmt"

func main() {
	var x, p, y int
	fmt.Scan(&x, &p, &y)
	//x - вклад
	//p - проценты
	//y- размер вклада через несколько лет
	var i = 1
	for ; ; i++ {
		x = (x + (x*p)/100) / 1
		if x >= y {
			break
		}
	}
	fmt.Println(i)
}
