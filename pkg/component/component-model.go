package component

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Component struct {
	ID            primitive.ObjectID `json:"_id" bson:"_id,omitempty"`                      // bson是用来创建后返回，omitempty可选
	Name          string             `json:"name" bson:"name,omitempty"`                    //组件名
	Category      string             `json:"category" bson:"category,omitempty"`            //分类
	Status        int                `json:"status" bson:"status,omitempty"`                // 可用状态
	TitleH1       string             `json:"title_h_1" bson:"titleH1,omitempty"`            // 一级标题
	TitleH2       string             `json:"title_h_2" bson:"titleH2,omitempty"`            // 二级标题
	TitleH3       string             `json:"title_h_3" bson:"titleH3,omitempty"`            // 三级标题
	Description   string             `json:"description" bson:"description,omitempty"`      // 文字描述
	BackgroundImg string             `json:"background_img" bson:"backgroundImg,omitempty"` // 背景图
	BigImg        string             `json:"big_img" bson:"bigImg,omitempty"`               // 展示大图
	Elements      []Element          `json:"elements" bson:"elements,omitempty"`            // 多个同类型元素列表展示
	Links         []Link             `json:"links" bson:"links,omitempty"`                  // 链接或按钮
	Extras        []Component        `json:"extras" bson:"extras,omitempty"`                // 其他补充内容
	CreatedAt     time.Time          `json:"created_at" bson:"createdAt,omitempty"`
	UpdatedAt     time.Time          `json:"updated_at" bson:"updatedAt,omitempty"`
}

type Element struct {
	Icon        string `json:"icon" bson:"icon,omitempty"` //图标
	TitleH1     string `json:"title_h_1" bson:"titleH1,omitempty"`
	TitleH2     string `json:"title_h_2" bson:"titleH2,omitempty"`
	TitleH3     string `json:"title_h_3" bson:"titleH3,omitempty"`
	Description string `json:"description" bson:"description,omitempty"`
	Link        string `json:"link" bson:"link,omitempty"`
}
type Link struct {
	Title string `json:"title" bson:"title,omitempty"`
	Url   string `json:"url" bson:"url,omitempty"`
}
