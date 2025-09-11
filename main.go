package main

import (
	"Hiro/Configs"
	"Hiro/Database"
	"Hiro/Models"
	"Hiro/Routes"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	fmt.Println("this is Database Connection Demo...")

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
	r.Use(sessions.Sessions("auth-session", store))

	Routes.RegisterUserRoutes(r)
	Routes.RegisterBlogRoutes(r)
	Routes.RegisterAuthRoutes(r)
	Routes.RegisterWebRoutes(r)

	r.Static("/assets", "./public/assets")

	// Load HTML templates
	r.LoadHTMLGlob("Resources/**/*.gohtml")

	r.Run(":8080")
}
