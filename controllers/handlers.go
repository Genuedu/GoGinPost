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
	if err := storage.DB.Find(&albums).Error; err != nil {
		fmt.Printf("error select Albums: %3v \n", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	} else {
		c.IndentedJSON(http.StatusOK, albums)
	}
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
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Album OK",
		})
	}
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func GetAlbumByID(c *gin.Context) {
	var albums = []entity.Album{}

	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.

	if err := storage.DB.First(&albums, "id = ?", id).Error; err != nil {
		fmt.Printf("error select Album: %s \n", id)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	} else {
		c.IndentedJSON(http.StatusOK, albums)
	}
}

func DeleteAlbum(c *gin.Context) {
	var albums = []entity.Album{}

	id := c.Param("id")

	// Loop over the list of albums and delete if found

	if err := storage.DB.Delete(&albums, "id = ?", id).Error; err != nil {
		fmt.Printf("error select Album: %s \n", id)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
		return
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{
			"message": "Album deleted succesfully",
		})
	}
}

func UpdateAlbum(c *gin.Context) {

	var album entity.Album

	id := c.Param("id")

	// Loop over the list of albums, grab a single album and make sure it exists

	if err := storage.DB.Where("id = ?", id).First(&album).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// we need to validate the user input with the UpdateAlbumInput schema
	var input entity.UpdateAlbumInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// we update the album model using the Updates method
	storage.DB.Model(&album).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": album})
}
