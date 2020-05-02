package db

import (
	"context"
	"fmt"
	"time"

	"github.com/yuedun/ginode-mongo/util"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var Connect *mongo.Client

func init() {
	conf := util.Conf{}
	c := conf.GetConf()
	client, err := mongo.NewClient(options.Client().ApplyURI(fmt.Sprintf("mongodb://%v:%v@%v/%v", c.User, c.Pwd, c.Host, c.Dbname)))
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.TODO(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Connected to MongoDB!")
	Connect = client
}

// 用于支持多数据库
func NewDB(dbname string) *mongo.Database {
	return Connect.Database(dbname)
}
