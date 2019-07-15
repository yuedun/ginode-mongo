package model

type User struct {
	Mobile      string `json:"mobile"`
	UserName    string `json:"user_name"`
	Gender      string `json:"gender"`
	Addr        string `json:"addr"`
	Email       string `json:"email"`
	Status      int    `json:"status"`
	Description string `json:"description"`
}
