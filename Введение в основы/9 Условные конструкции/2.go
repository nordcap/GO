package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)

	var digit1 int = a / 100
	var digit2 int = (a % 100) / 10
	var digit3 int = a % 10

	if (digit1-digit2 == 0) || (digit1-digit3 == 0) || (digit2-digit3 == 0) {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
	}

}
