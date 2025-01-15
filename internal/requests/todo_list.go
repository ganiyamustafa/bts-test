package requests

import (
	"github.com/google/uuid"
)

type GetTodoListsRequest struct {
	PaginateRequest
	FilterRequest
	UserID uuid.UUID `json:"-" query:"-"`
}

type CreateTodoListRequest struct {
	UserID        uuid.UUID  `json:"-"`
	LabelID       *uuid.UUID `json:"label_id,omitempty"`
	IsPinned      bool       `json:"is_pinned"`
	Title         string     `json:"title" validate:"required"`
	TodoListItems []string   `json:"todo_list_items"`
}
