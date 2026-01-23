package bank

import (
	"encoding/json"
	"fmt"
	"os"
)

const DataFileName = "bank_data.json"

func LoadBankData() map[string]*Account {
	datafile, err := os.ReadFile(DataFileName)
	if err != nil {
		fmt.Printf("文件初始化失败！%v\n", err)
		return map[string]*Account{
			"1001": {Name: "ambre", Balance: 500},
		}
	}
	data := make(map[string]*Account)
	err = json.Unmarshal(datafile, &data)
	if err != nil {
		return make(map[string]*Account)
	} else {
		fmt.Printf("数据加载成功\n")
	}
	return data
}

func SaveBankData(data map[string]*Account) {
	jsonData, err := json.MarshalIndent(data, "", "   ")
	if err != nil {
		fmt.Printf("数据初始化失败[%v]\n", err)
		return
	}
	err = os.WriteFile(DataFileName, jsonData, 0644)
	if err != nil {
		fmt.Printf("文件写入失败[%v]\n", err)
	} else {
		fmt.Printf("文件写入成功\n")
	}
}
