package model

import "time"

type User struct {
	Id          int       `json:"id"`
	Mobile      string    `json:"mobile"`
	UserName    string    `json:"user_name"`
	Gender      string    `json:"gender"`
	Addr        string    `json:"addr"`
	Email       string    `json:"email"`
	Status      int       `json:"status"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// 设置User的表名为`user`
func (User) TableName() string {
	return "user"
}
