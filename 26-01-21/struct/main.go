package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	users := []User{
		{Name: "zhang", Age: 20},
		{Name: "lv", Age: 39},
		{Name: "li", Age: 28},
	}

	for _, user := range users {
		if user.Age >= 30 {
			fmt.Println(user.Name)
		}
	}
}
