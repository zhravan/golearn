package cli_todo_list

// TODO:
// - Implement a simple JSON-backed todo list:
//   - NewTodoList: initialize with filepath and nextID starting at 1.
//   - Load/Save: read/write the list in JSON; handle missing/empty file gracefully.
//   - Add: append a new todo with incrementing ID.
//   - Complete: mark a todo as complete by ID or return an error if missing.

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
	// TODO: initialize todo list with filepath and nextID starting at 1
	return &TodoList{}
}

func (tl *TodoList) Load() error {
	// TODO: load todos from JSON file if present
	return nil
}

func (tl *TodoList) Save() error {
	// TODO: save todos to JSON file
	return nil
}

func (tl *TodoList) Add(task string) *Todo {
	// TODO: append a new todo and increment nextID
	return nil
}

func (tl *TodoList) Complete(id int) error {
	// TODO: mark todo as complete by ID or return error
	return nil
}
