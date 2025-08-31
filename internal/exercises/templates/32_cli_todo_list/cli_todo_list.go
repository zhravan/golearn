package cli_todo_list

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

type Todo struct {
	ID       int    `json:"id"`
	Task     string `json:"task"`
	Complete bool   `json:"complete"`
}

type TodoList struct {
	Todos    []Todo `json:"todos"`
	nextID   int
	filepath string
}

func NewTodoList(filepath string) *TodoList {
	return &TodoList{
		Todos:    []Todo{},
		nextID:   1,
		filepath: filepath,
	}
}

func (tl *TodoList) Load() error {
	if _, err := os.Stat(tl.filepath); errors.Is(err, os.ErrNotExist) {
		return nil // File does not exist, start with empty list
	}

	data, err := ioutil.ReadFile(tl.filepath)
	if err != nil {
		return fmt.Errorf("failed to read todo file: %w", err)
	}

	if len(data) == 0 {
		return nil // Empty file, no todos
	}

	if err := json.Unmarshal(data, &tl.Todos); err != nil {
		return fmt.Errorf("failed to unmarshal todos: %w", err)
	}
	// Find max ID to set nextID correctly
	for _, todo := range tl.Todos {
		if todo.ID >= tl.nextID {
			tl.nextID = todo.ID + 1
		}
	}
	return nil
}

func (tl *TodoList) Save() error {
	data, err := json.MarshalIndent(tl.Todos, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal todos: %w", err)
	}

	return ioutil.WriteFile(tl.filepath, data, 0644)
}

func (tl *TodoList) Add(task string) *Todo {
	todo := Todo{
		ID:       tl.nextID,
		Task:     task,
		Complete: false,
	}
	tl.Todos = append(tl.Todos, todo)
	tl.nextID++
	return &todo
}

func (tl *TodoList) Complete(id int) error {
	for i := range tl.Todos {
		if tl.Todos[i].ID == id {
			tl.Todos[i].Complete = true
			return nil
		}
	}
	return fmt.Errorf("todo with ID %d not found", id)
}

