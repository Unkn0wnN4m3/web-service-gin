// Go and Gin simple RESTful web service form go.dev

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album
type album struct {
	ID      string  `json:"id"`
	Title   string  `json:"title"`
	Artist  string  `json:"artist"`
	Price   float64 `json:"price"`
	InStock int64   `json:"inStock"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "The Eminem Show", Artist: "Eminem", Price: 56.99, InStock: 140},
	{ID: "2", Title: "Get Rich or Die Tryin'", Artist: "50 Cent", Price: 42.99, InStock: 353},
	{ID: "3", Title: "All Eyez On Me", Artist: "Tupac Shakur", Price: 45.99, InStock: 234},
}

// getAlbums responds with the list of all albums as JSON
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the recived JSON to
	// newAlbum
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
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

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	router.Run("127.0.0.1:5000")
}
