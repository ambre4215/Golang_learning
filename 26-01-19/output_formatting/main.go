package main

import "fmt"

func main() {
	var name string = "John"
	var age int = 25
	var url string = "名字：%s 年龄：%d"
	fmt.Printf(url, name, age)
}
