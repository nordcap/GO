package main

import "fmt"

func main() {
	var (
		a         uint8
		workArray [10]uint8
		workIndex [6]uint8
	)

	for idx := range workArray {
		fmt.Scan(&a)
		workArray[idx] = a
	}

	for i := 0; i < 6; i++ {
		fmt.Scan(&a)
		workIndex[i] = a
	}

	for i := 0; i < 6; i+=2 {
		workArray[workIndex[i]], workArray[workIndex[i+1]] = workArray[workIndex[i+1]], workArray[workIndex[i]]
	}

	for _, elem := range workArray {
		fmt.Print(elem, " ")
	}

}
