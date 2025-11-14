package backend

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListItems(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		items := []ListItem{}
		result := db.Find(&items)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		dtoList := make([]ListItemDto, len(items))
		for i, item := range items {
			dtoList[i] = ListItemDto{
				Id:    item.ID,
				Title: item.Title,
			}
		}

		c.JSON(http.StatusOK, dtoList)
	}
}

func GetTags(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		tags := []Tag{}
		result := db.Find(&tags)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		dtoList := make([]TagDto, len(tags))
		for i, item := range tags {
			dtoList[i] = TagDto{
				Id:   item.ID,
				Name: item.Name,
			}
		}

		c.JSON(http.StatusOK, dtoList)
	}
}

func AddItem(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var json struct {
			Title string `json:"title" binding:"required"`
		}

		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		item := ListItem{Title: json.Title}
		db.Create(&item)
		// Here you would typically add the item to the database
		// For demonstration, we just return the received item
		c.JSON(http.StatusCreated, gin.H{})
	}
}

func AddTag(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var json struct {
			Name string `json:"name" binding:"required"`
		}

		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		fmt.Println("Adding tag:", json.Name)

		item := Tag{Name: json.Name}
		db.Create(&item)
		// Here you would typically add the item to the database
		// For demonstration, we just return the received item
		c.JSON(http.StatusCreated, gin.H{})
	}
}
