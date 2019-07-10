package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

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

func GetUserInfo(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"message": "GetUserInfo",
	})
}