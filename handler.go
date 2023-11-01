package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var queue = New()

// POST /add
func add(c *gin.Context) {
	var req map[string]interface{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	queue.Add(&Item{body: req})
	c.JSON(http.StatusOK, gin.H{
		"message": "added to queue",
	})
}

// GET /get
func get(c *gin.Context) {
	item := queue.Get()
	if item == nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "no items in queue",
		})
		return
	}
	c.JSON(http.StatusOK, item.body)
}
