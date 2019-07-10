package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yuedun/ginode/middleware"
)

func main() {
	r := gin.Default()
	r.Use(middleware.Logger())

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello gin!!!!!",
		})
	})
	RouterRgister(r)
	r.Run() // listen and serve on 0.0.0.0:8080
}
