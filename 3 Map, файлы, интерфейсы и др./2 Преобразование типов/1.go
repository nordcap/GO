/*Напишите функцию с именем convert, которая конвертирует входное число типа int64 в uint16.*/
package main

import (
	"fmt"
	"math"
)

func main()  {
	fmt.Println(math.MaxInt8)
	fmt.Println(math.MaxInt16)
	fmt.Println(math.MaxInt32)
	fmt.Println(math.MaxInt64)

	fmt.Print(convert(64654652))
}

func convert(a int64) uint16 {
	return uint16(a)
}