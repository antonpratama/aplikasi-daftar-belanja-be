package main

import (
	"aplikasi-daftar-belanja/config"
	"aplikasi-daftar-belanja/routes"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"github.com/joho/godotenv"
)

func main() {
	//load env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env")
	}

	//Connect to DB
	config.ConnectDB()

	//start router
	r := gin.Default()

	//cors
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowMethods = []string{"POST", "GET", "PUT", "PATCH", "OPTIONS", "DELETE"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Type", "Authorization", "Accept", "User-Agent", "Cache-Control", "Pragma"}
	corsConfig.ExposeHeaders = []string{"Content-Length"}
	corsConfig.AllowCredentials = true
	r.Use(cors.New(corsConfig))

	//route daftar belanja
	routes.RegisterItemRoutes(r)

	//route user
	routes.RegisterUserRoutes(r)

	//dummy endpoint untuk cek server berjalan
	r.GET("/", func(c *gin.Context){
		c.JSON(200, gin.H{"message" : "Server is running"})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	fmt.Println("Server running on port " + port)
	r.Run(":" + port)
}