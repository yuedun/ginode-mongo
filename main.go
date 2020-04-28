package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/yuedun/ginode-mongo/db"
	"github.com/yuedun/ginode-mongo/router"
)

func main() {
	r := gin.Default()
	//r.Use(middleware.Logger())//全局中间件
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello ginode-mongo")
	})

	router.Register(r)
	r.Run(":8900") // listen and serve on 0.0.0.0:8080
}
