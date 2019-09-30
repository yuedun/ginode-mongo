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
	router.GET("/index", user.Index)
	router.GET("/users/:id", middleware.Auth(), user.GetUserInfo)
	router.GET("/users-by-sql/:id", user.GetUserInfoBySql)
	router.POST("/users/", user.CreateUser)
	router.PUT("/users/:id", user.UpdateUser)
	router.DELETE("/users/:id", user.DeleteUser)

	//user路由注册
	router.GET("/posts", post.Index)
	router.GET("/posts/:id", middleware.Auth(), post.GetPostInfo)
	router.GET("/posts-by-sql/:id", post.GetPostInfoBySql)
	router.POST("/posts/", post.CreatePost)
	router.PUT("/posts/:id", post.UpdatePost)
	router.DELETE("/posts/:id", post.DeletePost)
}
