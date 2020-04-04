package website

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	WebsiteService interface {
		GetWebsiteList(offset, limit int64, search Website) (website []Website, count int64, err error)
		GetWebsite(url string) (website Website, err error)
		CreateWebsite(website *Website) (err error)
		UpdateWebsite(website *Website) (err error)
		DeleteWebsite(websiteID primitive.ObjectID) (err error)
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
		fmt.Println(website)
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

func (this *websiteService) GetWebsite(url string) (website Website, err error) {
	//没有条件必须为空，不能包含键值对，go中对象会是零值作为查询，所以条件只能动态填充
	this.mongo.Collection("website").FindOne(context.Background(), bson.M{}).Decode(&website);
	fmt.Printf("数据:%v\n", website)
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
				"name":       website.Name,
				"category":   website.Category,
				"components": website.Components,
				"url":        website.Url,
			},
		})
	fmt.Println(result.MatchedCount, result.ModifiedCount)
	if err != nil {
		return err
	}
	return nil
}

func (this *websiteService) DeleteWebsite(websiteID primitive.ObjectID) (err error) {
	result, err := this.mongo.Collection("website").DeleteOne(context.Background(), bson.M{"_id": websiteID})
	if err != nil {
		return err
	}
	fmt.Println(result.DeletedCount)
	return nil
}
