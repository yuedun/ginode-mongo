package website

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
	WebsiteService interface {
		GetWebsiteList(offset, limit int64, search Website) (website []Website, count int64, err error)
		GetWebsite(userID primitive.ObjectID, url string) (website Website, err error)
		CreateWebsite(website *Website) (err error)
		UpdateWebsite(website *Website) (err error)
		DeleteWebsite(userID primitive.ObjectID, websiteID primitive.ObjectID) (err error)
		GetWebsiteComponents(userID primitive.ObjectID, websiteID primitive.ObjectID) (components []component.Component, err error)
		UpdateWebsiteComponents(userID primitive.ObjectID, id primitive.ObjectID, websiteComponents []component.Component) error
		CopyPage(userID primitive.ObjectID, id primitive.ObjectID, url string) error
	}
)
type websiteService struct {
	mongo *mongo.Database
}

func NewService(mongo *mongo.Database) WebsiteService {
	return &websiteService{
		mongo: mongo,
	}
}

func (this *websiteService) GetWebsiteList(offset, limit int64, search Website) (websites []Website, count int64, err error) {
	var cursor *mongo.Cursor
	if cursor, err = this.mongo.Collection("website").Find(
		context.Background(),
		bson.M{}, //没有条件必须为空，不能包含键值对，go中对象会是零值作为查询，所以条件只能动态填充
		options.Find().SetLimit(limit),
		options.Find().SetSkip(offset),
		options.Find().SetSort(bson.M{"_id": -1})); err != nil {
		return nil, 0, err
	}
	defer cursor.Close(context.Background())
	website := Website{}
	for cursor.Next(context.Background()) {
		if err = cursor.Decode(&website); err != nil {
			return nil, 0, err
		}
		websites = append(websites, website)
	}
	fmt.Printf("数据：%v\n", websites)
	//查询集合里面有多少数据
	if count, err = this.mongo.Collection("website").CountDocuments(context.Background(), bson.M{}); err != nil {
		return nil, 0, err
	}

	fmt.Printf("Count里面有多少条数据:%d\n", count)
	return websites, count, err
}

func (this *websiteService) GetWebsite(userID primitive.ObjectID, url string) (website Website, err error) {
	//没有条件必须为空，不能包含键值对，go中对象会是零值作为查询，所以条件只能动态填充
	if err = this.mongo.Collection("website").FindOne(context.Background(), bson.M{"url": url}).Decode(&website); err != nil {
		fmt.Println("get website err:", err.Error())
		return Website{}, err
	}
	fmt.Printf("数据:%+v\n", website)
	return website, nil
}

func (this *websiteService) CreateWebsite(website *Website) (err error) {
	_, err = this.mongo.Collection("website").InsertOne(context.Background(), website)
	if err != nil {
		return err
	}
	return nil
}

func (this *websiteService) UpdateWebsite(website *Website) (err error) {
	result, err := this.mongo.Collection("website").UpdateOne(
		context.Background(),
		bson.D{{"_id", website.ID}},
		bson.M{
			"$set": bson.M{
				"name":        website.Name,
				"category":    website.Category,
				"url":         website.URL,
				"icon":        website.Icon,
				"keywords":    website.Keywords,
				"description": website.Description,
				"status":      website.Status,
			},
		})
	fmt.Println(result.MatchedCount, result.ModifiedCount)
	if err != nil {
		return err
	}
	return nil
}

func (this *websiteService) DeleteWebsite(userID, websiteID primitive.ObjectID) (err error) {
	result, err := this.mongo.Collection("website").DeleteOne(context.Background(), bson.M{"_id": websiteID})
	if err != nil {
		return err
	}
	fmt.Println(result.DeletedCount)
	return nil
}

func (this *websiteService) GetWebsiteComponents(userID, websiteID primitive.ObjectID) (components []component.Component, err error) {
	website := Website{}
	err = this.mongo.Collection("website").FindOne(context.Background(), bson.M{"_id": websiteID}).Decode(&website)
	if err != nil {
		return nil, err
	}
	fmt.Println(website)
	return components, nil
}

func (this *websiteService) UpdateWebsiteComponents(userID, id primitive.ObjectID, websiteComponents []component.Component) error {
	result, err := this.mongo.Collection("website").UpdateOne(
		context.Background(),
		bson.D{{"_id", id}},
		bson.M{
			"$set": bson.M{
				"components": websiteComponents,
			},
		})
	fmt.Println(result.MatchedCount, result.ModifiedCount)
	if err != nil {
		return err
	}
	return nil
}

func (this *websiteService) CopyPage(userID, id primitive.ObjectID, url string) error {
	fmt.Println(id, url)
	website := Website{}
	err := this.mongo.Collection("website").FindOne(
		context.Background(),
		bson.D{{"_id", id}}).Decode(&website)
	if err != nil {
		fmt.Println(err)
		return err
	}
	website.ID = primitive.NewObjectID()
	website.URL = url
	result, err := this.mongo.Collection("website").InsertOne(
		context.Background(),
		website)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(result.InsertedID)
	return nil
}
