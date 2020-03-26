package middleware

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

//Logger 日志中间件
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// 设置 example 变量
		c.Set("example", "12345")

		// 设置cookie
		c.SetCookie("user_name", "yuedun", 3600, "/", "localhost", http.SameSiteDefaultMode, true, true)

		// 请求前

		c.Next()

		// 请求后
		latency := time.Since(t)
		log.Print("耗时：", latency)

		// 获取发送的 status
		status := c.Writer.Status()
		log.Println("状态：", status)
	}
}

// 权限校验
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("token")
		if err != nil {
			c.Abort() //不继续执行
			c.JSON(http.StatusForbidden, gin.H{
				"message": err.Error(),
			})
			return
		}
		log.Println(">>>", token)
		if token == "" {
			log.Print("权限验证未通过")
			c.Abort() //不继续执行
			c.JSON(http.StatusForbidden, gin.H{
				"message": "权限验证未通过",
			})
			return
		} else {
			c.Next() //如果通过中间件需要调用Next，使其继续执行下一个func
		}
	}
}
