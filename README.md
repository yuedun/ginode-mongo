# ginode
一个和nodejs很像的框架

## 本地开发
beego框架自带了本地调试工具，在修改代码后可以自动重启，幸运的是，`bee`工具同样可以在`gin`项目中使用
> bee run

但是`bee`不能使用在非`GOPATH`目录下

另外一种使用了`dogo`，本项目包含了dogo.json配置，可根据自己项目路径修改
> dogo

## 目录结构
本项目结构组织使用的是按职责划分，采用这种结构的原因是，目前流行的开发模式是微服务架构，但是一般项目都是由小到大再到拆分的过程，如果项目初始就使用微服务的架构开发的话估计还没等项目开发完公司就完蛋了。
所以最开始还是单体架构才是正确的方式，不过为了以后方便拆分，可以对项目目录进行合理的划分。
从路由入口看：
```go
func RouterRegister(router *gin.Engine) {
	//user路由注册
    	userRouter := router.Group("/user")
    	{
    		userRouter.GET("/index", user.Index)
    		userRouter.GET("/users/:id", middleware.Auth(), user.GetUserInfo)
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
```
本项目只是一个示例项目，模块较少，只开设了两个模块，`user`和`post`，注册路由的时候就可以按照不同的职责来分组注册，`pkg`目录下放置的就是不同职责的模块。
将来需要对服务进行拆分的时候只需要将`pkg`目录下的模块分离出去即可形成独立的服务，对依赖的的修改也较少。如果使用的MVC架构的话，要对不同职责的服务进行拆分则很困难，需要对每一层中对应的文件都拿出来重新组织。