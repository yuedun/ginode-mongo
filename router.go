package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yuedun/ginode/controller"
	"github.com/yuedun/ginode/middleware"
)

/**
 * 路由注册
 */
func RouterRgister(router *gin.Engine) {
	router.POST("/index", controller.Index)
	router.POST("/get-user-info", middleware.Auth(), controller.GetUserInfo)
}
