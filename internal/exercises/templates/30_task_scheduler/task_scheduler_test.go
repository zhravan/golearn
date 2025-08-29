package task_scheduler

import (
	"testing"
	"time"
)

func TestAddTask(t *testing.T) {
	scheduler := NewTaskScheduler()

	taskName := "Test Task"
	scheduledTime := time.Now().Add(1 * time.Hour)

	task, err := scheduler.AddTask(taskName, scheduledTime, func() {})
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if task.ID != 1 {
		t.Errorf("Expected task ID 1, got %d", task.ID)
	}

	if task.Name != taskName {
		t.Errorf("Expected task name %s, got %s", taskName, task.Name)
	}

	_, err = scheduler.AddTask("", time.Now(), func() {})
	if err == nil {
		t.Errorf("Expected error for empty task name, got nil")
	}
	if err.Code != 1 {
		t.Errorf("Expected error code 1, got %d", err.Code)
	}
}

func TestGetTask(t *testing.T) {
	scheduler := NewTaskScheduler()
	taskName := "Test Task"
	scheduledTime := time.Now().Add(1 * time.Hour)
	scheduler.AddTask(taskName, scheduledTime, func() {})

	task, err := scheduler.GetTask(1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if task.Name != taskName {
		t.Errorf("Expected task name %s, got %s", taskName, task.Name)
	}

	_, err = scheduler.GetTask(999)
	if err == nil {
		t.Errorf("Expected error for non-existent task, got nil")
	}
	if err.Code != 2 {
		t.Errorf("Expected error code 2, got %d", err.Code)
	}
}

func TestTaskIterator(t *testing.T) {
	scheduler := NewTaskScheduler()
	scheduler.AddTask("Task 1", time.Now(), func() {})
	scheduler.AddTask("Task 2", time.Now(), func() {})

	it := scheduler.Iterator()

	task, ok := it.Next()
	if !ok || task.Name != "Task 1" {
		t.Errorf("Expected Task 1, got %v", task)
	}

	task, ok = it.Next()
	if !ok || task.Name != "Task 2" {
		t.Errorf("Expected Task 2, got %v", task)
	}

	_, ok = it.Next()
	if ok {
		t.Errorf("Expected no more tasks, got one")
	}
}

