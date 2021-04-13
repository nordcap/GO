/*На стандартный ввод подаются данные о продолжительности периода (формат приведен в примере). Кроме того, вам дана дата в формате Unix-Time: 1589570165 в виде константы типа int64 (наносекунды для целей преобразования равны 0).
Требуется считать данные о продолжительности периода, преобразовать их в тип Duration, а затем вывести (в формате UnixDate) дату и время, получившиеся при добавлении периода к стандартной дате.
Небольшая подсказка: базовую дату необходимо явно перенести в зону UTC с помощью одноименного метода.*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

// вам это понадобится
const now int64 = 1589570165
const UnixDate = "Mon Jan _2 15:04:05 MST 2006"

func main() {
	rd, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	strInput := strings.Trim(rd, "\n")
	arr := strings.Split(strInput, " ")
	// преобразует строку в Duration с использованием аннотаций:
	// "ns" - наносекунды,
	// "us" - микросекунды,
	// "ms" - миллисекунды,
	// "s" - секунды,
	// "m" - минуты,
	// "h" - часы.
	strDuration := arr[0] + "m" + arr[2] + "s"
	dur, err := time.ParseDuration(strDuration)
	if err != nil {
		panic(err)
	}

	curTime := time.Unix(now, 0).UTC()
	futureTime := curTime.Add(dur)

	fmt.Println(futureTime.Format(UnixDate))

}
