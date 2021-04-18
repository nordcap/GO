/*Напишите элемент конвейера (функцию), что запоминает предыдущее значение и отправляет значения на следующий этап конвейера только если оно отличается от того, что пришло ранее.
Ваша функция должна принимать два канала - inputStream и outputStream, в первый вы будете получать строки, во второй вы должны отправлять значения без повторов. В итоге в outputStream должны остаться значения, которые не повторяются подряд. Не забудьте закрыть канал ;)
Функция должна называться removeDuplicates()*/
package main

import (
	"fmt"
)

func main() {

	inputStream := make(chan string)
	outputStream := make(chan string)

	arrString := [...]string{"aaa", "bbb", "bbb", "ccc"}
	go removeDuplicates(inputStream, outputStream)
	for _, r := range arrString {
		inputStream <- r
		fmt.Println("r")
	}
	for val:=range outputStream{
		fmt.Println(val)
	}
	close(inputStream)
}

func removeDuplicates(inputStream chan string, outputStream chan string) {

	var f string = ""
	for val := range inputStream {
		if len(f)==0 {
			f = val
			outputStream <-val
		} else {
			//сравниваем текущее и прошлое значение
			if f==val {
				continue //пропускаем итерацию
			} else{
				f=val
				outputStream <-val
			}
		}
	}
close(outputStream)
}
