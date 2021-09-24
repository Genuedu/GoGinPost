package controllers

import (
	"fmt"
	"net/http"
	"training/GolangAPI/entity"
	"training/GolangAPI/storage"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetAlbums(c *gin.Context) {
	var albums = []entity.Album{}

	c.IndentedJSON(http.StatusOK, albums)
}

func AddAlbums(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", "*")

	newAlbum := entity.Album{}
	c.BindJSON(&newAlbum)
	newAlbum.ID = uuid.NewString()

	if err := storage.DB.Create(&newAlbum).Error; err != nil {
		fmt.Printf("error add Album: %3v \n", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	/*if err != nil {
		fmt.Printf("error add Album: %3v \n", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Album OK",
		})
	}*/
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
//func getAlbumByID(c *gin.Context) {
//	id := c.Param("id")
//
//
// Loop over the list of albums, looking for
// an album whose ID value matches the parameter.
//	for _, a := range albums {
//		if fmt.Sprint(a.ID) == id {
//			c.IndentedJSON(http.StatusOK, a)
//			return
//		}
//	}
//	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
//}
