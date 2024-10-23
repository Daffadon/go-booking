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
		Login(ctx *gin.Context)
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
		response := utils.ReturnResponseError(400, err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	res := utils.ReturnResponseSuccess(200, dto.MESSAGE_SUCCESS_REGISTER_USER, nil, nil)
	ctx.JSON(http.StatusOK, res)
}

func (u *userController) Login(ctx *gin.Context) {
	var req dto.UserLoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response := utils.ReturnResponseError(400, dto.MESSAGE_FAILED_GET_USER_DATA)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	token, err := u.userService.LoginUser(ctx.Request.Context(), req)
	if err != nil {
		response := utils.ReturnResponseError(400, dto.ErrEmailOrPasswordIsWrong.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	tokenResponse := map[string]string{
		"token": token,
	}
	res := utils.ReturnResponseSuccess(200, dto.MESSAGE_SUCCESS_LOGIN_USER, tokenResponse, nil)
	ctx.JSON(http.StatusOK, res)
}
