package main

import (
	"log"

	"training/GolangAPI/controllers"
	"training/GolangAPI/storage"
	"training/GolangAPI/utils"

	"github.com/gin-gonic/gin"
)

// albums slice to seed record album data.
//var albums = []entity.Album{}

//{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
//{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
//{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
//}

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Start Project")
	//	config.RuntimeSetup, " and on port: ", config.AppPort)

	DB := storage.ConnectDB()
	defer DB.Close()

	log.Println(DB.RowsAffected)

	router := gin.Default()
	//	router.Static(relativePath:"/", root: "./static")
	router.GET("/albums", controllers.GetAlbums)
	router.POST("/albums", controllers.AddAlbums)
	router.GET("/albums/:id", controllers.GetAlbumByID)
	router.DELETE("/albums/:id", controllers.DeleteAlbum)
	router.PATCH("/albums/:id", controllers.UpdateAlbum)

	//router.Run("localhost:8080")
	router.Run(config.ServerAddress + ":" + config.AppPort)
}
