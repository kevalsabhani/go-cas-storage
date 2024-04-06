package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/kevalsabhani/go-cas-storage/config"
	"github.com/kevalsabhani/go-cas-storage/handlers"
)

func main() {

	config.LoadEnv()
	port := os.Getenv("PORT")
	// Create the storage directory if it doesn't exist
	if _, err := os.Stat(config.StoragePath); os.IsNotExist(err) {
		os.Mkdir(config.StoragePath, 0755)
	}

	r := gin.Default()

	setupRoutes(r)

	r.Run(port)
}

func setupRoutes(r *gin.Engine) {
	r.GET("/healthcheck", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"status": "OK"}) })

	// API v1 routes
	v1 := r.Group("/v1")
	v1.POST("/store", handlers.StoreContent)
	v1.GET("/retrieve/:key", handlers.RetrieveContent)

}
