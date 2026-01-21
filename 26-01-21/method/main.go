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

func main() {
	users := []User{
		{Name: "张", Age: 20, Language: "zh"},
		{Name: "吕", Age: 39, Language: "jp"},
		{Name: "李", Age: 28, Language: "zh"},
	}

	for _, user := range users {
		if user.IsSenior() && user.IsJpSpeaker() {
			fmt.Println(user.Name)
		}
	}
}
