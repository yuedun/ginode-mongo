package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"encoding/json"
	"io/ioutil"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello gin!!!!!",
		})
	})
	r.POST("/index", func(c *gin.Context) {
		namebody:=map[string]string{}
		name:=c.Request.Body
		namebyte,_:=ioutil.ReadAll(name)
		json.Unmarshal(namebyte, &namebody)
		fmt.Println(namebody)
		c.JSON(200, gin.H{
			"message": namebody["name"],
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
