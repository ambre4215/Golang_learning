package main

import "fmt"

type USBDevice interface {
	connect()
}

type Iphone struct {
	model string
}

func (i *Iphone) connect() {
	fmt.Println(i.model, "已连接，开始充电...")
}

type Ipad struct {
	size int
}

func (i *Ipad) connect() {
	fmt.Println(i.size, "寸Ipad已连接，开始充电")
}

type Mp3 struct {
	model string
}

func (m *Mp3) connect() {
	fmt.Println(m.model, "已连接，开始充电...")
}

func main() {
	device := []USBDevice{
		&Iphone{"Iphone 13 Pro"},
		&Ipad{17},
		&Mp3{"索尼"},
	}
	for _, v := range device {
		v.connect()
	}
}
