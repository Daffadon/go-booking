package controller

import (
	"go-booking/dto"
	"go-booking/service"
	"go-booking/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	UserController interface {
		Register(ctx *gin.Context)
	}
	userController struct {
		userService service.UserService
	}
)

func NewUserController(us service.UserService) UserController {
	return &userController{
		userService: us,
	}
}

func (u *userController) Register(ctx *gin.Context) {
	var req dto.UserCreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response := utils.ReturnResponseError(400, dto.MESSAGE_FAILED_GET_USER_DATA)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	if err := u.userService.RegisterUser(ctx.Request.Context(), req); err != nil {
		response := utils.ReturnResponseError(400, dto.MESSAGE_FAILED_REGISTER_USER)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	res := utils.ReturnResponseSuccess(200, dto.MESSAGE_SUCCESS_REGISTER_USER, nil)
	ctx.JSON(http.StatusOK, res)
}
