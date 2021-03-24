/*Представьте что вы работаете в большой компании где используется модульная архитектура. Ваш коллега написал модуль с какой-то логикой (вы не знаете) и передает в вашу программу какие-то данные. Вы же пишете функцию которая считывает две переменных типа string ,  а возвращает число типа int64 которое нужно получить сложением этих строк.
Но не было бы так все просто, ведь ваш коллега не пишет на Go, и он зол из-за того что гоферам платят больше. Поэтому он решил подшутить над вами и подсунул вам подвох. Он придумал вставлять мусор в строки перед тем как вызывать вашу функцию.
Поэтому предварительно вам нужно убрать из них мусор и конвертировать в числа. Под мусором имеются ввиду лишние символы и спец знаки. Разрешается использовать только определенные пакеты: fmt, strconv, unicode, strings,  bytes. */
package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	a := "%^80"
	b := "hhhhh20&&&&nd"
	c := adding(a, b)
	fmt.Println(c)
}

func adding(a string, b string) int64 {
	rs := []rune(a)
	for _, value := range rs {
		if unicode.IsDigit(value) {
			continue //если цифра идем дальше
		} else {
			//иначе, заменяем на пустой
			a = strings.ReplaceAll(a, string(value), "")
		}
	}
	rs = []rune(b)
	for _, value := range rs {
		if unicode.IsDigit(value) {
			continue //если цифра идем дальше
		} else {
			//иначе, заменяем на пустой
			b = strings.ReplaceAll(b, string(value), "")
		}
	}
	res1, err := strconv.ParseInt(a, 10, 64)
	if err != nil {
		panic(err)
	}
	res2, err := strconv.ParseInt(b, 10, 64)
	if err != nil {
		panic(err)
	}

	return res1 + res2
}
