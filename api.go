package main

import (
	"go-rest-api/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/gin-contrib/cors"
)

func main() {
	dsn := "root:@tcp(localhost:3307)/go_api?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	r := gin.Default()
	r.GET("/users", func(c *gin.Context) {
		var users []model.User
		db.Find(&users)
		c.JSON(200, users)
	})

	r.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		var user model.User
		db.First(&user, id)
		c.JSON(200, user)
	})

	r.POST("/users", func(c *gin.Context) {
		var user model.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		result := db.Create(&user)
		c.JSON(200, gin.H{"RowAffected": result.RowsAffected})
	})

	r.DELETE("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		var user model.User
		db.First(&user, id)
		db.Delete(&user)
		c.JSON(200, user)
	})

	r.PUT("/users/:id", func(c *gin.Context) {
		var user model.User
		var updatedUser model.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.First(&updatedUser, user.ID)
		updatedUser.Username = user.Username
		updatedUser.Email = user.Email
		updatedUser.Password = user.Password
		db.Save(&updatedUser)
		c.JSON(200, updatedUser)
	})

	r.Use(cors.Default())
	r.Run() // listen and serve on port 8080
}
