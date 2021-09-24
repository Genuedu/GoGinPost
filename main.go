package main

import (
	"log"

	"training/GolangAPI/controllers"
	"training/GolangAPI/storage"

	"github.com/gin-gonic/gin"
)

// albums slice to seed record album data.
//var albums = []entity.Album{}

//{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
//{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
//{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
//}

func main() {

	log.Println("Start Project")

	DB := storage.ConnectDB()
	defer DB.Close()

	log.Println(DB.RowsAffected)

	router := gin.Default()
	router.GET("/albums", controllers.GetAlbums)
	router.POST("/albums", controllers.AddAlbums)
	//	router.GET("/albums/:id", getAlbumByID)

	router.Run("localhost:8080")

}
