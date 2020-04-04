package website

import (
	"github.com/yuedun/ginode-mongo/pkg/component"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Website struct {
	ID         primitive.ObjectID    `bson:"_id",omitempty`                // bson是用来创建后返回，omitempty是可选
	Name       string                `json:"name",bson:"name"`                       //网站名
	Category   string                `json:"category",bson:"category"`               //网站分类
	Url        string                `json:"url",bson:"url"`                         //网站地址
	Components []component.Component `json:"components",bson:"components",omitempty` //包含组件
	Status     int                   `json:"status",bson:"status"`                   //状态
	CreatedAt  time.Time             `json:"created_at",bson:"createdAt"`
	UpdatedAt  time.Time             `json:"updated_at",bson:"updatedAt"`
}
