package middleware

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// Set example variable
		c.Set("example", "12345")

		// before request

		c.Next()

		// after request
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}

/**
 * 权限校验
 */
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Query("a") == "1" {
			log.Print("权限验证未通过")
			c.Abort() //不继续执行
			c.JSON(http.StatusForbidden, gin.H{
				"messge": "权限验证未通过",
			})
			return
		} else {
			c.Next() //如果通过中间件需要调用Next，使其继续执行下一个func
		}
	}
}
