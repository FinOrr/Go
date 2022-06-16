package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// An album represents data about a record album
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// Albums slice to seed record album data //
// To keep things simple, all data is stored in memory rather than a DB
// The memory will be lost when the server is stopped, but will be
// recreated when the server is started again!
var albums = []album{
	{ID: "1", Title: "C.R.E.A.M.", Artist: "Wu-Tang Clan", Price: 14.99},
	{ID: "2", Title: "Grown Man Sport", Artist: "Pete Rock", Price: 9.95},
	{ID: "3", Title: "All Caps", Artist: "Madvillain", Price: 12.49},
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}

//getAlbums responds with the list of all albums as JSON
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

//postAlbums adds an album from JSON received in the request body
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to newAlbum, which
	// binds the request body to newAlbum
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album struct to the albums slice, and add a 201
	// status code to the response, along with the JSON of the album
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
