package service

import (
	"errors"
	"go_wallet/model"
)

type Ledger struct {
	Records []*model.Record
}

func (l *Ledger) AddRecord(records ...*model.Record) error {
	if len(records) == 0 {
		return errors.New("记录不能为空")
	}
	l.Records = append(l.Records, records...)
	return nil
}

func (l Ledger) GetTotalBalance() (float64, error) {
	if len(l.Records) == 0 {
		return 0, errors.New("没有记账记录")
	}
	var totalBalance float64
	for _, r := range l.Records {
		if r.Type == model.Income {
			totalBalance += r.Amount
		} else {
			totalBalance -= r.Amount
		}
	}
	return totalBalance, nil
}
