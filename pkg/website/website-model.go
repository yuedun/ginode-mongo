package website

import "time"

type Website struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Category  string    `json:"category"`
	Url       string    `json:"url"`
	Content   string    `json:"content"`
	Status    int       `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// 设置User的表名为`website`
func (Website) CollectionName() string {
	return "website"
}
