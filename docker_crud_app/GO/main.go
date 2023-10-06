package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	initRoutes(r)
	r.Run(":8080")
}

type Storage struct {
	ID   int        `json:"id"`
	Time *time.Time `json:"time"`
	Text string     `json:"text"`
}

type POSTInput struct {
	Text string `json:"text"`
}

type DELETEdata struct {
	ID int `json:"id"`
}

type PUTdata struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

func initRoutes(r *gin.Engine) {
	var store []Storage
	ID := 0
	r.GET("/get", func(c *gin.Context) {
		c.JSON(http.StatusOK, store)
	})
	r.POST("/post", func(c *gin.Context) {
		var data POSTInput
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload in POST"})
			return
		}
		ID++
		currentTime := time.Now()
		dataPacket := Storage{ID, &currentTime, data.Text}

		store = append(store, dataPacket)
		c.JSON(http.StatusOK, gin.H{"status": "added"})
	})
	r.PUT("/put", func(c *gin.Context) {
		var data PUTdata
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload in PUT"})
			return
		}
		for i := 0; i < len(store); i++ {
			if store[i].ID == data.ID {
				store[i].Text = data.Text
				break
			}
		}
		c.JSON(http.StatusOK, gin.H{"status": "updated"})
	})
	r.DELETE("/delete", func(c *gin.Context) {
		var data DELETEdata
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload in DELETE"})
			return
		}
		for i := 0; i < len(store); i++ {
			if store[i].ID == data.ID {
				store = append(store[:i], store[i+1:]...)
				break
			}
		}
		c.JSON(http.StatusOK, gin.H{"status": "deleted"})
	})
}
