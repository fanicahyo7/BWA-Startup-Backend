package main

import (
	"bwastartup/handler"
	"bwastartup/user"
	"fmt"
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

	userByEmail, err := UserRepository.FindByEmail("fanidc7@gmail.com")
	if err != nil {
		fmt.Println(err.Error())
	}

	if userByEmail.ID == 0 {
		fmt.Println("Data tidak ditemukan")
	} else {
		fmt.Println(userByEmail.Name)
	}

	UserHandler := handler.NewUserHandler(UserService)
	router := gin.Default()
	api := router.Group("/api/v1")
	api.POST("users/", UserHandler.RegisterUser)

	router.Run()

}
