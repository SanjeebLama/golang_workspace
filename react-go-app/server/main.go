package main

import (
	"os"

	"github.com/SanjeebLama/golang_workspace/tree/react-go-app/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	router := gin.Default()
	router.Use(gin.Logger())   // logger
	router.Use(cors.Default()) // cors
	router.POST("/entry/create", routes.AddEntry)
	router.GET("/entries", routes.GetEntries)
	router.GET("entry/:id", routes.GetEntryById)

	// Update
	router.PUT("/entry/:id", routes.UpdateEntry)

	// delete
	router.DELETE("/entry/:id", routes.DeleteEntry)

	router.Run(":" + port)
}
