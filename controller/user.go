package controller

import (
	"encoding/json"
	"fmt"
	"github.com/yuedun/ginode/model"
	"github.com/yuedun/ginode/service"
	"io/ioutil"
	"net/http"
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
	user, err := service.GetUserInfo()
	if err != nil {
		fmt.Println("err:", err)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}

func GetUserInfoBySql(c *gin.Context) {
	user, err := service.GetUserInfoBySql()
	if err != nil {
		fmt.Println("err:", err)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}

func CreateUser(c *gin.Context) {
	user := model.User{}
	fmt.Println(">>>", c.PostForm("mobile"))
	user.Mobile = c.PostForm("mobile")
	user.CreatedAt = time.Now()
	err := service.CreateUser(&user)
	if err != nil {
		fmt.Println("err:", err)
	}
	c.JSON(http.StatusOK, gin.H{
		"data":    user,
		"message": "ok",
	})
}

func UpdateUser(c *gin.Context) {
	user := model.User{}
	err := service.UpdateUser(user)
	if err != nil {
		fmt.Println("err:", err)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}

func DeleteUser(c *gin.Context) {
	user, err := service.DeleteUser()
	if err != nil {
		fmt.Println("err:", err)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}
