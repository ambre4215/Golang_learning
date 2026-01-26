package utils

import (
	"fmt"
	"os"
	"time"
)

type Notifier interface {
	Log(msg string) error
	Print(msg string)
}

func Printer(msg string) {
	fmt.Printf("%s\n", msg)
}

func Logger(msg string) error {
	file, err := os.OpenFile("system.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return fmt.Errorf("log文件加载失败[%v]\n", err)
	}
	defer file.Close()
	if msg == "" {
		return fmt.Errorf("不能写入空信息\n")
	}
	logTime := time.Now().Format("[2006-01-02 15:04:05]")
	content := fmt.Sprintf("%s%s\n", logTime, msg)
	_, err = file.WriteString(content)
	if err != nil {
		return fmt.Errorf("Log写入失败[%v]\n", err)
	}
	return nil
}
