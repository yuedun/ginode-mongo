package component

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Component 网站组件
type Component struct {
	ID            primitive.ObjectID `json:"_id" bson:"_id,omitempty"`                       // bson是用来创建后返回，omitempty可选
	Name          string             `json:"name" bson:"name,omitempty"`                     //组件名
	Category      string             `json:"category" bson:"category,omitempty"`             //分类
	Status        int                `json:"status" bson:"status,omitempty"`                 // 可用状态
	Title1        string             `json:"title_1" bson:"title_1,omitempty"`               // 一级标题
	Title2        string             `json:"title_2" bson:"title_2,omitempty"`               // 二级标题
	Title3        string             `json:"title_3" bson:"title_3,omitempty"`               // 三级标题
	Description   string             `json:"description" bson:"description,omitempty"`       // 文字描述
	BackgroundImg string             `json:"background_img" bson:"background_img,omitempty"` // 背景图
	BigImg        string             `json:"big_img" bson:"big_img,omitempty"`               // 展示大图
	Elements      []Element          `json:"elements" bson:"elements,omitempty"`             // 多个同类型元素列表展示
	Links         []Link             `json:"links" bson:"links,omitempty"`                   // 链接或按钮
	Extras        []Component        `json:"extras" bson:"extras,omitempty"`                 // 其他补充内容
	CreatedAt     time.Time          `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt     time.Time          `json:"updated_at" bson:"updated_at,omitempty"`
}

//Element 组件中包含的同类型元素
type Element struct {
	Icon        string `json:"icon" bson:"icon,omitempty"` //图标
	Title1      string `json:"title_1" bson:"title_1,omitempty"`
	Title2      string `json:"title_2" bson:"title_2,omitempty"`
	Title3      string `json:"title_3" bson:"title_3,omitempty"`
	Description string `json:"description" bson:"description,omitempty"`
	Link        string `json:"link" bson:"link,omitempty"`
}

//Link 组件包含的链接
type Link struct {
	Title string `json:"title" bson:"title,omitempty"`
	URL   string `json:"url" bson:"url,omitempty"`
}
