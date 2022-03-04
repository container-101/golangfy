package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// load env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := ":" + os.Getenv("PORT")
	ginMode := os.Getenv("GIN_MODE")
	baseURL := os.Getenv("BASE_URL")

	gin.SetMode(ginMode)
	router := gin.Default()

	router.SetTrustedProxies([]string{baseURL})

	router.GET("/", func(c *gin.Context) {
		fmt.Printf("ClientIP: %s\n", c.ClientIP())
	})
	router.Run(port)
}
