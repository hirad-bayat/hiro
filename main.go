package main

import (
	"Hiro/Database"
	"Hiro/Models"
	"Hiro/Routes"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	fmt.Println("this is Database Connection Demo...")

	Database.Connect()
	err := Database.DB.AutoMigrate(&Models.User{}, &Models.Blog{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	fmt.Println("✅ Database migration completed!")

	r := gin.Default()
	Routes.RegisterUserRoutes(r)
	Routes.RegisterBlogRoutes(r)

	r.Run(":8080")
}
