/*На вход подается строка. Нужно определить, является ли она палиндромом. Если строка является паллиндромом - вывести Палиндром иначе - вывести Нет. Все входные строчки в нижнем регистре.
Палиндром — буквосочетание, слово или текст, одинаково читающееся в обоих направлениях (например, "топот", "око", "заказ").*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	text = strings.Trim(text, "\n")
	var rs = []rune(text)
	var flag = true
	for i := range rs {
		if rs[i] != rs[len(rs)-i-1] {
			flag = false
		}
	}
	if flag == true {
		fmt.Println("Палиндром")
	} else {
		fmt.Println("Нет")
	}

}
