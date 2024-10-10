package dto

import "errors"

var (
	ErrEmailAlreadyExists     = errors.New("email already exist")
	ErrEmailOrPasswordIsWrong = errors.New("email or password is wrong")
)

var (
	// Failed
	MESSAGE_FAILED_REGISTER_USER        = "Failed to register user"
	MESSAGE_FAILED_GET_USER_DATA        = "Failed to get user data"
	MESSAGE_FAILED_LOGIN_USER           = "Email or password is wrong"
	MESSAGE_FAILED_GET_TOKEN            = "Token is missing"
	MESSAGE_FAILED_TOKEN_NOT_VALID      = "Token is invalid"
	MESSAGE_FAILED_UNAUTHORIZED         = "Unauthorized"
	MESSAGE_FAILED_TOKEN_NOT_ASSOCIATED = "Token is not associated with any user"
	// Success
	MESSAGE_SUCCESS_REGISTER_USER = "User successfully registered"
	MESSAGE_SUCCESS_LOGIN_USER    = "User successfully login"
)

type (
	UserCreateRequest struct {
		Name       string `json:"name" form:"name" binding:"required"`
		Email      string `json:"email" form:"email" binding:"required"`
		TelpNumber string `json:"telp_number" form:"telp_number" binding:"required"`
		Age        uint8  `json:"age" form:"age" binding:"required"`
		Password   string `json:"password" form:"password" binding:"required"`
	}

	UserCreateResponse struct {
		Message string `json:"message"`
	}

	UserLoginRequest struct {
		Email    string `json:"email" form:"email" binding:"required"`
		Password string `json:"password" form:"password" binding:"required"`
	}
)
