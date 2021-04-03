/*Задача состоит в следующем: на стандартный ввод подаются целые числа в диапазоне 0-100, каждое число подается на стандартный ввод с новой строки (количество чисел не известно). Требуется прочитать все эти числа и вывести в стандартный вывод их сумму.*/
package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

func main() {
	var count int = 0
	var sum int = 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if value, err := strconv.Atoi(scanner.Text()); err == nil {
			sum = sum + value
		}

		count++
		if count == 5 {
			break
		}
	}
	if _, err := io.WriteString(os.Stdout, strconv.Itoa(sum)); err != nil {
		panic(err)
	}
}
