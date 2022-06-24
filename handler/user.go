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

func (h *userHandler) Login(c *gin.Context) {
	var input user.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := gin.H{"errors": errors}

		c.JSON(http.StatusUnprocessableEntity, helper.JsonResponse("Something went wrong", http.StatusUnprocessableEntity, "error", errorMessage))
		return
	}

	newUser, err := h.userService.Login(input)
	if err != nil {
		errorsMessage := gin.H{"errors": err.Error()}
		c.JSON(http.StatusBadRequest, helper.JsonResponse("Something went wrong", http.StatusBadRequest, "error", errorsMessage))
		return
	}

	formatter := user.FormatUser(newUser, "token")

	response := helper.JsonResponse("User has been logged in", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) EmailCheck(c *gin.Context) {
	var input user.EmailCheck

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := gin.H{"errors": errors}

		c.JSON(http.StatusUnprocessableEntity, helper.JsonResponse("Something went wrong", http.StatusUnprocessableEntity, "error", errorMessage))
		return
	}

	isAvailable, err := h.userService.EmailCheck(input)
	if err != nil {
		errorsMessage := gin.H{"errors": err.Error()}
		c.JSON(http.StatusBadRequest, helper.JsonResponse("Something went wrong", http.StatusBadRequest, "error", errorsMessage))
		return
	}

	data := gin.H{"is_available": isAvailable}

	var metaMessage string

	if isAvailable {
		metaMessage = "Email is available"
	} else {
		metaMessage = "Email is not available"
	}

	response := helper.JsonResponse(metaMessage, http.StatusOK, "success", data)

	c.JSON(http.StatusOK, response)
}