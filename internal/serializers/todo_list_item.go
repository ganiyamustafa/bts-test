package serializers

import (
	"time"

	"github.com/ganiyamustafa/bts/internal/models"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

type GetTodoListItemResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	IsChecked bool      `json:"is_checked"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u GetTodoListItemResponse) FromModel(m *models.TodoListItem) *GetTodoListItemResponse {
	var res GetTodoListItemResponse
	copier.CopyWithOption(&res, &m, copier.Option{IgnoreEmpty: true})
	return &res
}

type CreateTodoListItemResponse struct {
	ID uuid.UUID `json:"id"`
}

func (u CreateTodoListItemResponse) FromModel(m *models.TodoListItem) *CreateTodoListItemResponse {
	var res CreateTodoListItemResponse
	copier.CopyWithOption(&res, &m, copier.Option{IgnoreEmpty: true})
	return &res
}
