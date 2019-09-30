package post

import (
	"encoding/json"
	"fmt"
	"github.com/yuedun/ginode/db"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	namebody := map[string]string{}
	name := c.Request.Body
	namebyte, _ := ioutil.ReadAll(name)
	json.Unmarshal(namebyte, &namebody)
	fmt.Println(namebody)
	c.JSON(200, gin.H{
		"message": namebody["name"],
	})
}

func GetPostInfo(c *gin.Context) {
	userService := NewService(db.Mysql)
	user, err := userService.GetPostInfo()
	if err != nil {
		fmt.Println("err:", err)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}

func GetPostInfoBySql(c *gin.Context) {
	userService := NewService(db.Mysql)
	user, err := userService.GetPostInfoBySql()
	if err != nil {
		fmt.Println("err:", err)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}

func CreatePost(c *gin.Context) {
	userService := NewService(db.Mysql)
	user := Post{}
	fmt.Println(">>>", c.PostForm("mobile"))
	user.Mobile = c.PostForm("mobile")
	user.CreatedAt = time.Now()
	err := userService.CreatePost(&user)
	if err != nil {
		fmt.Println("err:", err)
	}
	c.JSON(http.StatusOK, gin.H{
		"data":    user,
		"message": "ok",
	})
}

func UpdatePost(c *gin.Context) {
	userService := NewService(db.Mysql)
	user := Post{}
	userId, _ := strconv.Atoi(c.Param("id"))
	user.Addr = c.PostForm("addr")
	err := userService.UpdatePost(userId, &user)
	if err != nil {
		fmt.Println("err:", err)
	}
	c.JSON(http.StatusOK, gin.H{
		"data":    user,
		"message": "ok",
	})
}

func DeletePost(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))
	userService := NewService(db.Mysql)
	err := userService.DeletePost(userId)
	if err != nil {
		fmt.Println("err:", err)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
