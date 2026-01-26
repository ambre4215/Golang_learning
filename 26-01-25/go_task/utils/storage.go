package utils

import (
	"encoding/json"
	"errors"
	"go_task/model"
	"io"
	"os"
)

const FileName = "task_data.json"

func LoadFile() ([]*model.Task, error) {
	file, err := os.OpenFile(FileName, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return []*model.Task{}, errors.New("文件打开失败")
	}
	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, errors.New("文件读取失败")
	}
	defer file.Close()
	var data []*model.Task
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, errors.New("文件加载失败")
	}
	return data, nil
}

func SavaFile(data []*model.Task) error {
	if len(data) == 0 {
		return errors.New("无数据可保存")
	}
	jsonData, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return errors.New("数据序列化失败")
	}
	err = os.WriteFile(FileName, jsonData, 0644)
	if err != nil {
		return errors.New("数据保存失败")
	}
	return nil
}
