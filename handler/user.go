package handler

import (
	"bwastartup/auth"
	"bwastartup/helper"
	"bwastartup/user"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
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

	token, err := h.authService.GenerateToken(newUser.ID)
	if err != nil {
		response := helper.ApiResponse("Register account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	userformat := user.FormatUser(newUser, token)

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

	token, err := h.authService.GenerateToken(loginuser.ID)
	if err != nil {
		response := helper.ApiResponse("login failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	loginformat := user.FormatUser(loginuser, token)
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

func (h *userHandler) UploadAvatar(c *gin.Context) {
	file, err := c.FormFile("avatar")
	if err != nil {
		errormessagge := gin.H{"is_uploaded": false}

		response := helper.ApiResponse("upload avatar failed", http.StatusUnprocessableEntity, "error", errormessagge)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	UserID := 9
	// path := "images/" + file.Filename
	path := fmt.Sprintf("images/%d-%s", UserID, file.Filename)
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		errormessagge := gin.H{"is_uploaded": false}

		response := helper.ApiResponse("upload avatar failed", http.StatusUnprocessableEntity, "error", errormessagge)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.userService.UpdateAvatar(UserID, path)
	if err != nil {
		errormessagge := gin.H{"is_uploaded": false}

		response := helper.ApiResponse("upload avatar failed", http.StatusUnprocessableEntity, "error", errormessagge)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := helper.ApiResponse("upload success", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)

}
