/*Ваша задача сделать проверку подходит ли пароль вводимый пользователем под заданные требования. Длина пароля должна быть не менее 5 символов, он должен содержать только цифры и буквы латинского алфавита. На вход подается строка-пароль. Если пароль соответствует требованиям - вывести "Ok", иначе вывести "Wrong password"*/
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var pass string
	var isLatin = true
	fmt.Scan(&pass)

	rs := []rune(pass)
	for i := range rs {
		if (rs[i] >= 'A' && rs[i] <= 'z') || (rs[i] >= '0' && rs[i] <= '9') {
			continue
		} else {
			isLatin = false
		}
	}
	if utf8.RuneCountInString(pass) < 5 || isLatin == false {
		fmt.Println("Wrong password")
	} else {
		fmt.Println("Ok")
	}

}
