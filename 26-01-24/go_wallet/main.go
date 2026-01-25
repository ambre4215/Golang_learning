package main

import (
	"fmt"
	"go_wallet/model"
	"go_wallet/report"
	"go_wallet/service"
	"time"
)

func main() {
	service := &service.Ledger{}
	var time = time.Now()
	var Ledger = []*model.Record{
		{ID: 1, Amount: 4.9, Type: model.Income, Category: "工作", Description: "工资", Time: time},
		{ID: 2, Amount: 20.9, Type: model.Expense, Category: "餐饮", Description: "盖浇饭", Time: time},
		{ID: 3, Amount: 100.9, Type: model.Income, Category: "工作", Description: "工资", Time: time},
		{ID: 4, Amount: 30.9, Type: model.Expense, Category: "其他", Description: "理发", Time: time},
		{ID: 5, Amount: 220.9, Type: model.Income, Category: "工作", Description: "工资", Time: time},
		{ID: 6, Amount: 41.9, Type: model.Expense, Category: "娱乐", Description: "游戏充值", Time: time},
	}
	service.AddRecord(Ledger...)
	Formatter := &report.TextFormatter{}
	content, err := Formatter.Format(service.Records)
	if err != nil {
		fmt.Printf("格式化失败[%v]\n", err)
	} else {
		fmt.Println(content)
	}
}
