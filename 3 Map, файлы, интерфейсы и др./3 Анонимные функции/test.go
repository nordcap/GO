package main

import "fmt"

func main()  {
	ExampleClosure()
}

func ExampleClosure() {
	fn := func() func(int) int {
		count := 0
		return func(i int) int {
			count++
			return count * i
		}
	}()

	for i := 1; i <= 5; i++ {
		fmt.Println(fn(i))
	}

	// Output:
	// 1
	// 4
	// 9
	// 16
	// 25
}
