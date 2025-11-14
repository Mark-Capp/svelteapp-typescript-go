package backend

import (
	"log"
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
			log.Println("got asdfadf", item)
			dtoList[i] = ListItemDto{
				Id:    item.ID,
				Title: item.Title,
			}
		}

		c.JSON(http.StatusOK, dtoList)
	}
}

func Tags(c *gin.Context) {
	// return JSON
	c.JSON(http.StatusOK, []string{"tag1", "tag2", "tag3", "tag4"})
}

func AddItem(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var json struct {
			Name string `json:"name" binding:"required"`
		}

		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		item := ListItem{Title: json.Name}
		db.Create(&item)

		log.Println("added", item)

		// Here you would typically add the item to the database
		// For demonstration, we just return the received item
		c.JSON(http.StatusCreated, gin.H{})
	}
}
