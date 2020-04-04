package user

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	Id          primitive.ObjectID `json:"_id",bson:"_id,omitempty"` // bson是用来创建后返回，omitempty是创建时自动创建
	Mobile      string             `json:"mobile",bson:"mobile,omitempty"`
	UserName    string             `json:"user_name"bson:"userName,omitempty"`
	Password    string             `json:"password",bson:"password,omitempty"`
	Gender      string             `json:"gender",bson:"gender,omitempty"`
	Addr        string             `json:"addr",bson:"addr,omitempty"`
	Email       string             `json:"email",bson:"email,omitempty"`
	Status      int                `json:"status",bson:"status,omitempty"`
	Description string             `json:"description",bson:"description,omitempty"`
	CreatedAt   time.Time          `json:"created_at",bson:"createdAt,omitempty"`
	UpdatedAt   time.Time          `json:"updated_at",bson:"updatedAt,omitempty"`
}
