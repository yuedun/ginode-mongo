package middleware

import (
	jwt "github.com/appleboy/gin-jwt/v2"
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
		c.SetCookie("user_name", "yuedun", 3600, "/", "localhost", true, true)

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

type User struct {
	UserName  string
	FirstName string
	LastName  string
}
type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func Jwt() *jwt.GinJWTMiddleware {
	var identityKey = "id"
	// the jwt middleware
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*User); ok {
				return jwt.MapClaims{
					identityKey: v.UserName,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &User{
				UserName: claims[identityKey].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals login
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			userID := loginVals.Username
			password := loginVals.Password
			log.Println(">>>>>>>>>", userID, password)
			if (userID == "admin" && password == "admin") || (userID == "test" && password == "test") {
				return &User{
					UserName:  userID,
					LastName:  "Bo-Yi",
					FirstName: "Wu",
				}, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(*User); ok && v.UserName == "test" {
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// 获取jwt token的方法，从header中获取，从query中获取，从cookie中获取
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}
	return authMiddleware
}
