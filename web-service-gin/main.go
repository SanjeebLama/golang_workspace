package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albus slice to seed record album data.

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums add an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, gin.H{"message": "Album added successfully."})
}

func getAlbumsByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// deleteAlbum removes an album from the list by its ID.
func deleteAlbum(c *gin.Context) {
	id := c.Param("id")

	for i, a := range albums {
		if a.ID == id {
			// Delete the album from the slice.
			albums = append(albums[:i], albums[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "album deleted"})
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// putAlbum replaces an album's data.
func putAlbum(c *gin.Context) {
	id := c.Param("id")
	var updatedAlbum album

	if err := c.BindJSON(&updatedAlbum); err != nil {
		return
	}

	for i, a := range albums {
		if a.ID == id {
			albums[i] = updatedAlbum
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Album updated successfully."})
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// patchAlbum updates an album's data partially.
func patchAlbum(c *gin.Context) {
	id := c.Param("id")
	var patchData album

	if err := c.BindJSON(&patchData); err != nil {
		return
	}

	for i, a := range albums {
		if a.ID == id {
			// Update fields if present in patchData
			if patchData.Title != "" {
				albums[i].Title = patchData.Title
			}
			if patchData.Artist != "" {
				albums[i].Artist = patchData.Artist
			}
			if patchData.Price != 0 {
				albums[i].Price = patchData.Price
			}
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Album updated successfully."})
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
	router := gin.Default()    //initialize the gin router
	router.Use(gin.Logger())   // logger
	router.Use(cors.Default()) // cors

	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/albums/:id", getAlbumsByID)
	router.DELETE("/albums/:id", deleteAlbum)
	router.PUT("/albums/:id", putAlbum)
	router.PATCH("/albums/:id", patchAlbum)

	router.Run("localhost:8080")
}
