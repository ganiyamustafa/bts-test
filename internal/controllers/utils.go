package controllers

import (
	"net/http"
	"time"

	"github.com/ganiyamustafa/bts/internal/serializers"
	apperror "github.com/ganiyamustafa/bts/utils/app_error"
	"github.com/gin-gonic/gin"
)

func ErrorResponse(ctx *gin.Context, err *apperror.AppError) {
	errMessage := err.Error()

	if err.HttpStatusCode() == http.StatusInternalServerError {
		errMessage = "something went wrong with our server"
	}

	ctx.AbortWithStatusJSON(err.HttpStatusCode(), serializers.BaseResponse{
		Status:     err.HttpStatusMessage(),
		StatusCode: err.HttpStatusCode(),
		Message:    errMessage,
		Timestamp:  time.Now(),
	})
}

func SuccessResponse(ctx *gin.Context, data any, meta *serializers.MetaResponse, message string, statusCode int) {
	ctx.JSON(statusCode, serializers.BaseResponse{
		Status:     "success",
		StatusCode: statusCode,
		Message:    message,
		Timestamp:  time.Now(),
		Data:       data,
		Meta:       meta,
	})
}
