package middleware

import (
	"go-booking/dto"
	"go-booking/service"
	"go-booking/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RoleMiddleware(jwtService service.JWTService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		role, ok := ctx.Get("role")
		if !ok {
			res := utils.ReturnResponseError(401, dto.MESSAGE_FAILED_TOKEN_NOT_ASSOCIATED)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
			return
		}
		if role != "admin" {
			res := utils.ReturnResponseError(401, dto.MESSAGE_FAILED_UNAUTHORIZED)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
			return
		}
		ctx.Next()
	}
}
