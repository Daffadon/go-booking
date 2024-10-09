package dto

import "errors"

var (
	ErrEmailAlreadyExists = errors.New("email already exist")
)

var (
	// Failed
	MESSAGE_FAILED_REGISTER_USER = "Failed to register user"
	MESSAGE_FAILED_GET_USER_DATA = "Failed to get user data"

	// Success
	MESSAGE_SUCCESS_REGISTER_USER = "User successfully registered"
)

type (
	UserCreateRequest struct {
		Name       string `json:"name" form:"name" binding:"required"`
		Email      string `json:"email" form:"email" binding:"required"`
		TelpNumber string `json:"telp_number" form:"telp_number" binding:"required"`
		Age        int    `json:"age" form:"age" binding:"required"`
		Password   string `json:"password" form:"password" binding:"required"`
	}

	UserCreateResponse struct {
		Message string `json:"message"`
	}
)
