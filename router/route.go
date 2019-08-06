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
	router.GET("/index", controller.Index)
	router.GET("/users/:id", middleware.Auth(), controller.GetUserInfo)
	router.GET("/users-by-sql/:id", controller.GetUserInfoBySql)
	router.POST("/users/", controller.CreateUser)
	router.PUT("/users/:id", controller.UpdateUser)
	router.DELETE("/users/:id", controller.DeleteUser)
}
