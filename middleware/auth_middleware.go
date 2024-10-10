package middleware

import (
	"go-booking/dto"
	"go-booking/service"
	"go-booking/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(jwtService service.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			res := utils.ReturnResponseError(401, dto.MESSAGE_FAILED_GET_TOKEN)
			c.AbortWithStatusJSON(http.StatusUnauthorized, res)
			return
		}
		if !strings.Contains(token, "Bearer ") {
			res := utils.ReturnResponseError(401, dto.MESSAGE_FAILED_TOKEN_NOT_VALID)
			c.AbortWithStatusJSON(http.StatusUnauthorized, res)
			return
		}
		token = strings.Replace(token, "Bearer ", "", -1)
		validatedToken, err := jwtService.ValidateToken(token)
		if err != nil {
			res := utils.ReturnResponseError(401, dto.MESSAGE_FAILED_TOKEN_NOT_VALID)
			c.AbortWithStatusJSON(http.StatusUnauthorized, res)
			return
		}
		if !validatedToken.Valid {
			res := utils.ReturnResponseError(401, dto.MESSAGE_FAILED_UNAUTHORIZED)
			c.AbortWithStatusJSON(http.StatusUnauthorized, res)
			return
		}
		userId, err := jwtService.GetUserIDByToken(validatedToken)
		if err != nil {
			res := utils.ReturnResponseError(401, dto.MESSAGE_FAILED_TOKEN_NOT_ASSOCIATED)
			c.AbortWithStatusJSON(http.StatusUnauthorized, res)
			return
		}
		c.Set("userId", userId)
		c.Set("token", token)
		c.Next()
	}
}
