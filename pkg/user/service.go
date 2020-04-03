package user

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	UserService interface {
		GetUserInfo(userObj User) (user User, err error)
		GetUserList(offset, limit int, search User) (user []User, count int64, err error)
		CreateUser(user *User) (err error)
		UpdateUser(user *User) (err error)
		DeleteUser(userID int) (err error)
	}
)
type userService struct {
	mongo *mongo.Database
}

func NewService(mongo *mongo.Database) UserService {
	return &userService{
		mongo: mongo,
	}
}

func (this *userService) GetUserInfo(usersearch User) (user User, err error) {
	if err = this.mongo.Collection("user").FindOne(context.Background(), bson.M{"_id": usersearch.Id}).Decode(&user); err != nil {
		return user, err
	}
	return user, nil
}

func (this *userService) GetUserList(offset, limit int, search User) (users []User, count int64, err error) {
	//一次查询多条数据
	// 查询createtime>=3
	// 限制取2条
	// createtime从大到小排序的数据
	var cursor *mongo.Cursor
	if cursor, err = this.mongo.Collection("user").Find(context.Background(), bson.M{"createtime": bson.M{"$gte": 2}}, options.Find().SetLimit(2), options.Find().SetSort(bson.M{"createtime": -1})); err != nil {
		return nil, 0, err
	}
	defer cursor.Close(context.Background())
	user := User{}
	for cursor.Next(context.Background()) {
		if err = cursor.Decode(&user); err != nil {

		}
		users = append(users, user)
	}
	//查询集合里面有多少数据
	if count, err = this.mongo.Collection("user").CountDocuments(context.Background(), bson.D{}); err != nil {
		return nil, 0, err
	}

	fmt.Printf("Count里面有多少条数据:%d\n", count)
	return users, count, err
}

func (this *userService) CreateUser(user *User) (err error) {
	_, err = this.mongo.Collection("user").InsertOne(context.Background(), user)
	if err != nil {
		return err
	}
	return nil
}

func (this *userService) UpdateUser(user *User) (err error) {
	err = this.mongo.Collection("user").FindOneAndUpdate(context.Background(), bson.D{{"name", "howie_4"}}, bson.M{"$set": bson.M{"name": "这条数据我需要修改了"}}).Decode(&user)
	if err != nil {
		return err
	}
	return nil
}

func (this *userService) DeleteUser(userID int) (err error) {
	this.mongo.Collection("user").DeleteOne(context.Background(), bson.D{})
	if err != nil {
		return err
	}
	return nil
}
