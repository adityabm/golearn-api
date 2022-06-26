package main

import (
	"golearn/auth"
	"golearn/handler"
	"golearn/models/user"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	_ "golearn/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {

    err := godotenv.Load(".env")

    if err != nil {
        log.Fatal("Error loading .env file")
    }
}

// @title Crowdfunding API v1
// @version 1.0.0

// @host localhost:8822
// @BasePath /api/v1

func main() {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	authService := auth.NewService()
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService, authService)

	router := gin.Default()
	api := router.Group("/api/v1")


	api.POST("/user", userHandler.Register)
	api.POST("/login", userHandler.Login)
	api.POST("/email-check", userHandler.EmailCheck)
	api.POST("/upload-profile-picture", userHandler.UploadProfilePicture)
	
	envType := os.Getenv("APP_ENV")
	 
	if envType != "production" {
		router.GET("/documentation/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}

	appPort := os.Getenv("APP_PORT")
	router.Run(":" + appPort)
}