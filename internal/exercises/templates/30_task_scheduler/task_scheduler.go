package task_scheduler

import (
    "time"
)

// TODO:
// - Build a simple in-memory task scheduler:
//   - AddTask: validate input and schedule a task with an auto-increment ID.
//   - GetTask: fetch by ID or return a typed error when missing.
//   - Iterator: provide Next() to walk scheduled tasks.
//   - RunScheduledTasks: execute tasks scheduled before now.
// - Keep signatures; tests assert error codes/messages and iteration behavior.

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
    // TODO: format error text
    return ""
}

type TaskScheduler struct {
    tasks  []*Task
    nextID int
}

func NewTaskScheduler() *TaskScheduler {
    // TODO: initialize scheduler state
    return &TaskScheduler{}
}

func (ts *TaskScheduler) AddTask(name string, scheduled time.Time, execFn func()) (*Task, *SchedulerError) {
    // TODO: validate and append to scheduler
    return nil, nil
}

func (ts *TaskScheduler) GetTask(id int) (*Task, *SchedulerError) {
    // TODO: find task by ID or return error
    return nil, nil
}

type TaskIterator struct {
	scheduler    *TaskScheduler
	currentIndex int
}

func (ts *TaskScheduler) Iterator() *TaskIterator {
    // TODO: return an iterator over tasks
    return &TaskIterator{}
}

func (it *TaskIterator) Next() (*Task, bool) {
    // TODO: return next task if available
    return nil, false
}

func (ts *TaskScheduler) RunScheduledTasks() {
    // TODO: execute scheduled tasks
}
