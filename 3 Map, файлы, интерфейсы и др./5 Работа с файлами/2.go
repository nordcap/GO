/*Данная задача в основном ориентирована на изучение типа bufio.Reader, поскольку этот тип позволяет считывать данные постепенно.
В тестовом файле, который вы можете скачать из нашего репозитория на github.com, содержится длинный ряд чисел, разделенных символом ";". Требуется найти, на какой позиции находится число 0 и указать её в качестве ответа. Требуется вывести именно позицию числа, а не индекс (то-есть порядковый номер, нумерация с 1).*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./task.data")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	text, _ := bufio.NewReaderSize(file, 999999).ReadString('\n')
	text = strings.Trim(text, ";") //убираем последний ;
	arrData := strings.Split(text, ";")
	for i, valueStr := range arrData {
		valueInt, err := strconv.ParseInt(valueStr, 10, 64)
		if err != nil {
			panic(err)
		}
		if valueInt == 0 {
			fmt.Println(i + 1)
		}
	}

}
