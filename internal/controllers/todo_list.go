package controllers

import (
	"net/http"

	"github.com/ganiyamustafa/bts/internal/models"
	"github.com/ganiyamustafa/bts/internal/requests"
	"github.com/ganiyamustafa/bts/internal/serializers"
	"github.com/ganiyamustafa/bts/internal/services"
	"github.com/ganiyamustafa/bts/utils"
	apperror "github.com/ganiyamustafa/bts/utils/app_error"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TodoListController struct {
	TodoListService     services.TodoListService
	TodoListItemService services.TodoListItemService
}

func (u *TodoListController) GetTodoListByUserIds(ctx *gin.Context) {
	var query requests.GetTodoListsRequest
	var user = ctx.MustGet("user").(models.User)
	query.UserID = user.ID

	ctx.Bind(&query)

	if products, meta, err := u.TodoListService.GetTodoListByUserIds(&query); err != nil {
		ErrorResponse(ctx, err)
	} else {
		// serialize carts
		datas := utils.Map[*models.TodoList](products, func(i int, u *models.TodoList) *serializers.GetTodoListResponse {
			return serializers.GetTodoListResponse{}.FromModel(u)
		})

		SuccessResponse(ctx, datas, meta, "Get Todo Lists Successfully", http.StatusOK)
	}
}

func (u *TodoListController) CreateTodoList(ctx *gin.Context) {
	var user = ctx.MustGet("user").(models.User)
	var payload = requests.CreateTodoListRequest{
		UserID: user.ID,
	}

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ErrorResponse(ctx, apperror.FromError(err).SetHttpCustomStatusCode(http.StatusBadRequest))
		return
	}

	if err := u.TodoListService.Handler.Validator.Struct(&payload); err != nil {
		ErrorResponse(ctx, apperror.FromError(err).SetHttpCustomStatusCode(http.StatusBadRequest))
		return
	}

	// create flash sale product
	flashSaleProduct, err := u.TodoListService.CreateTodoList(&payload)
	if err != nil {
		ErrorResponse(ctx, err)
		return
	}

	SuccessResponse(ctx, serializers.CreateTodoListResponse{}.FromModel(flashSaleProduct), nil, "Create Todo List Successfully", http.StatusCreated)
}

func (u *TodoListController) DeleteTodoList(ctx *gin.Context) {
	var user = ctx.MustGet("user").(models.User)
	var id = ctx.Param("id")

	// create flash sale product
	err := u.TodoListService.DeleteTodoList(uuid.MustParse(id), user.ID)
	if err != nil {
		ErrorResponse(ctx, err)
		return
	}

	SuccessResponse(ctx, nil, nil, "delete Todo List Successfully", http.StatusCreated)
}

func (u *TodoListController) GetDetailTodoList(ctx *gin.Context) {
	var user = ctx.MustGet("user").(models.User)
	var id = ctx.Param("id")

	// create flash sale product
	todo_list, err := u.TodoListService.GetTodoListDetailByID(uuid.MustParse(id), user.ID)
	if err != nil {
		ErrorResponse(ctx, err)
		return
	}

	SuccessResponse(ctx, serializers.GetTodoListDetailResponse{}.FromModel(todo_list), nil, "Get Todo List Successfully", http.StatusCreated)
}

func (u *TodoListController) CreateTodoListItem(ctx *gin.Context) {
	var user = ctx.MustGet("user").(models.User)
	var todo_list_id = ctx.Param("todo_list_id")
	var payload = requests.CreateTodoListItemRequest{
		TodoListID: uuid.MustParse(todo_list_id),
		UserID:     user.ID,
	}

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ErrorResponse(ctx, apperror.FromError(err).SetHttpCustomStatusCode(http.StatusBadRequest))
		return
	}

	if err := u.TodoListItemService.Handler.Validator.Struct(&payload); err != nil {
		ErrorResponse(ctx, apperror.FromError(err).SetHttpCustomStatusCode(http.StatusBadRequest))
		return
	}

	// create flash sale product
	todo_list_item, err := u.TodoListItemService.CreateTodoListItem(&payload)
	if err != nil {
		ErrorResponse(ctx, err)
		return
	}

	SuccessResponse(ctx, serializers.CreateTodoListItemResponse{}.FromModel(todo_list_item), nil, "Create Todo List Item Successfully", http.StatusCreated)
}

func (u *TodoListController) CheckTodoListItem(ctx *gin.Context) {
	var user = ctx.MustGet("user").(models.User)
	var todo_list_item_id = ctx.Param("todo_list_item_id")

	// create flash sale product
	todo_list_item, err := u.TodoListItemService.CheckTodoListItemByID(uuid.MustParse(todo_list_item_id), user.ID)
	if err != nil {
		ErrorResponse(ctx, err)
		return
	}

	SuccessResponse(ctx, serializers.CreateTodoListItemResponse{}.FromModel(todo_list_item), nil, "Update Todo List Item Successfully", http.StatusCreated)
}

func (u *TodoListController) GetTodoListItemDetail(ctx *gin.Context) {
	var todo_list_item_id = ctx.Param("todo_list_item_id")

	// create flash sale product
	todo_list_item, err := u.TodoListItemService.GetTodoListItemDetailByID(uuid.MustParse(todo_list_item_id))
	if err != nil {
		ErrorResponse(ctx, err)
		return
	}

	SuccessResponse(ctx, serializers.CreateTodoListItemResponse{}.FromModel(todo_list_item), nil, "Update Todo List Item Successfully", http.StatusCreated)
}

func (u *TodoListController) DeleteTodoListItem(ctx *gin.Context) {
	var todo_list_item_id = ctx.Param("todo_list_item_id")

	// create flash sale product
	err := u.TodoListItemService.DeleteTodoListItem(uuid.MustParse(todo_list_item_id))
	if err != nil {
		ErrorResponse(ctx, err)
		return
	}

	SuccessResponse(ctx, nil, nil, "delete Todo List Item Successfully", http.StatusCreated)
}
