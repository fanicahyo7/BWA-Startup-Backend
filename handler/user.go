package handler

import (
	"bwastartup/helper"
	"bwastartup/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errormessagge := gin.H{"errors": errors}

		response := helper.ApiResponse("Register account failed", http.StatusUnprocessableEntity, "error", errormessagge)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.ApiResponse("Register account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	userformat := user.FormatUser(newUser, "tokentokentoken")

	response := helper.ApiResponse("Account has been registered", 200, "success", userformat)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) Login(c *gin.Context) {
	var input user.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errormessagge := gin.H{"errors": errors}

		response := helper.ApiResponse("login failed", http.StatusUnprocessableEntity, "error", errormessagge)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	loginuser, err := h.userService.LoginUser(input)
	if err != nil {
		errormessagge := gin.H{"errors": err.Error()}

		response := helper.ApiResponse("login failed", http.StatusUnprocessableEntity, "error", errormessagge)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	loginformat := user.FormatUser(loginuser, "tokentokentoken")
	response := helper.ApiResponse("login sukses", 200, "success", loginformat)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) CekEmailUser(c *gin.Context) {
	var input user.CekEmailInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errormessagge := gin.H{"errors": errors}

		response := helper.ApiResponse("cek email failed", http.StatusUnprocessableEntity, "error", errormessagge)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	cekemail, err := h.userService.CekEmailUser(input)
	if err != nil {
		errormessagge := gin.H{"errors": err.Error()}

		response := helper.ApiResponse("cek email failed", http.StatusUnprocessableEntity, "error", errormessagge)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{
		"is_available": cekemail,
	}

	message := "Email has been registered"
	if cekemail {
		message = "Email Available"
	}

	response := helper.ApiResponse(message, 200, "success", data)

	c.JSON(http.StatusOK, response)
}
