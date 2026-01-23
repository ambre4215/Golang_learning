package utils

import (
	"fmt"
	"os"
	"time"
)

type LogInterface interface {
	Log(msg string) error
}

type FileLogger struct {
	FileName string
}

func (f *FileLogger) Log(msg string) error {
	file, err := os.OpenFile(f.FileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("日志文件打开失败\n[%v]\n", err)
	}
	defer file.Close()

	_, err = file.WriteString(time.Now().Format("[2006-01-02 15:04:05]-") + msg + "\n")
	if err != nil {
		return fmt.Errorf("日志写入失败\n[%v]\n", err)
	}
	return nil
}
