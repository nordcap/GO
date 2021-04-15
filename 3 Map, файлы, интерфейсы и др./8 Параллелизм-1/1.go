package main

import "fmt"

func main() {
	c := make(chan int, 5)
	for i := 0; i < 5; i++ {
		go myFunc(c)
	}
	fmt.Println(<-c)
}

//горутина
func myFunc(c chan int) {
	c <- cap(c)
}
