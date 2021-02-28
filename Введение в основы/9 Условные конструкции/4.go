package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	t6 := a % 10
	a = a / 10

	t5 := a % 10
	a = a / 10

	t4 := a % 10
	a = a / 10

	t3 := a % 10
	a = a / 10

	t2 := a % 10
	a = a / 10

	t1 := a % 10

	sum1 := t1 + t2 + t3
	sum2 := t4 + t5 + t6

	if sum1 == sum2 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}
