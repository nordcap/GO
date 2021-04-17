/*Напишите функцию которая принимает канал и строку. Необходимо отправить эту же строку в заданный канал 5 раз, добавив к ней пробел.
Функция должна называться task2().*/
package main

import "fmt"

func main() {
	c := make(chan string)
	go task2("hello", c)
	for v:=range c{
		fmt.Print(v)
	}

}

func task2(str string, c chan string) {
	for i := 0; i < 5; i++ {
		c <- str + " "
	}
	close(c)
}
