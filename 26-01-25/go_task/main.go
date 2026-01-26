package main

import (
	"flag"
	"fmt"
	"go_task/model"
	"go_task/service"
	"go_task/utils"
	"sync"
)

func main() {
	data, err := utils.LoadFile()
	manager := &service.TaskManager{Task: data}
	if err != nil {
		fmt.Println(err)
	}
	var wait sync.WaitGroup
	addCmd := flag.String("add", "", "添加任务的内容")
	addPriorityCmd := flag.Int("p", 0, "设置优先级(0:高, 1:中, 2:低)")
	listCmd := flag.Bool("list", false, "列出所有任务")
	doneCmd := flag.Int("done", 0, "完成指定ID的任务")
	flag.Parse()
	if *addCmd != "" {
		err := manager.Add(*addCmd, model.Priority(*addPriorityCmd))
		if err != nil {
			fmt.Printf("任务添加添加失败[%v]\n", err)
		}
		fmt.Printf("任务添加成功\n")
		wait.Add(1)
		go func(tasks []*model.Task) {
			defer wait.Done()
			utils.SavaFile(tasks)
		}(manager.Task)
	}
	if *doneCmd != 0 {
		err := manager.Complete(*doneCmd)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("完成任务成功")
		wait.Add(1)
		go func(tasks []*model.Task) {
			defer wait.Done()
			utils.SavaFile(tasks)
		}(manager.Task)
	}
	if *listCmd {
		list, err := manager.List()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(list)
	}
	wait.Wait()
}
