package main

import "fmt"

func main() {
	var d int
	fmt.Scan(&d)

	var deg int = 2 //на 1 градус приходится 2минуты
	var minAll int = d * deg
	var hour int = minAll / 60

	fmt.Println("It is", hour, "hours", minAll-hour*60, "minutes.")

}
