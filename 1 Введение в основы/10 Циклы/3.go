package main

import "fmt"

func main() {
	var count, n int
	fmt.Scan(&count)

	sum := 0
	for i := 0; i < count; i++ {
		//анализируем n
		fmt.Scan(&n)
		if (n%100 == n) && (n > 9) && (n%8 == 0) {
			sum = sum + n
		}
	}
	fmt.Println(sum)
}
