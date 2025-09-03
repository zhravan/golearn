package cli_todo_list

import (
	"os"
	"testing"
)

func TestNewTodoList(t *testing.T) {
	filepath := "test_todo.json"
	tl := NewTodoList(filepath)
	if tl == nil {
		t.Errorf("Expected a new TodoList, got nil")
	}
	if tl.filepath != filepath {
		t.Errorf("Expected filepath %s, got %s", filepath, tl.filepath)
	}
}

func TestAddAndComplete(t *testing.T) {
	filepath := "test_add_complete.json"
	defer os.Remove(filepath)

	tl := NewTodoList(filepath)
	tl.Add("Buy groceries")
	tl.Add("Walk the dog")

	if len(tl.Todos) != 2 {
		t.Errorf("Expected 2 todos, got %d", len(tl.Todos))
	}

	if tl.Todos[0].Task != "Buy groceries" || tl.Todos[1].Task != "Walk the dog" {
		t.Errorf("Tasks not added correctly")
	}

	// Complete the first task
	err := tl.Complete(1)
	if err != nil {
		t.Errorf("Unexpected error completing task: %v", err)
	}
	if !tl.Todos[0].Complete {
		t.Errorf("Task 1 should be complete")
	}

	// Try to complete a non-existent task
	err = tl.Complete(99)
	if err == nil {
		t.Errorf("Expected error for non-existent task, got nil")
	}
}

func TestLoadAndSave(t *testing.T) {
	filepath := "test_load_save.json"
	defer os.Remove(filepath)

	// Create a list and save it
	tl1 := NewTodoList(filepath)
	tl1.Add("Task A")
	tl1.Add("Task B")
	_ = tl1.Complete(1)
	err := tl1.Save()
	if err != nil {
		t.Fatalf("Failed to save todo list: %v", err)
	}

	// Load into a new list
	tl2 := NewTodoList(filepath)
	err = tl2.Load()
	if err != nil {
		t.Fatalf("Failed to load todo list: %v", err)
	}

	if len(tl2.Todos) != 2 {
		t.Errorf("Expected 2 todos after load, got %d", len(tl2.Todos))
	}
	if tl2.Todos[0].Task != "Task A" || !tl2.Todos[0].Complete || tl2.Todos[1].Task != "Task B" || tl2.Todos[1].Complete {
		t.Errorf("Loaded todos do not match saved state")
	}
}
