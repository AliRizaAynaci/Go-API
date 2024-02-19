package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func GetHandler(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func PostHandler(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func GetByIdHandler(c *gin.Context) {
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

func UpdateHandler(c *gin.Context) {
	id := c.Param("id")

	var updatedAlbum album

	if err := c.BindJSON(&updatedAlbum); err != nil {
		return
	}

	for i, a := range albums {
		if a.ID == id {
			albums[i].Title = updatedAlbum.Title
			albums[i].Artist = updatedAlbum.Artist
			albums[i].Price = updatedAlbum.Price
			c.IndentedJSON(http.StatusOK, albums[i])
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})

}

func DeleteHandler(c *gin.Context) {
	id := c.Param("id")

	for i, a := range albums {
		if a.ID == id {
			albums = append(albums[:i], albums[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "album deleted!"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
