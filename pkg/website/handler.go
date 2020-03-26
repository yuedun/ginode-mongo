package website

import (
	"fmt"
	"github.com/yuedun/ginode/db"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

//WebsiteList列表
func WebsiteList(c *gin.Context) {
	offset, _ := strconv.Atoi(c.Query("offset"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	wbService := NewService(db.Mysql)
	webs, err := wbService.GetWebsiteList(offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data":    webs,
	})
}

//Create
func Create(c *gin.Context) {
	websiteService := NewService(db.Mysql)
	wbObj := Website{}
	c.BindJSON(&wbObj)
	wbObj.CreatedAt = time.Now()
	wbObj.Status = 1
	err := websiteService.CreateWebsite(&wbObj)
	if err != nil {
		fmt.Println("err:", err)
	}
	c.JSON(http.StatusOK, gin.H{
		"data":    wbObj,
		"message": "ok",
	})
}

//Update
func Update(c *gin.Context) {
	websiteService := NewService(db.Mysql)
	website := Website{}
	c.BindJSON(&website)
	err := websiteService.UpdateWebsite(&website)
	if err != nil {
		fmt.Println("err:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data":    website,
			"message": "ok",
		})
	}
}

//Delete
func Delete(c *gin.Context) {
	websiteId, _ := strconv.Atoi(c.Param("id"))
	websiteService := NewService(db.Mysql)
	err := websiteService.DeleteWebsite(websiteId)
	if err != nil {
		fmt.Println("err:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	}
}
