package report

import (
	"errors"
	"fmt"
	"go_wallet/model"
)

type Formatter interface {
	Format(records []*model.Record) (string error)
}

type TextFormatter struct{}

func (t *TextFormatter) Format(records []*model.Record) (string, error) {
	if len(records) == 0 {
		return "", errors.New("无可用数据")
	}
	var textOut string
	for _, r := range records {
		if r.Type == model.Income {
			textOut += fmt.Sprintf("%v. [%v] %v (+%v)\n==========================\n", r.ID, r.Category, r.Description, r.Amount)
		} else {
			textOut += fmt.Sprintf("%v. [%v] %v (-%v)\n==========================\n", r.ID, r.Category, r.Description, r.Amount)
		}
	}
	return textOut, nil
}

type CSVFormatter struct{}

func (c *CSVFormatter) Format(records []*model.Record) (string, error) {
	if len(records) == 0 {
		return "", errors.New("无可用数据")
	}
	var CSVOut string
	for _, r := range records {
		if r.Type == model.Income {
			CSVOut += fmt.Sprintf("%v,%v,%v,%v", r.ID, r.Category, r.Description, r.Amount)
		} else {
			CSVOut += fmt.Sprintf("%v,%v,%v,-%v", r.ID, r.Category, r.Description, r.Amount)
		}
	}
	return CSVOut, nil
}
