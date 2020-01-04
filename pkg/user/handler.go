package user

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

func GetUserInfo(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Param("id"))
	userService := NewUserService(db.Mysql)
	user, err := userService.GetUserInfo(userID)
	if err != nil {
		fmt.Println("err:", err)
	}
	c.JSON(http.StatusOK, gin.H{
		"data":    user,
		"message": "ok",
	})
}

func GetUserInfoBySql(c *gin.Context) {
	userService := NewService(db.Mysql)
	user, err := userService.GetUserInfoBySQL()
	if err != nil {
		fmt.Println("err:", err)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}

func CreateUser(c *gin.Context) {
	userService := NewService(db.Mysql)
	user := User{}
	fmt.Println(">>>", c.PostForm("mobile"))
	user.Mobile = c.PostForm("mobile")
	user.CreatedAt = time.Now()
	err := userService.CreateUser(&user)
	if err != nil {
		fmt.Println("err:", err)
	}
	c.JSON(http.StatusOK, gin.H{
		"data":    user,
		"message": "ok",
	})
}

func UpdateUser(c *gin.Context) {
	userService := NewService(db.Mysql)
	user := User{}
	userID, _ := strconv.Atoi(c.Param("id"))
	user.Addr = c.PostForm("addr")
	err := userService.UpdateUser(userID, &user)
	if err != nil {
		fmt.Println("err:", err)
	}
	c.JSON(http.StatusOK, gin.H{
		"data":    user,
		"message": "ok",
	})
}

func DeleteUser(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Param("id"))
	userService := NewService(db.Mysql)
	err := userService.DeleteUser(userID)
	if err != nil {
		fmt.Println("err:", err)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
