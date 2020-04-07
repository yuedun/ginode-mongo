package website

import (
	"github.com/yuedun/ginode-mongo/pkg/component"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Website struct {
	ID          primitive.ObjectID    `json:"_id" bson:"_id,omitempty"`                 // bson是用来创建后返回，omitempty是可选
	Name        string                `json:"name" bson:"name,omitempty"`               //网站名 TDK title
	Description string                `json:"description" bson:"description,omitempty"` // 描述 TDK description
	Keywords    string                `json:"keywords" bson:"keywords,omitempty"`       //关键字 TDK keywords
	Icon        string                `json:"icon" bson:"icon,omitempty"`               //图标
	Category    string                `json:"category" bson:"category,omitempty"`       //网站分类
	Url         string                `json:"url" bson:"url,omitempty"`                 //网站地址
	Components  []component.Component `json:"components" bson:"components,omitempty"`   //包含组件
	Status      int                   `json:"status" bson:"status,omitempty"`           //状态
	CreatedAt   time.Time             `json:"created_at" bson:"createdAt,omitempty"`
	UpdatedAt   time.Time             `json:"updated_at" bson:"updatedAt,omitempty"`
}
