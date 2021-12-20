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
	"path/filepath"
	"strings"

	webHandler "bwastartup/web/handler"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
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

	webUserHandler := webHandler.NewUserHandler(UserService)
	webCampaignhandler := webHandler.NewCampaignHandler(CampaignService, UserService)
	webTransactionHandler := webHandler.NewTransactionHandler(TransactionService)
	webSessionHandler := webHandler.NewSessionHandler(UserService)

	router := gin.Default()
	router.Use(cors.Default())

	cookieStore := cookie.NewStore([]byte(auth.SECRET_KEY))
	router.Use(sessions.Sessions("bwastartup", cookieStore))

	router.HTMLRender = loadTemplates("./web/templates")

	router.Static("/images", "./images")
	router.Static("/css", "./web/assets/css")
	router.Static("/js", "./web/assets/js")
	router.Static("/webfonts", "./web/assets/webfonts")

	api := router.Group("/api/v1")
	api.POST("users/", UserHandler.RegisterUser)
	api.POST("sessions/", UserHandler.Login)
	api.POST("checkemail/", UserHandler.CekEmailUser)
	api.POST("avatars/", authMiddleware(AuthService, UserService), UserHandler.UploadAvatar)
	api.GET("users/fetch", authMiddleware(AuthService, UserService), UserHandler.FetchUser)

	api.GET("campaigns/", CampaignHandler.GetCampaigns)
	api.GET("campaigns/:id", CampaignHandler.GetCampaign)
	api.POST("campaigns/", authMiddleware(AuthService, UserService), CampaignHandler.CreateCampaign)
	api.PUT("campaigns/:id", authMiddleware(AuthService, UserService), CampaignHandler.UpdateCampaign)
	api.POST("campaigns-images/", authMiddleware(AuthService, UserService), CampaignHandler.UploadImage)

	api.GET("campaigns/:id/transactions", authMiddleware(AuthService, UserService), TransactionHandler.GetCampaignTransactionByCampaignID)
	api.GET("transactions/", authMiddleware(AuthService, UserService), TransactionHandler.GetCampaignTransactionByUserID)
	api.POST("transactions/", authMiddleware(AuthService, UserService), TransactionHandler.CreateTransaction)

	router.GET("users/", authAdminMiddleware(), webUserHandler.Index)
	router.GET("users/new", webUserHandler.New)
	router.POST("users/", webUserHandler.Create)
	router.GET("users/edit/:id", webUserHandler.Edit)
	router.POST("users/update/:id", authAdminMiddleware(), webUserHandler.Update)
	router.GET("users/avatar/:id", authAdminMiddleware(), webUserHandler.NewAvatar)
	router.POST("users/avatar/:id", authAdminMiddleware(), webUserHandler.CreateAvatar)

	router.GET("campaigns/", authAdminMiddleware(), webCampaignhandler.Index)
	router.GET("campaigns/new", authAdminMiddleware(), webCampaignhandler.New)
	router.POST("campaigns/", authAdminMiddleware(), webCampaignhandler.Create)
	router.GET("campaigns/image/:id", authAdminMiddleware(), webCampaignhandler.NewImage)
	router.POST("campaigns/image/:id", authAdminMiddleware(), webCampaignhandler.CreateImage)
	router.GET("campaigns/edit/:id", authAdminMiddleware(), webCampaignhandler.Edit)
	router.POST("campaigns/update/:id", authAdminMiddleware(), webCampaignhandler.Update)
	router.GET("campaigns/show/:id", authAdminMiddleware(), webCampaignhandler.Show)

	router.GET("transactions/", authAdminMiddleware(), webTransactionHandler.Index)

	router.GET("login/", webSessionHandler.New)
	router.POST("/session", webSessionHandler.Create)
	router.GET("/logout", webSessionHandler.Destroy)

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

func authAdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)

		userIDSession := session.Get("userID")

		if userIDSession == nil {
			c.Redirect(http.StatusFound, "/login")
			return
		}
	}
}

func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	layouts, err := filepath.Glob(templatesDir + "/layouts/*")
	if err != nil {
		panic(err.Error())
	}

	includes, err := filepath.Glob(templatesDir + "/**/*")
	if err != nil {
		panic(err.Error())
	}

	// Generate our templates map from our layouts/ and includes/ directories
	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)
		r.AddFromFiles(filepath.Base(include), files...)
	}
	return r
}
