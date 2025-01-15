package serializers

import (
	"time"

	"github.com/ganiyamustafa/bts/internal/models"
	"github.com/ganiyamustafa/bts/utils"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

type GetTodoListResponse struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"-"`
	IsPinned  bool      `json:"is_pinned"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func (u GetTodoListResponse) FromModel(m *models.TodoList) *GetTodoListResponse {
	var res GetTodoListResponse
	copier.CopyWithOption(&res, &m, copier.Option{IgnoreEmpty: true})

	return &res
}

type CreateTodoListResponse struct {
	ID uuid.UUID `json:"id"`
}

func (u CreateTodoListResponse) FromModel(m *models.TodoList) *CreateTodoListResponse {
	var res CreateTodoListResponse
	copier.CopyWithOption(&res, &m, copier.Option{IgnoreEmpty: true})
	return &res
}

type DeleteTodoListResponse struct {
	ID uuid.UUID `json:"id"`
}

func (u DeleteTodoListResponse) FromModel(m *models.TodoList) *CreateTodoListResponse {
	var res CreateTodoListResponse
	copier.CopyWithOption(&res, &m, copier.Option{IgnoreEmpty: true})
	return &res
}

type GetTodoListDetailResponse struct {
	ID            uuid.UUID                  `json:"id"`
	UserID        uuid.UUID                  `json:"-"`
	Label         *GetLabelResponse          `json:"label,omitempty"`
	TodoListItems []*GetTodoListItemResponse `json:"todo_list_items,omitempty"`
	IsPinned      bool                       `json:"is_pinned"`
	Title         string                     `json:"title"`
	CreatedAt     time.Time                  `json:"created_at,omitempty"`
	UpdatedAt     time.Time                  `json:"updated_at,omitempty"`
}

func (u GetTodoListDetailResponse) FromModel(m *models.TodoList) *GetTodoListDetailResponse {
	var res GetTodoListDetailResponse
	copier.CopyWithOption(&res, &m, copier.Option{IgnoreEmpty: true})

	if m.Label != nil {
		res.Label = GetLabelResponse{}.FromModel(m.Label)
	}

	res.TodoListItems = utils.Map[*models.TodoListItem, *GetTodoListItemResponse](m.TodoListItems, func(i int, tli *models.TodoListItem) *GetTodoListItemResponse {
		return GetTodoListItemResponse{}.FromModel(tli)
	})

	return &res
}
