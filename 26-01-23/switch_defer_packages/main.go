package main

import (
	"fmt"
	"packages/bank"
)

func main() {
	bankDB := bank.LoadBankData()
LOGIN:
	for {
		fmt.Printf("请输入账户ID按回车登录")
		var inputId string
		fmt.Scan(&inputId)
		account, exists := bankDB[inputId]
		if !exists {
			fmt.Printf("用户不存在！")
			continue LOGIN
		}
	OP:
		for {
			fmt.Printf("欢迎回来[%s],请选择你需要的操作\n", account.Name)
			fmt.Printf("———————————————————————————————\n")
			fmt.Printf("[1]查询余额\n")
			fmt.Printf("[2]存款\n")
			fmt.Printf("[3]取款\n")
			fmt.Printf("[4]转账\n")
			fmt.Printf("[5]退出登录\n")
			var inputOpration int
			var inputToId string
			var amount float64
			fmt.Scan(&inputOpration)
			switch inputOpration {
			case 1:
				account.CheckBalance()
			case 2:
				fmt.Printf("请输入要存入的金额\n")
				fmt.Scan(&amount)
				err := account.Deposit(amount)
				if err != nil {
					fmt.Printf("交易失败！[%v]\n", err)
				} else {
					account.CheckBalance()
					bank.SaveBankData(bankDB)
				}

			case 3:
				fmt.Printf("请输入要取出的金额\n")
				fmt.Scan(&amount)
				err := account.Withdraw(amount)
				if err != nil {
					fmt.Printf("交易失败！[%v]\n", err)
				} else {
					account.CheckBalance()
					bank.SaveBankData(bankDB)
				}
			case 4:
				fmt.Printf("请输入要转向的账户号\n")
				fmt.Scan(&inputToId)
				toAccount, exists := bankDB[inputToId]
				if !exists {
					fmt.Printf("用户不存在！\n")
					continue OP
				}
				fmt.Printf("请输入要转账的金额\n")
				fmt.Scan(&amount)
				err := bank.Transfer(account, toAccount, amount)
				if err != nil {
					fmt.Printf("交易失败！[%v]\n", err)
				} else {
					account.CheckBalance()
					bank.SaveBankData(bankDB)
				}
			case 5:
				continue LOGIN
			}
		}
	}
}
