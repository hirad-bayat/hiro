package main

import (
	"Hiro/Configs"
	"Hiro/Database"
	UserController "Hiro/Internal/User/Controllers"
	"Hiro/Internal/User/Repositories"
	"Hiro/Internal/User/Services"
	"Hiro/Middlewares"
	"Hiro/Models"
	"Hiro/Routes"
	"Hiro/docs"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
	"log"
	"net/http"
)

// @title Your API
// @version 1.0
// @description Your API Description
// @host localhost:8080
// @BasePath /api/
func main() {
	// Load configuration
	cfg, err := Configs.LoadConfig("./Configs")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	// Initialize database connection
	Database.Connect(cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.DBName)
	err = Database.DB.AutoMigrate(&Models.User{}, &Models.Blog{}, &Models.AccessToken{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	fmt.Println("✅ Database migration completed!")

	r := gin.Default()

	// Session middleware setup
	store := cookie.NewStore([]byte("your-secret-key")) // Change this to a secure random key
	store.Options(sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7, // 7 days
		HttpOnly: true,
		Secure:   false, // Set to true in production with HTTPS
		SameSite: http.SameSiteStrictMode,
	})
	r.Use(
		sessions.Sessions("auth-session", store),
		Middlewares.CORSMiddleware(),
	)

	// Initialize swagger
	docs.SwaggerInfo.Title = "Your API"
	docs.SwaggerInfo.Description = "Your API Description"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	// Swagger route - make sure this is exactly like this
	url := ginSwagger.URL("/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	//Routes.RegisterApiRoutes(r)

	// Initialize handlers
	userHandler := InitUserModule(Database.DB)

	// Setup router
	Routes.SetupRouter(userHandler, r)
	//Routes.RegisterBlogRoutes(r)
	//Routes.RegisterAuthRoutes(r)
	//Routes.RegisterWebRoutes(r)

	r.Static("/assets", "./public/assets")

	// Load HTML templates
	r.LoadHTMLGlob("Resources/**/*.gohtml")

	r.Run(":8080")
}

func InitUserModule(db *gorm.DB) *UserController.UserHandler {
	repo := Repositories.NewUserRepository(db)
	service := Services.NewUserService(repo)
	return UserController.NewUserHandler(service)
}
