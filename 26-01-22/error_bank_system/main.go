package main

import (
	"errors"
	"fmt"
)

type Account struct {
	Name    string
	Balance float64
}

func (a *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("存款金额必须大于0\n")
	}
	a.Balance += amount
	fmt.Printf("成功存入%f元,余额%f元\n", amount, a.Balance)
	return nil
}

func (a *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("取款金额不能为0\n")
	}
	if a.Balance < amount {
		return errors.New("余额不足\n")
	}
	a.Balance -= amount
	fmt.Printf("成功取出%f元,余额%f元\n", amount, a.Balance)
	return nil
}

func Transfer(from *Account, to *Account, amount float64) error {
	if amount <= 0 {
		return errors.New("取款金额不能为0\n")
	}
	if from.Balance < amount {
		return errors.New("账户余额不足\n")
	}

	from.Balance -= amount
	to.Balance += amount
	fmt.Printf("成功向%s转账%f元\n", to.Name, amount)
	return nil
}

func (a Account) CheckBalance() {
	fmt.Printf("用户[%s]当前余额为%f\n", a.Name, a.Balance)
}

func main() {
	bankDB := map[string]*Account{
		"1001": {Name: "ambre", Balance: 500},
		"1002": {Name: "david", Balance: 1000},
		"1003": {Name: "snake", Balance: 1500},
		"1004": {Name: "neo", Balance: 2000},
		"1005": {Name: "dante", Balance: 2500},
	}
LOGIN:
	for {
		var inputId string

		fmt.Print("请输入您的用户id\n")
		fmt.Scan(&inputId)

		account, exists := bankDB[inputId]

		if !exists {
			fmt.Printf("用户不存在！\n")
			return
		}
	OP:
		for {
			fmt.Printf("欢迎回来![%s]\n----------------------------\n[1]查询余额\n[2]存款\n[3]取款\n[4]转账\n[5]退出账户\n", account.Name)

			var inputOpration string

			fmt.Scan(&inputOpration)

			if inputOpration == "1" {
				account.CheckBalance()
			}
			var inputDeposit float64
			if inputOpration == "2" {
				fmt.Printf("请输入存款金额\n")

				fmt.Scan(&inputDeposit)

				err := account.Deposit(inputDeposit)

				if err != nil {
					fmt.Println("交易失败\n", err)
				} else {
					account.CheckBalance()
				}
			}

			var inputWithdraw float64

			if inputOpration == "3" {
				fmt.Printf("请输入取款金额\n")

				fmt.Scan(&inputWithdraw)

				err := account.Withdraw(inputWithdraw)

				if err != nil {
					fmt.Println("交易失败\n", err)
				} else {
					account.CheckBalance()
				}
			}

			if inputOpration == "4" {
				var (
					toId        string
					inputAmount float64
				)

				fmt.Printf("请输入转账账户")

				fmt.Scan(&toId)

				toAccount, exists := bankDB[toId]

				if !exists {
					fmt.Printf("用户不存在！\n")
					continue OP
				}

				fmt.Printf("转账金额\n")

				fmt.Scan(&inputAmount)

				err := Transfer(account, toAccount, inputAmount)

				if err != nil {
					fmt.Println("交易失败\n", err)
				} else {
					account.CheckBalance()
				}
			}
			if inputOpration == "5" {
				continue LOGIN
			}
		}
	}
}
