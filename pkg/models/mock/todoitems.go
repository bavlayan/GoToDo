package mock

import (
	"time"

	"github.com/bavlayan/GoToDo/pkg/models"
)

var mockTodoItem = &models.TodoItems{
	ID:          "6cf57483-921b-4104-87f8-6337a4a59310",
	Completed:   false,
	Description: "mock todo item",
	CreatedDate: time.Now(),
	Deleted:     false,
}

type TodoItemModel struct{}

func (m *TodoItemModel) Save(description string) (int, error) {
	return 2, nil
}

func (m *TodoItemModel) Get(id string) (*models.TodoItems, error) {
	if id == "6cf57483-921b-4104-87f8-6337a4a59310" {
		return mockTodoItem, nil
	} else {
		return nil, models.ErrNoRecord
	}
}

func (m *TodoItemModel) GetDaily() ([]*models.TodoItems, error) {
	return []*models.TodoItems{mockTodoItem}, nil
}
