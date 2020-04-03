package website

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	WebsiteService interface {
		GetWebsiteList(offset, limit int, search Website) (website []Website, count int64, err error)
		CreateWebsite(website *Website) (err error)
		UpdateWebsite(website *Website) (err error)
		DeleteWebsite(websiteID int) (err error)
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

func (this *websiteService) GetWebsiteList(offset, limit int, search Website) (websites []Website, count int64, err error) {
	//一次查询多条数据
	// 查询createtime>=3
	// 限制取2条
	// createtime从大到小排序的数据
	var cursor *mongo.Cursor
	if cursor, err = this.mongo.Collection("test").Find(context.Background(), bson.M{"createtime": bson.M{"$gte": 2}}, options.Find().SetLimit(2), options.Find().SetSort(bson.M{"createtime": -1})); err != nil {
		return nil, 0, err
	}
	defer cursor.Close(context.Background())
	website := Website{}
	for cursor.Next(context.Background()) {
		if err = cursor.Decode(&website); err != nil {

		}
		websites = append(websites, website)
	}
	//查询集合里面有多少数据
	if count, err = this.mongo.Collection("").CountDocuments(context.Background(), bson.D{}); err != nil {
		return nil, 0, err
	}

	fmt.Printf("Count里面有多少条数据:%d\n", count)
	return websites, count, err
}

func (this *websiteService) CreateWebsite(website *Website) (err error) {
	_, err = this.mongo.Collection("").InsertOne(context.Background(), website)
	if err != nil {
		return err
	}
	return nil
}

func (this *websiteService) UpdateWebsite(website *Website) (err error) {
	err = this.mongo.Collection("").FindOneAndUpdate(context.Background(), bson.D{{"name", "howie_4"}}, bson.M{"$set": bson.M{"name": "这条数据我需要修改了"}}).Decode(&website)
	if err != nil {
		return err
	}
	return nil
}

func (this *websiteService) DeleteWebsite(websiteID int) (err error) {
	this.mongo.Collection("").DeleteOne(context.Background(), bson.D{})
	if err != nil {
		return err
	}
	return nil
}
