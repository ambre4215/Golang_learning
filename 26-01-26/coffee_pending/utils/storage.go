package utils

import (
	"coffee_pending/model"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

const FileName = "turnover.json"

func Load() ([]*model.Order, error) {
	file, err := os.OpenFile(FileName, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return []*model.Order{}, fmt.Errorf("打开文件失败%v\n", err)
	}
	jsonData, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("文件读取失败%v\n", err)
	}
	defer file.Close()
	var data []*model.Order
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		return nil, fmt.Errorf("文件反序列化失败%v\n", err)
	}
	return data, nil
}

func Save(data []*model.Order) error {
	file, err := os.OpenFile(FileName, os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("文件打开失败%v\b", err)
	}
	defer file.Close()
	jsonData, err := json.MarshalIndent(file, "", "\t")
	if err != nil {
		return fmt.Errorf("文件序列化失败%v\n", err)
	}
	_, err = file.Write(jsonData)
	if err != nil {
		return fmt.Errorf("文件写入失败%v\n", err)
	}
	return nil
}
