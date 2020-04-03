package user

import (
	"fmt"
	"github.com/yuedun/ginode-mongo/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

//GetUserInfo
func GetUserInfo(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.(error).Error(),
			})
		}
	}()
	userID := c.Param("id")
	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		panic(err)
	}
	username := c.Param("username")
	mobile := c.Param("mobile")
	userService := NewService(db.Mongodb)
	userObj := User{
		Id:       id,
		UserName: username,
		Mobile:   mobile,
	}
	user, err := userService.GetUserInfo(userObj)
	if err != nil {
		fmt.Println("err:", err)
	}
	c.JSON(http.StatusOK, gin.H{
		"data":    user,
		"message": "ok",
	})
}

//CreateUser
func CreateUser(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.(error).Error(),
			})
		}
	}()
	userService := NewService(db.Mongodb)
	user := User{}
	if err := c.ShouldBind(&user); err != nil {
		panic(err)
	}
	user.CreatedAt = time.Now()
	err := userService.CreateUser(&user)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"data":    user,
		"message": "ok",
	})
}

//UpdateUser post json
func UpdateUser(c *gin.Context) {
	userService := NewService(db.Mongodb)
	var user User
	userID, _ := strconv.Atoi(c.Param("id"))
	fmt.Println(userID)
	//user.Addr = c.PostForm("addr")
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"data":    nil,
			"message": "err",
		})
	} else {
		err := userService.UpdateUser(&user)
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
	userService := NewService(db.Mongodb)
	err := userService.DeleteUser(userID)
	if err != nil {
		fmt.Println("err:", err)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
