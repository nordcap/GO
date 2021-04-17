/*Напишите функцию которая принимает канал и число N типа int. Необходимо вернуть значение N+1 в канал.
Функция должна называться task().
*/
package main

import (
	"fmt"
)

func main() {
	channal := make(chan int)
	go task(100, channal)
	fmt.Println(<-channal)
	//time.Sleep(time.Millisecond * 1000)
}

func task(N int, c chan int) {
	c <- N + 1 //запись в канал
}
