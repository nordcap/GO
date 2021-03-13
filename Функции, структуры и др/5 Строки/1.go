/*На вход подается строка. Нужно определить, является ли она правильной или нет. Правильная строка начинается с заглавной буквы и заканчивается точкой. Если строка правильная - вывести Right иначе - вывести Wrong*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	text = strings.Trim(text, "\n")
	var rs = []rune(text)
	var last = len(rs) - 1

	if unicode.IsUpper(rs[0]) == true && rs[last] == '.' {
		fmt.Println("Right")
	} else {
		fmt.Println("Wrong")
	}

}
