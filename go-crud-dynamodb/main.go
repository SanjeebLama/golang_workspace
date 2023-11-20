package main

import (
	"log"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	log.Println(strings.Repeat("-", 88))

	log.Println("Welcome to the Amazon DynamoDB ")
	log.Println(strings.Repeat("-", 88))

	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile | log.LUTC)
	log.Println("starting")

	router := gin.Default()    //initialize the gin router
	router.Use(gin.Logger())   // logger
	router.Use(cors.Default()) // cors

	router.Run("localhost:8080")

}
