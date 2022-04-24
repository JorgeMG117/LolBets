package routes

import (
	"net/http"

	"github.com/JorgeMG117/LolBets/backend/models"
	"github.com/gin-gonic/gin"
)

func GetGames(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, models.GetGames())
}

// postAlbums adds an album from JSON received in the request body.
func PostGames(c *gin.Context) {
	var newGame models.Game

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newGame); err != nil {
		return
	}

	// Add the new album to the slice.
	models.AddGame(newGame)

	c.IndentedJSON(http.StatusCreated, newGame)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func GetAlbumByLeague(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range models.GetGames() {
		if a.Id == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
