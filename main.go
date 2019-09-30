package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/yuedun/ginode/middleware"
	_ "github.com/yuedun/ginode/db"
	"github.com/yuedun/ginode/router"
)

func main() {
	r := gin.Default()
	r.Use(middleware.Logger())

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello gin!!!!!",
			"code":    0,
		})
	})

	router.RouterRegister(r)
	r.Run() // listen and serve on 0.0.0.0:8080
}
