package bank

import (
	"errors"
	"fmt"
)

type Account struct {
	Name    string
	Balance float64
}

func (a Account) CheckBalance() {
	fmt.Printf("[%s]账户余额为%.2f\n", a.Name, a.Balance)
}

func (a *Account) Deposit(amount float64) error {
	switch {
	case amount <= 0:
		return errors.New("存款金额不能为0或负数\n")
	case amount >= 0:
		a.Balance += amount
		fmt.Printf("存款成功,存入%.2f元\n", amount)
		return nil
	default:
		return errors.New("未知错误\n")
	}
}

func (a *Account) Withdraw(amount float64) error {
	switch {
	case amount <= 0:
		return errors.New("取款金额不能为0或负数\n")
	case a.Balance < amount:
		return errors.New("余额不足\n")
	case a.Balance >= amount:
		a.Balance -= amount
		fmt.Printf("取款成功,取出%.2f元\n", amount)
		return nil
	default:
		return errors.New("未知错误\n")
	}
}

func Transfer(from *Account, to *Account, amount float64) error {
	switch {
	case amount <= 0:
		return errors.New("转账金额不能为0或负数\n")
	case amount > from.Balance:
		return errors.New("余额不足\n")
	case amount <= from.Balance:
		from.Balance -= amount
		to.Balance += amount
		fmt.Printf("转账成功,转账%.2f元\n", amount)
		return nil
	default:
		return errors.New("未知错误\n")
	}
}
