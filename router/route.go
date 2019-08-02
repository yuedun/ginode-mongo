package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yuedun/ginode/controller"
	"github.com/yuedun/ginode/middleware"
)

/**
 * 路由注册
 */
func RouterRegister(router *gin.Engine) {
	router.POST("/index", controller.Index)
	router.GET("/get-user-info", middleware.Auth(), controller.GetUserInfo)
	router.GET("/get-user-info-by-sql", controller.GetUserInfoBySql)
}

