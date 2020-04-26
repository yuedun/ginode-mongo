package user

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	UserService interface {
		GetUserInfoByName(username string) (user User, err error)
		GetUserList(offset, limit int64, search User) (user []User, count int64, err error)
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

func (this *userService) GetUserInfoByName(username string) (user User, err error) {
	if err = this.mongo.Collection("user").FindOne(context.Background(), bson.M{"username": username}).Decode(&user); err != nil {
		return user, err
	}
	return user, nil
}

func (this *userService) GetUserList(offset, limit int64, search User) (users []User, count int64, err error) {
	//一次查询多条数据
	// 查询createtime>=3
	// 限制取2条
	// createtime从大到小排序的数据
	var cursor *mongo.Cursor
	if cursor, err = this.mongo.Collection("user").Find(
		context.Background(),
		bson.M{},
		options.Find().SetSkip(offset), options.Find().SetLimit(limit), options.Find().SetSort(bson.M{"createtime": -1})); err != nil {
		return nil, 0, err
	}
	//for cursor.Next(context.Background()) {
	//	user := User{}
	//	if err = cursor.Decode(&user); err != nil {
	//
	//	}
	//	users = append(users, user)
	//}
	if err = cursor.All(context.Background(), &users); err != nil {
		return nil, 0, err
	}
	//查询集合里面有多少数据
	if count, err = this.mongo.Collection("user").CountDocuments(context.Background(), bson.D{}); err != nil {
		return nil, 0, err
	}

	fmt.Printf("Count里面有多少条数据:%d\n", count)
	return users, count, err
}

func (this *userService) CreateUser(user *User) (err error) {
	result, err := this.mongo.Collection("user").InsertOne(context.Background(), user)
	if err != nil {
		return err
	}
	user.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (this *userService) UpdateUser(user *User) (err error) {
	result := this.mongo.Collection("user").FindOneAndUpdate(
		context.Background(),
		bson.D{{"_id", user.ID}},
		bson.M{
			"$set": bson.M{
				"name":   user.Username,
				"mobile": user.Mobile,
			},
		})
	if result.Err() != nil {
		return result.Err()
	}
	return nil
}

func (this *userService) DeleteUser(userID int) (err error) {
	result, err:=this.mongo.Collection("user").UpdateOne(context.Background(), bson.M{"_id": userID}, bson.M{"status": 0})
	if err != nil {
		return err
	}
	fmt.Println(result.ModifiedCount)
	return nil
}
