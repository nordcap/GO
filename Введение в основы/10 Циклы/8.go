package main

import "fmt"

func main() {
	var a, b string
	fmt.Scan(&a, &b)
	var length_a = len(a)
	var length_b = len(b)
	for i := 0; i < length_a; i++ {
		var count = 0
		for j := 0; j < length_b; j++ {
			if string(a[i]) == string(b[j]) {
				count++
			}
		}
		if count > 0 {
			fmt.Print(string(a[i]), " ")
		}
	}

}
