package routes

import (
	"github.com/ganiyamustafa/bts/internal/controllers"
	"github.com/ganiyamustafa/bts/internal/services"
	"github.com/ganiyamustafa/bts/utils"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.RouterGroup, handler *utils.Handler) {
	userService := services.UserService{Handler: handler}
	controller := controllers.AuthController{UserService: userService}

	router.POST("/register", controller.Register)
	router.POST("/login", controller.Login)
}
