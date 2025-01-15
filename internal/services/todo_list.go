package services

import (
	"net/http"

	"github.com/ganiyamustafa/bts/internal/models"
	"github.com/ganiyamustafa/bts/internal/requests"
	"github.com/ganiyamustafa/bts/internal/serializers"
	"github.com/ganiyamustafa/bts/utils"
	apperror "github.com/ganiyamustafa/bts/utils/app_error"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

type TodoListService struct {
	Handler *utils.Handler
}

// get all todolist data
func (u *TodoListService) GetTodoListByUserIds(payload *requests.GetTodoListsRequest) ([]*models.TodoList, *serializers.MetaResponse, *apperror.AppError) {
	var db = u.Handler.Postgre.Table("todo_lists")
	var todo_lists []*models.TodoList
	var metaResponse serializers.MetaResponse

	// count all data
	if err := db.Count(&metaResponse.Total).Error; err != nil {
		return nil, nil, apperror.FromError(err)
	}

	// generate paginate data
	metaResponse.GeneratePaginateData(payload.Limit, payload.Page)

	// filter database
	db = db.Scopes(
		models.UtilScopes{}.PaginateScope(payload.PaginateRequest),
		models.UtilScopes{}.OrderByScope(payload.FilterRequest),
		models.TodoListScopes{}.SearchScope(payload.Search),
	)

	// find todolist
	if err := db.Where("user_id = ?", payload.UserID).Find(&todo_lists).Error; err != nil {
		return nil, nil, apperror.FromError(err)
	}

	return todo_lists, &metaResponse, nil
}

// get all todolist data
func (u *TodoListService) GetTodoListDetailByID(ID uuid.UUID, userID uuid.UUID) (*models.TodoList, *apperror.AppError) {
	var db = u.Handler.Postgre.Table("todo_lists")
	var todo_list *models.TodoList

	// filter database
	db = db.Scopes(
		models.TodoListScopes{}.PreloadTodoListItem(nil),
		models.TodoListScopes{}.PreloadLabel(nil),
	)

	// find todolist
	if err := db.Where("id = ?", ID).First(&todo_list).Error; err != nil {
		return nil, apperror.FromError(err)
	}

	if todo_list.UserID != userID {
		return nil, apperror.New("Access Forbidden").SetHttpCustomStatusCode(http.StatusForbidden)
	}

	return todo_list, nil
}

// create todolist data
func (u *TodoListService) CreateTodoList(payload *requests.CreateTodoListRequest) (*models.TodoList, *apperror.AppError) {
	// copy payload to user models
	var todo_list models.TodoList
	copier.Copy(&todo_list, &payload)

	todo_list.TodoListItems = models.TodoListItem{}.FromArrayString(payload.TodoListItems)

	return &todo_list, apperror.FromError(u.Handler.Postgre.Create(&todo_list).Error)
}

// delete todolist data
func (u *TodoListService) DeleteTodoList(ID uuid.UUID, userID uuid.UUID) *apperror.AppError {
	// copy payload to user models
	return apperror.FromError(u.Handler.Postgre.Table("todo_lists").Where("user_id = ? and id = ?", userID, ID).Delete(nil).Error)
}
