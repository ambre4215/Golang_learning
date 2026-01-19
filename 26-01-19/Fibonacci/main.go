package main

import "fmt"

func main() {
	a, b := 0, 1
	c := a + b
	for i := 0; i < 20; i++ {
		fmt.Printf("第%d次运算 a=%d b=%d, c=%d\n", i+1, a, b, c)
		a, b = b, c
		c = a + b
	}
}
