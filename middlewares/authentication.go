package middlewares

import (
	"strings"

	"github.com/ganiyamustafa/bts/db/connections"
	"github.com/ganiyamustafa/bts/internal/models"
	"github.com/ganiyamustafa/bts/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type AuthClaims struct {
	jwt.StandardClaims
	Payload map[string]interface{} `json:"payload"`
}

// user authentication middleware
func IsUser(ctx *gin.Context) {
	header := ctx.Request.Header
	authorization := header["Authorization"]

	if len(authorization) == 0 {
		ctx.AbortWithStatusJSON(401, gin.H{"error": "Invalid token"})
		return
	}

	tokenString := strings.Split(authorization[0], " ")[1]
	token, err := utils.DecodeJWT(tokenString)
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"error": "Failed decoding jwt token"})
		return
	}

	claims := token.Claims.(jwt.MapClaims)
	email := claims["payload"].(map[string]interface{})["email"]

	ctx.Set("email", email)
	ctx.Next()
}

// attach user to context middleware after authentication
func AttachUserCtx(ctx *gin.Context) {
	db := connections.Postgre

	var user models.User
	db.Where("email = ?", ctx.MustGet("email")).First(&user)

	ctx.Set("user", user)
	ctx.Next()
}
