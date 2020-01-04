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

//Index
func Index(c *gin.Context) {
	nameBody := map[string]string{}
	name := c.Request.Body
	nameByte, _ := ioutil.ReadAll(name)
	json.Unmarshal(nameByte, &nameBody)
	fmt.Println(nameBody)
	c.JSON(200, gin.H{
		"message": nameBody["name"],
	})
}

//GetUserInfo
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

//GetUserInfoBySql
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

//CreateUser
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

//UpdateUser post json
func UpdateUser(c *gin.Context) {
	userService := NewService(db.Mysql)
	var user User
	userID, _ := strconv.Atoi(c.Param("id"))
	//user.Addr = c.PostForm("addr")
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"data":    nil,
			"message": "err",
		})
	} else {
		err := userService.UpdateUser(userID, &user)
		if err != nil {
			fmt.Println("err:", err)
		}
		c.JSON(http.StatusOK, gin.H{
			"data":    user,
			"message": "ok",
		})
	}
}

//DeleteUser
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
