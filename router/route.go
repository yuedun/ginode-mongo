package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yuedun/ginode/middleware"
	"github.com/yuedun/ginode/pkg/post"
	"github.com/yuedun/ginode/pkg/user"
)

/**
 * 路由注册
 */
func RouterRegister(router *gin.Engine) {
	//user路由注册
	userRouter := router.Group("/user")
	{
		userRouter.GET("/index", user.Index)
		userRouter.GET("/users/:id", middleware.Auth(), user.GetUserInfo)//单独给某个路由添加中间件
		userRouter.GET("/users-by-sql/:id", user.GetUserInfoBySql)
		userRouter.POST("/", user.CreateUser)
		userRouter.PUT("/update/:id", user.UpdateUser)
		userRouter.DELETE("/del/:id", user.DeleteUser)
	}
	//user路由注册
	postRouter := router.Group("/post")
	{
		postRouter.GET("/", post.Index)
		postRouter.GET("/posts/:id", middleware.Auth(), post.GetPostInfo)
		postRouter.GET("/posts-by-sql/:id", post.GetPostInfoBySql)
		postRouter.POST("/", post.CreatePost)
		postRouter.PUT("/:id", post.UpdatePost)
		postRouter.DELETE("/:id", post.DeletePost)
	}
}
