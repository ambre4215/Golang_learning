package service

import (
	"fmt"
	"go_task/model"
	"testing"
	"time"
)

func TestAdd(t *testing.T) {
	manager := &TaskManager{
		Task: []*model.Task{},
	}
	err := manager.Add("测试数据", model.High)

	if err != nil {
		t.Errorf("add方法不应返回错误[%v]", err)
	}
	if len(manager.Task) != 1 {
		t.Errorf("数据添加失败[%v]", err)
	}
	if manager.Task[0].Title != "测试数据" {
		t.Errorf("期望得到的title是“测试数据”,实际为%v [%v]", manager.Task[0].Title, err)
	}
}

func TestList(t *testing.T) {
	manager := &TaskManager{
		Task: []*model.Task{{
			ID:       1,
			Title:    "学习 Go 语言指针",
			Done:     true,
			CreateAt: time.Now().Add(-48 * time.Hour), // 两天前
			Priority: model.High,
		},
			{
				ID:       2,
				Title:    "编写 TaskManager 单元测试",
				Done:     false,
				CreateAt: time.Now().Add(-24 * time.Hour), // 昨天
				Priority: model.High,
			},
			{
				ID:       3,
				Title:    "买咖啡豆",
				Done:     false,
				CreateAt: time.Now().Add(-1 * time.Hour), // 一小时前
				Priority: model.Low,
			},
			{
				ID:       4,
				Title:    "整理 GitHub 仓库",
				Done:     false,
				CreateAt: time.Now(),
				Priority: model.Middle,
			}},
	}
	list, err := manager.List()
	if err != nil {
		t.Errorf("测试列表出错%v", err)
	}
	fmt.Println(list)
}
