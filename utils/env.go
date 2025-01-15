package utils

import (
	"os"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/joho/godotenv"
)

var RootPath *string

func Env(key string) string {
	rootPath := GetRootPath()

	if err := godotenv.Load(rootPath + ".env"); err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv(key)
}

func GetRootPath() string {
	return *RootPath
}

func SetRootPath(cwd string) {
	RootPath = &cwd
}
