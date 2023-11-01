package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/add", add)
	r.GET("/get", get)

	r.Run()
}
