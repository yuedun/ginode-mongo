package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yuedun/ginode-mongo/middleware"
	"github.com/yuedun/ginode-mongo/pkg/component"
	"github.com/yuedun/ginode-mongo/pkg/user"
	"github.com/yuedun/ginode-mongo/pkg/website"
)

/**
 * 路由注册
 */
func Register(router *gin.Engine) {
	userRouter := router.Group("/user")
	//user路由注册,可以给各个group加中间件
	userRouter.Use(middleware.Logger())
	{
		//userRouter.POST("/login", user.Login)
		userRouter.POST("/login", middleware.Jwt().LoginHandler)
		userRouter.GET("/refresh_token", middleware.Jwt().RefreshHandler) // 刷新token
		userRouter.GET("/logout", middleware.Jwt().LogoutHandler)
		userRouter.GET("/info/:id", user.GetUserInfo) //单独给某个路由添加中间件
		userRouter.POST("/", user.CreateUser)
		userRouter.GET("/list", user.GetUserList)
		userRouter.PUT("/update/:id", user.UpdateUser)
		userRouter.DELETE("/del/:id", user.DeleteUser)
	}
	//website路由注册
	websiteRouter := router.Group("/website")
	websiteRouter.Use(middleware.Jwt().MiddlewareFunc())
	{
		websiteRouter.GET("/", website.WebsiteList)
		websiteRouter.POST("/create", website.Create)
		websiteRouter.PUT("/update", website.Update)
		websiteRouter.DELETE("/delete/:id", website.Delete)
		websiteRouter.GET("/getWebsiteComponents/:id", website.GetWebsiteComponents)
		websiteRouter.PUT("/updateWebsiteComponents/:id", website.UpdateWebsiteComponents)
	}
	//component路由注册
	componentRouter := router.Group("/component")
	componentRouter.Use(middleware.Jwt().MiddlewareFunc())
	{
		componentRouter.GET("/", component.ComponentList)
		componentRouter.GET("/getComponent/:id", component.GetComponent)
		componentRouter.POST("/create", component.Create)
		componentRouter.PUT("/update", component.Update)
		componentRouter.DELETE("/delete/:id", component.Delete)
	}
	//website路由注册
	websiteAPI := router.Group("/api/website")
	{
		websiteAPI.GET("/", website.WebsiteList)
		websiteAPI.GET("/get-website", website.GetWebsite)
		websiteAPI.POST("/create", website.Create)
		websiteAPI.PUT("/update", website.Update)
		websiteAPI.DELETE("/delete/:id", website.Delete)
	}
}
