package requests

import (
	"github.com/google/uuid"
)

type CreateTodoListItemRequest struct {
	UserID     uuid.UUID `json:"-"`
	TodoListID uuid.UUID `json:"-"`
	Name       string    `json:"name"`
}
