package models


type User struct {
	Id uint `json:"id" gorm:"column:id"`
	UserName string `json:"user_name" gorm:"column:user_name"`
	Name string `json:"name" gorm:"column:name"`
	Email string `json:"email" gorm:"column:email"`
	Token string `json:"token" gorm:"column:token"`
}

func (User) TableName() string {
	return "users"
}