package main

import (
	"bwastartup/auth"
	"bwastartup/handler"
	"bwastartup/user"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	UserRepository := user.NewRepository(db)
	UserService := user.NewService(UserRepository)
	AuthService := auth.NewService()

	UserHandler := handler.NewUserHandler(UserService, AuthService)
	router := gin.Default()
	api := router.Group("/api/v1")
	api.POST("users/", UserHandler.RegisterUser)
	api.POST("sessions/", UserHandler.Login)
	api.POST("checkemail/", UserHandler.CekEmailUser)
	api.POST("avatars/", UserHandler.UploadAvatar)
	router.Run()

}
