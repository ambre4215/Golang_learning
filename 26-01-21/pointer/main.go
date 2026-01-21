package main

import "fmt"

type User struct {
	Name     string
	Age      int
	Language string
}

func (u User) IsSenior() bool {
	return u.Age >= 30
}

func (u User) IsJpSpeaker() bool {
	return u.Language == "jp"
}

func (u *User) HappyBirthday() {
	u.Age = u.Age + 1
}

func main() {
	users := []User{
		{Name: "张", Age: 20, Language: "zh"},
		{Name: "吕", Age: 39, Language: "jp"},
		{Name: "李", Age: 28, Language: "zh"},
		{Name: "赵", Age: 29, Language: "zh"},
	}

	for i := range users {
		if users[i].Age == 29 {
			fmt.Printf("用户之前%d岁\n", users[i].Age)
			users[i].HappyBirthday()
			fmt.Printf("过完生日了，现在%d岁", users[i].Age)
		}
	}
}
