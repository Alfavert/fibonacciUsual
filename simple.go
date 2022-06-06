package main

import "fmt"

func main() {
	x1 := 0
	x2 := 1
	for x := 0; x < 10; x++ {
		if x == 0 {
			fib := 0
			fmt.Println(fib)
		} else if x == 1 {
			fib := 1
			fmt.Println(fib)
		} else {
			fib := x1 + x2
			x1 = x2
			x2 = fib
			fmt.Println(fib)
		}
	}
}
