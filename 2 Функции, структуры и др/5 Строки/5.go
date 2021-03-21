/*Дается строка. Нужно удалить все символы, которые встречаются более одного раза и вывести получившуюся строку*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	fmt.Scan(&str)
	rs := []rune(str)
	for _, elem := range rs {
		if strings.Count(str, string(elem)) > 1 {
			str = strings.ReplaceAll(str, string(elem), "")
		}
	}
	fmt.Println(str)
}
