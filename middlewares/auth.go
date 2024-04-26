package middlewares

import (
	"github.com/gin-gonic/gin"
	"main.go/utils"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			ctx.Abort()
			utils.ApiResponse(ctx, http.StatusUnauthorized, "unauthorized", nil, nil)
			return
		}

		token := strings.Replace(authHeader, "Bearer ", "", -1)

		claims, err := utils.VerifyToken(token)
		if err != nil {
			ctx.Abort()
			utils.ApiResponse(ctx, http.StatusUnauthorized, err.Error(), nil, nil)
			return
		}

		ctx.Set("user", claims)
	}
}
