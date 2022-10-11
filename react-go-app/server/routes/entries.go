package routes

import (
	"github.com/gin-gonic/gin"
)

func StatusOK(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
	})
}

func AddEntry(c *gin.Context) {

}

func GetEntries(c *gin.Context) {

}

func GetEntryById(c *gin.Context) {

}

func UpdateEntry(c *gin.Context) {

}

func DeleteEntry(c *gin.Context) {

}
