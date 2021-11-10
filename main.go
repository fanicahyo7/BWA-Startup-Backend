package main

import (
	"bwastartup/auth"
	"bwastartup/campaign"
	"bwastartup/handler"
	"bwastartup/helper"
	"bwastartup/transaction"
	"bwastartup/user"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
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
	CampaignRepository := campaign.NewRepository(db)
	TransactionRepository := transaction.NewRepository(db)

	UserService := user.NewService(UserRepository)
	AuthService := auth.NewService()
	CampaignService := campaign.NewService(CampaignRepository)
	TransactionService := transaction.NewService(TransactionRepository, CampaignRepository)

	UserHandler := handler.NewUserHandler(UserService, AuthService)
	CampaignHandler := handler.NewCampaignHandler(CampaignService)
	TransactionHandler := handler.NewTransactionHandler(TransactionService)

	router := gin.Default()
	router.Static("/images", "./images")
	api := router.Group("/api/v1")
	api.POST("users/", UserHandler.RegisterUser)
	api.POST("sessions/", UserHandler.Login)
	api.POST("checkemail/", UserHandler.CekEmailUser)
	api.POST("avatars/", authMiddleware(AuthService, UserService), UserHandler.UploadAvatar)

	api.GET("campaigns/", CampaignHandler.GetCampaigns)
	api.GET("campaigns/:id", CampaignHandler.GetCampaign)
	api.POST("campaigns/", authMiddleware(AuthService, UserService), CampaignHandler.CreateCampaign)
	api.PUT("campaigns/:id", authMiddleware(AuthService, UserService), CampaignHandler.UpdateCampaign)
	api.POST("campaigns-images/", authMiddleware(AuthService, UserService), CampaignHandler.UploadImage)

	api.GET("campaigns/:id/transactions", authMiddleware(AuthService, UserService), TransactionHandler.GetCampaignTransactionByCampaignID)
	api.GET("transactions/", authMiddleware(AuthService, UserService), TransactionHandler.GetCampaignTransactionByUserID)

	router.Run()
}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.ApiResponse("unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.ApiResponse("unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := helper.ApiResponse("unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		UserID := int(claim["user_id"].(float64))

		user, err := userService.GetUserByID(UserID)
		if err != nil {
			response := helper.ApiResponse("unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		c.Set("currentUser", user)
	}
}
