package main

import "fmt"

func main() {
	var a, solve int
	fmt.Scan(&a)

	solve = (a % 100) / 10
	fmt.Println(solve)

}
