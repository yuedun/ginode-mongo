package user

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	Id          primitive.ObjectID `bson:"_id"`
	Mobile      string             `json:"mobile"`
	UserName    string             `json:"userName"`
	Password    string             `json:"password"`
	Gender      string             `json:"gender"`
	Addr        string             `json:"addr"`
	Email       string             `json:"email"`
	Status      int                `json:"status"`
	Description string             `json:"description"`
	CreatedAt   time.Time          `json:"createdAt"`
	UpdatedAt   time.Time          `json:"updatedAt"`
}
