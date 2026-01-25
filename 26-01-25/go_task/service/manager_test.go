package service

import (
	"go_task/model"
	"testing"
)

func TestAdd(t *testing.T) {
	// 1. 初始化一个空的 TaskManager
	manager := &TaskManager{
		Task: []*model.Task{},
	}

	// 2. 调用 Add 方法添加一个测试任务
	// 假设你定义的 Priority 类型中 model.High 是合法的
	err := manager.Add("测试任务", model.Middle)

	// 3. 判断是否返回了错误
	if err != nil {
		t.Errorf("Add 方法不应该返回错误，但是拿到了: %v", err)
	}

	// 4. 判断 Tasks 切片的长度是否变成了 1
	if len(manager.Task) != 1 {
		t.Errorf("期望任务数量为 1，但实际得到的是 %d", len(manager.Task))
	}

	// 5. 进一步验证数据准确性（可选，但推荐）
	if manager.Task[0].Title != "测试任务" {
		t.Errorf("期望标题为 '测试任务'，实际得到的是 %s", manager.Task[0].Title)
	}
}

func TestAddInvalid(t *testing.T) {
	manager := &TaskManager{}

	// 1. 测试标题为空的情况
	err := manager.Add("", model.High)
	if err == nil {
		t.Errorf("期望添加空标题时返回错误，但实际上 err 为 nil")
	}

	// 2. 测试优先级越界的情况 (假设你的优先级定义是 1, 2, 3)
	err = manager.Add("合法标题", model.Priority(99))
	if err == nil {
		t.Errorf("期望添加无效优先级时返回错误，但实际上 err 为 nil")
	}

	// 3. 确保虽然调用了 Add，但切片依然是空的
	if len(manager.Task) != 0 {
		t.Errorf("由于校验失败，任务列表应该保持为空，但实际长度为: %d", len(manager.Task))
	}
}

func TestAddMultiple(t *testing.T) {
	manager := &TaskManager{}

	// 1. 连续添加三个任务
	manager.Add("任务 A", model.High)
	manager.Add("任务 B", model.Middle)
	manager.Add("任务 C", model.Low)

	// 2. 检查长度
	if len(manager.Task) != 3 {
		t.Errorf("期望长度为 3，实际为 %d", len(manager.Task))
	}

	// 3. 检查 ID 自增是否正确 (应该是 1, 2, 3)
	for i, task := range manager.Task {
		expectedID := i + 1
		if task.ID != expectedID {
			t.Errorf("任务 [%s] 的 ID 错误：期望 %d，实际 %d", task.Title, expectedID, task.ID)
		}
	}
}
