package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func (u *User) MarkJunior() {
	u.Name = "[JUNIOR]" + u.Name
}

func main() {
	users := []User{
		{Name: "张", Age: 15},
		{Name: "吕", Age: 16},
		{Name: "赵", Age: 28},
		{Name: "钱", Age: 29},
		{Name: "孙", Age: 32},
		{Name: "李", Age: 55},
		{Name: "周", Age: 12},
		{Name: "吴", Age: 17},
		{Name: "郑", Age: 77},
	}
	count := 0
	for i := range users {
		if users[i].Age < 18 {
			users[i].MarkJunior()
			count++
		}
	}
	fmt.Println(users, count)
}
