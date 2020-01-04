package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/yuedun/ginode/db"
	"github.com/yuedun/ginode/middleware"
	"github.com/yuedun/ginode/router"
	"net/http"
)

func main() {
	r := gin.Default()
	r.Use(middleware.Logger())//全局中间件
	r.LoadHTMLGlob("templates/*")//加载模板
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tpl", gin.H{
			"title": "Hello World!",
		})
	})

	router.RouterRegister(r)
	r.Run(":8900") // listen and serve on 0.0.0.0:8080
}
