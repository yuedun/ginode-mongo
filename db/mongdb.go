package db

import (
	"context"
	"fmt"
	"github.com/yuedun/ginode-mongo/util"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

var Mongodb *mongo.Database

func init() {
	conf := util.Conf{}
	c := conf.GetConf()
	client, err := mongo.NewClient(options.Client().ApplyURI(fmt.Sprintf("mongodb://%v:%v@%v/%v", c.User, c.Pwd, c.Host, c.Dbname)))
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	if err = client.Ping(context.Background(), readpref.Primary()); err != nil {
		panic(err)
	}
	Mongodb = client.Database("website")
}
