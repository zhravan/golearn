package task_scheduler

import (
	"fmt"
	"time"
)

type Task struct {
	ID        int
	Name      string
	Scheduled time.Time
	Execute   func()
}

type SchedulerError struct {
	Code    int
	Message string
}

func (e *SchedulerError) Error() string {
	return fmt.Sprintf("Scheduler Error %d: %s", e.Code, e.Message)
}

type TaskScheduler struct {
	tasks  []*Task
	nextID int
}

func NewTaskScheduler() *TaskScheduler {
	return &TaskScheduler{
		tasks:  make([]*Task, 0),
		nextID: 1,
	}
}

func (ts *TaskScheduler) AddTask(name string, scheduled time.Time, execFn func()) (*Task, *SchedulerError) {
	if name == "" {
		return nil, &SchedulerError{Code: 1, Message: "Task name cannot be empty"}
	}

	task := &Task{
		ID:        ts.nextID,
		Name:      name,
		Scheduled: scheduled,
		Execute:   execFn,
	}
	ts.tasks = append(ts.tasks, task)
	ts.nextID++
	return task, nil
}

func (ts *TaskScheduler) GetTask(id int) (*Task, *SchedulerError) {
	for _, task := range ts.tasks {
		if task.ID == id {
			return task, nil
		}
	}
	return nil, &SchedulerError{Code: 2, Message: fmt.Sprintf("Task with ID %d not found", id)}
}

type TaskIterator struct {
	scheduler    *TaskScheduler
	currentIndex int
}

func (ts *TaskScheduler) Iterator() *TaskIterator {
	return &TaskIterator{scheduler: ts, currentIndex: 0}
}

func (it *TaskIterator) Next() (*Task, bool) {
	if it.currentIndex < len(it.scheduler.tasks) {
		task := it.scheduler.tasks[it.currentIndex]
		it.currentIndex++
		return task, true
	}
	return nil, false
}

func (ts *TaskScheduler) RunScheduledTasks() {
	now := time.Now()
	for _, task := range ts.tasks {
		if task.Scheduled.Before(now) {
			fmt.Printf("Executing task %s (ID: %d) at %s\n", task.Name, task.ID, now.Format(time.RFC3339))
			task.Execute()
		}
	}
}

