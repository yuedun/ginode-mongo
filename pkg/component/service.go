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
		GetComponentList(offset, limit int64, search Component) (component []Component, count int64, err error)
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

func (u *componentService) GetComponentList(offset, limit int64, search Component) (components []Component, count int64, err error) {
	var cursor *mongo.Cursor
	if cursor, err = u.mongo.Collection("component").Find(context.Background(),
		bson.M{},
		options.Find().SetLimit(limit),
		options.Find().SetSkip(offset),
		options.Find().SetSort(bson.M{"_id": -1})); err != nil {
		return nil, 0, err
	}
	defer cursor.Close(context.Background())
	if err = cursor.All(context.Background(), &components); err != nil {
		return nil, 0, err
	}
	//如需对数据迭代处理使用下面方法，不需要则直接使用上面方法
	//for cursor.Next(context.Background()) {
	//	component := Component{}
	//	if err = cursor.Decode(&component); err != nil {
	//		fmt.Println(">>>>>>>>>.")
	//		return nil, 0, err
	//	}
	//	fmt.Println(component)
	//	components = append(components, component)
	//}
	fmt.Println(components)
	//查询集合里面有多少数据
	if count, err = u.mongo.Collection("component").CountDocuments(context.Background(), bson.D{}); err != nil {
		return nil, 0, err
	}

	fmt.Printf("Count里面有多少条数据:%d\n", count)
	return components, count, err
}

func (u *componentService) CreateComponent(component *Component) (err error) {
	_, err = u.mongo.Collection("component").InsertOne(context.Background(), component)
	if err != nil {
		return err
	}
	return nil
}

func (u *componentService) UpdateComponent(component *Component) (err error) {
	err = u.mongo.Collection("component").FindOneAndUpdate(context.Background(), bson.D{{"name", "howie_4"}}, bson.M{"$set": bson.M{"name": "这条数据我需要修改了"}}).Decode(&component)
	if err != nil {
		return err
	}
	return nil
}

func (u *componentService) DeleteComponent(componentID int) (err error) {
	u.mongo.Collection("component").DeleteOne(context.Background(), bson.D{})
	if err != nil {
		return err
	}
	return nil
}
