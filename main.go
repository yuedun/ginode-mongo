package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/yuedun/ginode/controller"
	"github.com/yuedun/ginode/middleware"
	_ "github.com/yuedun/ginode/model"
)

func main() {
	r := gin.Default()
	r.Use(middleware.Logger())

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello gin!!!!!",
		})
	})

	routerRegister(r)
	r.Run() // listen and serve on 0.0.0.0:8080
}

/**
 * 路由注册
 */
func routerRegister(router *gin.Engine) {
	router.POST("/index", controller.Index)
	router.GET("/get-user-info", middleware.Auth(), controller.GetUserInfo)
}
