package main

import (
	"go_server/config"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func GetUsers(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	} else {
		var user User
		result := config.Cfg.GormDB.Where("id = ?", id).First(&user)
		if result.Error != nil {
			println("Error fetching user: ", result.Error.Error())
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"user": user})
	}
}

func main() {
	r := gin.Default()
	r.GET("/users/:id", GetUsers)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
