package main

import "fmt"

func main() {
	var a float64
	fmt.Scan(&a)

	if a <= 0 {
		fmt.Printf("число %2.2f не подходит", a)
	} else if a > 10000 {
		fmt.Printf("%e", a)
	} else {
		b := a * a
		fmt.Printf("%.4f", b)
	}
}
