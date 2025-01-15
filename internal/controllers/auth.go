package controllers

import (
	"net/http"

	"github.com/ganiyamustafa/bts/internal/requests"
	"github.com/ganiyamustafa/bts/internal/serializers"
	"github.com/ganiyamustafa/bts/internal/services"
	"github.com/ganiyamustafa/bts/utils"
	apperror "github.com/ganiyamustafa/bts/utils/app_error"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct {
	UserService services.UserService
}

func (u *AuthController) Login(ctx *gin.Context) {
	var payload requests.LoginRequest

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ErrorResponse(ctx, apperror.FromError(err).SetHttpCustomStatusCode(http.StatusBadRequest))
		return
	}

	if err := u.UserService.Handler.Validator.Struct(&payload); err != nil {
		ErrorResponse(ctx, apperror.FromError(err).SetHttpCustomStatusCode(http.StatusBadRequest))
		return
	}

	// check user email
	user, err := u.UserService.GetUserByEmail(payload.Email)
	if err != nil {
		ErrorResponse(ctx, apperror.New("Email or Password Invalid").SetHttpCustomStatusCode(http.StatusUnauthorized))
		return
	}

	// check user password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password)); err != nil {
		ErrorResponse(ctx, apperror.New("Email or Password Invalid").SetHttpCustomStatusCode(http.StatusUnauthorized))
		return
	}

	// generate auth token
	authToken, nErr := utils.EncodeJWT(map[string]string{"id": user.ID.String(), "email": user.Email}, []byte(utils.Env("SECRET_KEY")))
	if nErr != nil {
		ErrorResponse(ctx, apperror.FromError(err))
		return
	}

	SuccessResponse(ctx, serializers.LoginResponse{AuthToken: authToken}, nil, "login Successfully", http.StatusOK)
}

func (u *AuthController) Register(ctx *gin.Context) {
	var payload requests.RegisterRequest

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ErrorResponse(ctx, apperror.FromError(err).SetHttpCustomStatusCode(http.StatusBadRequest))
		return
	}

	if err := u.UserService.Handler.Validator.Struct(&payload); err != nil {
		ErrorResponse(ctx, apperror.FromError(err).SetHttpCustomStatusCode(http.StatusBadRequest))
		return
	}

	// validate existing user
	_, err := u.UserService.GetUserByEmail(payload.Email)
	if err == nil {
		ErrorResponse(ctx, apperror.New("Email Has Been Used").SetHttpCustomStatusCode(http.StatusConflict))
		return
	}

	// create user
	user, err := u.UserService.CreateUser(payload)
	if err != nil {
		ErrorResponse(ctx, err)
		return
	}

	// generate auth token
	authToken, nErr := utils.EncodeJWT(map[string]string{"id": user.ID.String(), "email": user.Email}, []byte(utils.Env("SECRET_KEY")))
	if nErr != nil {
		ErrorResponse(ctx, apperror.FromError(err))
		return
	}

	SuccessResponse(ctx, serializers.LoginResponse{AuthToken: authToken}, nil, "Register Successfully", http.StatusCreated)
}
