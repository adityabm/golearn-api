package handler

import (
	"fmt"
	"golearn/helper"
	"golearn/models/user"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

// Register godoc
// @Summary First step for Registering user
// @Description Register new user from form than insert it to Database and it will response the user with token
// @Tags Auth API
// @Accept json
// @Produce plain
// @Param data body user.RegisterUserInput true "User info"
// @Success 200 {object} helper.Response 
// @Failure 400,422 {object} helper.Response
// @Router /user [post]
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
		errorMessage := gin.H{"errors": err.Error()}
		c.JSON(http.StatusBadRequest, helper.JsonResponse("Something went wrong", http.StatusBadRequest, "error", errorMessage))
		return
	}

	formatter := user.FormatUser(newUser, "token")

	response := helper.JsonResponse("User has been registered", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

// Login godoc
// @Summary First step for Loging in user
// @Description Login user and it will response the user with token
// @Tags Auth API
// @Accept json
// @Produce plain
// @Param data body user.LoginInput true "Login info"
// @Success 200 {object} helper.Response 
// @Failure 400,422 {object} helper.Response
// @Router /login [post]
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

// Email Check godoc
// @Summary To check availability of email
// @Description Check availability of email in Database
// @Tags Auth API
// @Accept json
// @Produce plain
// @Param data body user.EmailCheck true "Email"
// @Success 200 {object} helper.Response 
// @Failure 400,422 {object} helper.Response
// @Router /email-check [post]
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

func (h *userHandler) UploadProfilePicture(c *gin.Context) {
	file, err := c.FormFile("profile_picture")
	if err != nil {
		errorMessage := gin.H{"is_uploaded": false}

		c.JSON(http.StatusBadRequest, helper.JsonResponse("Failed to Upload Profile Picture", http.StatusBadRequest, "error", errorMessage))
		return
	}

	userID := 1

	// Set name to epoch
	now := time.Now()
	epoch := now.Unix()

	path := fmt.Sprintf("uploads/profile_pictures/%d-%s", epoch, file.Filename)
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		errorMessage := gin.H{"is_uploaded": false}

		c.JSON(http.StatusBadRequest, helper.JsonResponse("Failed to Upload Profile Picture", http.StatusBadRequest, "error", errorMessage))
		return
	}

	// TODO : Change it to JWT Token result for user
	_, err = h.userService.SaveProfilePicture(userID, path)
	if err != nil {
		errorMessage := gin.H{"is_uploaded": false}

		c.JSON(http.StatusBadRequest, helper.JsonResponse("Failed to Upload Profile Picture", http.StatusBadRequest, "error", errorMessage))
		return
	}

	data := gin.H{"is_uploaded": true}
	c.JSON(http.StatusOK, helper.JsonResponse("Profile Picture has been uploaded", http.StatusOK, "success", data))
}