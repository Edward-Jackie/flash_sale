package models

type User struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (u User) TableName() string {
	return "user"
}
