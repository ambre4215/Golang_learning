package service

import (
	"errors"
	"go_task/model"
	"time"
)

type TaskManager struct {
	Task []*model.Task
}

func (t *TaskManager) MaxId() int {
	maxID := 0
	for _, item := range t.Task {
		if item.ID > maxID {
			maxID = item.ID
		}
	}
	maxID = maxID + 1
	return maxID
}

func (t *TaskManager) Add(title string, priority model.Priority) error {
	if title == "" {
		return errors.New("标题不能为空")
	}
	if priority < model.High || priority > model.Low {
		return errors.New("无效等级参数")
	}
	newTask := &model.Task{
		ID:       t.MaxId(),
		Title:    title,
		Done:     false,
		CreateAt: time.Now(),
		Priority: priority,
	}
	t.Task = append(t.Task, newTask)
	return nil
}
