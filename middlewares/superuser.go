package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func SuperUserMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := ctx.MustGet("user").(*jwt.MapClaims)
		//roleNames := (*user)["role"].(any)

		fmt.Println(user)
		//isSuperUser := slices.ContainsFunc(roleNames, func(roleName any) bool {
		//	roleNameStr := roleName.(string)
		//	return roleNameStr == constants.ROLE["SUPER_USER"]
		//})
		//
		//if !isSuperUser {
		//	ctx.Abort()
		//	utils.ApiResponse(ctx, http.StatusUnauthorized, "only superuser can access this endpoint", nil, nil)
		//	return
		//}

		ctx.Next()
	}
}
