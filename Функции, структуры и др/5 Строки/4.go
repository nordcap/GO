/* На вход дается строка, из нее нужно сделать другую строку, оставив только нечетные символы (считая с нуля)*/
package main

import "fmt"

func main()  {
	var str string
	fmt.Scan(&str)

	var rs = []rune(str)
	for i, elem:=range rs{
		if i%2 == 1 {
			//fmt.Print(string(rs[i]))
			fmt.Print(string(elem))
		}
	}
}

