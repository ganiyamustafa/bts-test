package services

import (
	"github.com/ganiyamustafa/bts/internal/models"
	"github.com/ganiyamustafa/bts/internal/requests"
	"github.com/ganiyamustafa/bts/utils"
	apperror "github.com/ganiyamustafa/bts/utils/app_error"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

type TodoListItemService struct {
	Handler *utils.Handler
}

// create todolist data
func (u *TodoListItemService) CreateTodoListItem(payload *requests.CreateTodoListItemRequest) (*models.TodoListItem, *apperror.AppError) {
	// copy payload to user models
	var todo_list_item models.TodoListItem
	copier.Copy(&todo_list_item, &payload)

	return &todo_list_item, apperror.FromError(u.Handler.Postgre.Create(&todo_list_item).Error)
}

func (u *TodoListItemService) CheckTodoListItemByID(ID uuid.UUID, userID uuid.UUID) (*models.TodoListItem, *apperror.AppError) {
	// copy payload to user models
	var todo_list_item models.TodoListItem
	if err := u.Handler.Postgre.Where("id = ?", ID).First(&todo_list_item).Error; err != nil {
		return nil, apperror.FromError(err)
	}

	return &todo_list_item, apperror.FromError(u.Handler.Postgre.Table("todo_list_items").Where("id = ?", ID).Update("is_checked", !todo_list_item.IsChecked).Error)
}

// delete todolist data
func (u *TodoListItemService) DeleteTodoListItem(ID uuid.UUID) *apperror.AppError {
	// copy payload to user models
	return apperror.FromError(u.Handler.Postgre.Table("todo_list_items").Where("id = ?", ID).Delete(nil).Error)
}

// get all todolist data
func (u *TodoListItemService) GetTodoListItemDetailByID(ID uuid.UUID) (*models.TodoListItem, *apperror.AppError) {
	var db = u.Handler.Postgre.Table("todo_list_items")
	var todo_list_item *models.TodoListItem

	// find todolist
	if err := db.Where("id = ?", ID).First(&todo_list_item).Error; err != nil {
		return nil, apperror.FromError(err)
	}

	return todo_list_item, nil
}
