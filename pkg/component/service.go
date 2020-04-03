package component

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	ComponentService interface {
		GetComponentList(offset, limit int, search Component) (component []Component, count int64, err error)
		CreateComponent(component *Component) (err error)
		UpdateComponent(component *Component) (err error)
		DeleteComponent(componentID int) (err error)
	}
)
type componentService struct {
	mongo *mongo.Database
}

func NewService(mongo *mongo.Database) ComponentService {
	return &componentService{
		mongo: mongo,
	}
}

func (u *componentService) GetComponentList(offset, limit int, search Component) (components []Component, count int64, err error) {
	//一次查询多条数据
	// 查询createtime>=3
	// 限制取2条
	// createtime从大到小排序的数据
	var cursor *mongo.Cursor
	if cursor, err = u.mongo.Collection("test").Find(context.Background(), bson.M{"createtime": bson.M{"$gte": 2}}, options.Find().SetLimit(2), options.Find().SetSort(bson.M{"createtime": -1})); err != nil {
		return nil, 0, err
	}
	defer cursor.Close(context.Background())
	component := Component{}
	for cursor.Next(context.Background()) {
		if err = cursor.Decode(&component); err != nil {

		}
		components = append(components, component)
	}
	//查询集合里面有多少数据
	if count, err = u.mongo.Collection("").CountDocuments(context.Background(), bson.D{}); err != nil {
		return nil, 0, err
	}

	fmt.Printf("Count里面有多少条数据:%d\n", count)
	return components, count, err
}

func (u *componentService) CreateComponent(component *Component) (err error) {
	_, err = u.mongo.Collection("").InsertOne(context.Background(), component)
	if err != nil {
		return err
	}
	return nil
}

func (u *componentService) UpdateComponent(component *Component) (err error) {
	err = u.mongo.Collection("").FindOneAndUpdate(context.Background(), bson.D{{"name", "howie_4"}}, bson.M{"$set": bson.M{"name": "这条数据我需要修改了"}}).Decode(&component)
	if err != nil {
		return err
	}
	return nil
}

func (u *componentService) DeleteComponent(componentID int) (err error) {
	u.mongo.Collection("").DeleteOne(context.Background(), bson.D{})
	if err != nil {
		return err
	}
	return nil
}
