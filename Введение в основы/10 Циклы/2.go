package main

import "fmt"

func main() {
	var A, B int
	fmt.Scan(&A, &B)

	sum := 0
	for i := A; i <= B; i++ {
		sum = sum + i
	}

	fmt.Println(sum)
}
