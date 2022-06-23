package handler

import (
	"golearn/helper"
	"golearn/models/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) Register(c *gin.Context) {
	var input user.RegisterUserInput
	
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := gin.H{"errors": errors}

		c.JSON(http.StatusUnprocessableEntity, helper.JsonResponse("Something went wrong", http.StatusUnprocessableEntity, "error", errorMessage))
		return
	}

	newUser, err := h.userService.Register(input)

	if err != nil {
		c.JSON(http.StatusBadRequest, helper.JsonResponse("Something went wrong", http.StatusBadRequest, "error", nil))
		return
	}

	formatter := user.FormatUser(newUser, "token")

	response := helper.JsonResponse("User has been registered", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}