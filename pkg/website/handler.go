package website

import (
	"fmt"
	"github.com/yuedun/ginode-mongo/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type any = interface{}

//WebsiteList列表
func WebsiteList(c *gin.Context) {
	offset, err := strconv.ParseInt(c.Query("offset"), 10, 64)
	if err != nil {
		offset = 0
	}
	limit, err := strconv.ParseInt(c.Query("limit"), 10, 64)
	if err != nil {
		limit = 10
	}
	name := c.Query("name")
	category := c.Query("category")
	websiteSearch := Website{
		Name:     name,
		Category: category,
		Status:   1,
	}
	wbService := NewService(db.Mongodb)
	list, total, err := wbService.GetWebsiteList(offset, limit, websiteSearch)
	data := map[string]any{
		"result": list,
		"count":  total,
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data":    data,
	})
}

// 获取单个网站
func GetWebsite(c *gin.Context)  {
	webService:=NewService(db.Mongodb)
	website, err:=webService.GetWebsite("xes")
	if err!=nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err":err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"data":website,
	})
}

//Create
func Create(c *gin.Context) {
	websiteService := NewService(db.Mongodb)
	wbObj := Website{}
	c.ShouldBind(&wbObj)
	wbObj.ID = primitive.NewObjectID()
	wbObj.CreatedAt = time.Now()
	wbObj.UpdatedAt = time.Now()
	wbObj.Status = 1
	fmt.Println(wbObj)
	err := websiteService.CreateWebsite(&wbObj)
	if err != nil {
		fmt.Println("err:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"data":    wbObj,
			"message": "ok",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data":    wbObj,
		"message": "ok",
	})
}

//Update
func Update(c *gin.Context) {
	websiteService := NewService(db.Mongodb)
	website := Website{}
	c.ShouldBind(&website)
	err := websiteService.UpdateWebsite(&website)
	if err != nil {
		fmt.Println("err:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data":    website,
			"message": "ok",
		})
	}
}

//Delete
func Delete(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.(error).Error(),
			})
		}
	}()
	websiteId := c.Param("id")
	id, err := primitive.ObjectIDFromHex(websiteId)
	if err != nil {
		panic(err)
	}
	websiteService := NewService(db.Mongodb)
	err = websiteService.DeleteWebsite(id)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})

}
