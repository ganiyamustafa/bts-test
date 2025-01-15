package routes

import (
	"github.com/ganiyamustafa/bts/internal/controllers"
	"github.com/ganiyamustafa/bts/internal/services"
	"github.com/ganiyamustafa/bts/middlewares"
	"github.com/ganiyamustafa/bts/utils"
	"github.com/gin-gonic/gin"
)

func TodoListRoutes(router *gin.RouterGroup, handler *utils.Handler) {
	todoListService := services.TodoListService{Handler: handler}
	todoListItemService := services.TodoListItemService{Handler: handler}
	controller := controllers.TodoListController{TodoListService: todoListService, TodoListItemService: todoListItemService}

	router = router.Group("/todo-lists")
	router.GET("/", middlewares.IsUser, middlewares.AttachUserCtx, controller.GetTodoListByUserIds)
	router.POST("/", middlewares.IsUser, middlewares.AttachUserCtx, controller.CreateTodoList)
	router.GET("/:id", middlewares.IsUser, middlewares.AttachUserCtx, controller.GetDetailTodoList)
	router.DELETE("/:id", middlewares.IsUser, middlewares.AttachUserCtx, controller.DeleteTodoList)

	router.POST("/:todo_list_id/items", middlewares.IsUser, middlewares.AttachUserCtx, controller.CreateTodoListItem)
	router.POST("/:todo_list_id/items/:todo_list_item_id/check", middlewares.IsUser, middlewares.AttachUserCtx, controller.CheckTodoListItem)
	router.POST("/:todo_list_id/items/:todo_list_item_id/delete", middlewares.IsUser, middlewares.AttachUserCtx, controller.DeleteTodoListItem) // must be delete but got error
	router.POST("/:todo_list_id/items/:todo_list_item_id/detail", middlewares.IsUser, middlewares.AttachUserCtx, controller.GetDetailTodoList)  // must be get but got error T.T
}
