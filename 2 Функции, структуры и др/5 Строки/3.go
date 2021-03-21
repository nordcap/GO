/*Даются две строки X и S. Нужно найти и вывести первое вхождение подстроки S в строке X. Если подстроки S нет в строке X - вывести -1*/
package main

import (
	"fmt"
	"strings"
)

func main()  {
	var x,s string
	var idx int
	fmt.Scan(&x)
	fmt.Scan(&s)
	idx = strings.Index(x,s)
	fmt.Println(idx)
}

