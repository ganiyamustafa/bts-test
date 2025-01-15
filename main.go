package main

import (
	"log"
	"os"

	"github.com/ganiyamustafa/bts/db/connections"
	"github.com/ganiyamustafa/bts/db/migrations"
	"github.com/ganiyamustafa/bts/db/seeders"
	"github.com/ganiyamustafa/bts/internal/routes"
	"github.com/ganiyamustafa/bts/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// init application root path
func init() {
	cwd, _ := os.Getwd()
	utils.SetRootPath(cwd + "/")
}

// init routes
func initRoute(router *gin.RouterGroup, handler *utils.Handler) {
	routes.AuthRoutes(router, handler)
	routes.TodoListRoutes(router, handler)
}

// init main app
func initMainApp() {
	// init gin
	app := gin.Default()
	app.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"*"},
		AllowHeaders:    []string{"*"},
	}))

	// init base handler
	handler := utils.Handler{}
	handler.Validator = validator.New()
	handler.Postgre = connections.Postgre

	// init gin routes
	routes := app.Group("/api/v1")
	initRoute(routes, &handler)

	// run gin
	if err := app.Run(utils.Env("HOST") + ":" + utils.Env("PORT")); err != nil {
		return
	}
}

// init app for migration command
func initMigratorApp() {
	var id string
	command := os.Args[1]

	if len(os.Args) >= 3 {
		id = os.Args[2]
	}

	switch command {
	case "migrate":
		migrations.Migrate(connections.Postgre)
	case "rollback":
		migrations.Rollback(connections.Postgre)
	case "seed":
		seeders.Seed(connections.Postgre, id)
	case "wipe":
		seeders.Wipe(connections.Postgre, id)
	}
}

func main() {
	// connect postgre
	if err := connections.ConnectPostgre(); err != nil {
		log.Fatal("Failed to connect to postgre: " + err.Error())
	}

	if len(os.Args) >= 2 {
		initMigratorApp()
	} else {
		initMainApp()
	}
}

// yg belum
// 1. add product to cart
// 2. checkout
// 3. give discount voucher

// optional
// 1. get all product
// 2. get all user voucher
