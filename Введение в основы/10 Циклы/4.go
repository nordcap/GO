package main

import "fmt"

func main() {
	var n int
	var max = 0
	var count = 0

	for {
		fmt.Scan(&n)
		if n == 0 {
			break
		}
		if n > max {
			max = n
			count = 1
		} else if n == max {
			count++
		}

	}
	fmt.Println(count)

}
