package main

import (
	"github.com/JorgeMG117/LolBets/backend/configs"
	"github.com/JorgeMG117/LolBets/backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", routes.GetGames)
	router.GET("/:league", routes.GetAlbumByLeague)
	//router.GET("/games/:team", routes.GetAlbumByLeague)
	router.POST("/", routes.PostGames)

	/*v1 := router.Group("/games")
	{
		v1.POST("/login", loginEndpoint)
		v1.POST("/submit", submitEndpoint)
		v1.POST("/read", readEndpoint)
	}*/

	configs.ConnectDB()

	router.Run("localhost:8080")
}
