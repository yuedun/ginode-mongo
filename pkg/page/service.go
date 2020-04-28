package page

import (
	"context"
	"fmt"

	"github.com/yuedun/ginode-mongo/pkg/component"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	PageService interface {
		GetPageList(offset, limit int64, search Page) (page []Page, count int64, err error)
		GetPage(url string) (page Page, err error)
		CreatePage(page *Page) (err error)
		UpdatePage(page *Page) (err error)
		DeletePage(pageID primitive.ObjectID) (err error)
		GetPageComponents(pageID primitive.ObjectID) (components []component.Component, err error)
		UpdatePageComponents(id primitive.ObjectID, pageComponents []component.Component) error
		CopyPage(id primitive.ObjectID, url string) error
	}
)
type pageService struct {
	mongo *mongo.Database
}

func NewService(mongo *mongo.Database) PageService {
	return &pageService{
		mongo: mongo,
	}
}

func (this *pageService) GetPageList(offset, limit int64, search Page) (pages []Page, count int64, err error) {
	var cursor *mongo.Cursor
	if cursor, err = this.mongo.Collection("page").Find(
		context.Background(),
		bson.M{"website_id": search.WebsiteID, "status": 1}, //没有条件必须为空，不能包含键值对，go中对象会是零值作为查询，所以条件只能动态填充
		options.Find().SetLimit(limit),
		options.Find().SetSkip(offset),
		options.Find().SetSort(bson.M{"_id": -1})); err != nil {
		return nil, 0, err
	}
	defer cursor.Close(context.Background())
	page := Page{}
	for cursor.Next(context.Background()) {
		if err = cursor.Decode(&page); err != nil {
			return nil, 0, err
		}
		if page.Components == nil {
			page.Components = []component.Component{}
		}
		pages = append(pages, page)
	}
	fmt.Printf("数据：%v\n", pages)
	//查询集合里面有多少数据
	if count, err = this.mongo.Collection("page").CountDocuments(context.Background(), bson.M{"website_id": search.WebsiteID, "status": 1}); err != nil {
		return nil, 0, err
	}

	fmt.Printf("Count里面有多少条数据:%d\n", count)
	if count == 0 {
		pages = make([]Page, 0)
	}
	return pages, count, err
}

func (this *pageService) GetPage(url string) (page Page, err error) {
	//没有条件必须为空，不能包含键值对，go中对象会是零值作为查询，所以条件只能动态填充
	if err = this.mongo.Collection("page").FindOne(context.Background(), bson.M{"url": url}).Decode(&page); err != nil {
		fmt.Println("get page err:", err.Error())
		return Page{}, err
	}
	fmt.Printf("数据:%+v\n", page)
	return page, nil
}

func (this *pageService) CreatePage(page *Page) (err error) {
	result, err := this.mongo.Collection("page").InsertOne(context.Background(), page)
	if err != nil {
		return err
	}
	page.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (this *pageService) UpdatePage(page *Page) (err error) {
	result, err := this.mongo.Collection("page").UpdateOne(
		context.Background(),
		bson.D{{"_id", page.ID}},
		bson.M{
			"$set": bson.M{
				"name":        page.Name,
				"components":  page.Components,
				"url":         page.URL,
				"keywords":    page.Keywords,
				"description": page.Description,
				"status":      page.Status,
			},
		})
	fmt.Println(result.MatchedCount, result.ModifiedCount)
	if err != nil {
		return err
	}
	return nil
}

func (this *pageService) DeletePage(pageID primitive.ObjectID) (err error) {
	result, err := this.mongo.Collection("page").UpdateOne(context.Background(),
		bson.M{"_id": pageID},
		bson.M{
			"$set": bson.M{"status": 0},
		})
	if err != nil {
		return err
	}
	fmt.Println(result.ModifiedCount)
	return nil
}

func (this *pageService) GetPageComponents(pageID primitive.ObjectID) (components []component.Component, err error) {
	page := Page{}
	err = this.mongo.Collection("page").FindOne(context.Background(), bson.M{"_id": pageID}).Decode(&page)
	if err != nil {
		return nil, err
	}
	fmt.Println(page)
	components = page.Components
	return components, nil
}

func (this *pageService) UpdatePageComponents(id primitive.ObjectID, pageComponents []component.Component) error {
	result, err := this.mongo.Collection("page").UpdateOne(
		context.Background(),
		bson.D{{"_id", id}},
		bson.M{
			"$set": bson.M{
				"components": pageComponents,
			},
		})
	fmt.Println(result.MatchedCount, result.ModifiedCount)
	if err != nil {
		return err
	}
	return nil
}

func (this *pageService) CopyPage(id primitive.ObjectID, url string) error {
	fmt.Println(id, url)
	page := Page{}
	err := this.mongo.Collection("page").FindOne(
		context.Background(),
		bson.D{{"_id", id}}).Decode(&page)
	if err != nil {
		fmt.Println(err)
		return err
	}
	page.ID = primitive.NewObjectID()
	page.URL = url
	result, err := this.mongo.Collection("page").InsertOne(
		context.Background(),
		page)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(result.InsertedID)
	return nil
}
