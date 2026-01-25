package service

import (
	"errors"
	"fmt"
	"go_task/model"
	"sort"
	"strings"
	"time"
)

type TaskManager struct {
	Task []*model.Task
}

type personSlice []*model.Task

func (p personSlice) Len() int {
	return len(p)
}

func (p personSlice) Less(i, j int) bool {
	if p[i].Priority != p[j].Priority {
		return p[i].Priority < p[j].Priority
	} else {
		return p[i].CreateAt.Before(p[j].CreateAt)
	}
}

func (p personSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
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

func (t *TaskManager) List() (string, error) {

	if len(t.Task) == 0 {
		return "", errors.New("无数据")
	}
	sort.Sort(personSlice(t.Task))
	var builder strings.Builder
	for _, list := range t.Task {
		line := fmt.Sprintf(
			"%-4d %-30s %-10d %-10v %s\n",
			list.ID,
			list.Title,
			list.Priority,
			list.Done,
			list.CreateAt.Format("2006-01-02 15:04:05"),
		)
		builder.WriteString(line)
	}
	return builder.String(), nil
}
