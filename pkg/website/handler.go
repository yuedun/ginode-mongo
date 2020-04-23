package website

import (
	"fmt"
	jwt "github.com/appleboy/gin-jwt/v2"
	"net/http"
	"strconv"
	"time"

	"github.com/yuedun/ginode-mongo/db"
	"github.com/yuedun/ginode-mongo/pkg/component"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gin-gonic/gin"
)

type any = interface{}

//WebsiteList列表
func WebsiteList(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.(error).Error(),
			})
		}
	}()
	claims := jwt.ExtractClaims(c)
	userID, err := primitive.ObjectIDFromHex(claims["user_id"].(string))
	if err != nil {
		panic(err)
	}
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
		UserID:   userID,
		Name:     name,
		Category: category,
		Status:   1,
	}
	wbService := NewService(db.NewDB("website"))
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
func GetWebsite(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.(error).Error(),
			})
		}
	}()
	claims := jwt.ExtractClaims(c)
	userID, err := primitive.ObjectIDFromHex(claims["user_id"].(string))
	if err != nil {
		panic(err)
	}
	webService := NewService(db.NewDB("website"))
	name := c.Query("name")
	fmt.Println("url:", name)
	website, err := webService.GetWebsite(userID, name)
	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{
			"err": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": website,
	})

}

//Create
func Create(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.(error).Error(),
			})
		}
	}()
	claims := jwt.ExtractClaims(c)
	userID, err := primitive.ObjectIDFromHex(claims["user_id"].(string))
	if err != nil {
		panic(err)
	}
	websiteService := NewService(db.NewDB("website"))
	wbObj := Website{}
	c.ShouldBind(&wbObj)
	wbObj.ID = primitive.NewObjectID()
	wbObj.UserID = userID
	wbObj.CreatedAt = time.Now()
	wbObj.UpdatedAt = time.Now()
	wbObj.Status = 1
	fmt.Println(wbObj)
	err = websiteService.CreateWebsite(&wbObj)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"data":    wbObj,
		"message": "ok",
	})
}

//Update
func Update(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.(error).Error(),
			})
		}
	}()
	claims := jwt.ExtractClaims(c)
	userID, err := primitive.ObjectIDFromHex(claims["user_id"].(string))
	if err != nil {
		panic(err)
	}
	websiteService := NewService(db.NewDB("website"))
	website := Website{}
	c.ShouldBind(&website)
	website.UserID = userID
	fmt.Println(website)
	err = websiteService.UpdateWebsite(&website)
	if err != nil {
		panic(err)
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
	claims := jwt.ExtractClaims(c)
	userID, err := primitive.ObjectIDFromHex(claims["user_id"].(string))
	if err != nil {
		panic(err)
	}
	websiteId := c.Param("id")
	id, err := primitive.ObjectIDFromHex(websiteId)
	if err != nil {
		panic(err)
	}
	websiteService := NewService(db.NewDB("website"))
	err = websiteService.DeleteWebsite(userID, id)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})

}

//GetWebsiteComponents
func GetWebsiteComponents(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.(error).Error(),
			})
		}
	}()
	claims := jwt.ExtractClaims(c)
	userID, err := primitive.ObjectIDFromHex(claims["user_id"].(string))
	if err != nil {
		panic(err)
	}
	websiteId := c.Param("id")
	id, err := primitive.ObjectIDFromHex(websiteId)
	if err != nil {
		panic(err)
	}
	websiteService := NewService(db.NewDB("website"))
	components, err := websiteService.GetWebsiteComponents(userID, id)
	if err != nil {
		panic(err)
	}
	fmt.Println("components", components)
	c.JSON(http.StatusOK, gin.H{
		"data":    components,
		"message": "ok",
	})

}

// UpdateWebsiteComponents 单独修改网站组件
func UpdateWebsiteComponents(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.(error).Error(),
			})
		}
	}()
	claims := jwt.ExtractClaims(c)
	userID, err := primitive.ObjectIDFromHex(claims["user_id"].(string))
	if err != nil {
		panic(err)
	}
	websiteId := c.Param("id")
	id, err := primitive.ObjectIDFromHex(websiteId)
	var websiteComponents []component.Component
	c.ShouldBind(&websiteComponents)
	fmt.Println(websiteComponents)
	websiteService := NewService(db.NewDB("website"))
	err = websiteService.UpdateWebsiteComponents(userID, id, websiteComponents)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

// CopyPage 复制网站（页面），多个页面组成一个网站
func CopyPage(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.(error).Error(),
			})
		}
	}()
	claims := jwt.ExtractClaims(c)
	userID, err := primitive.ObjectIDFromHex(claims["user_id"].(string))
	if err != nil {
		panic(err)
	}
	websiteId := c.Param("id")
	id, err := primitive.ObjectIDFromHex(websiteId)
	if err != nil {
		panic(err)
	}
	url := c.Param("url")
	websiteService := NewService(db.NewDB("website"))
	err = websiteService.CopyPage(userID, id, url)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
