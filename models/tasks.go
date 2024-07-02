package models

type Taskz struct {
	Id      uint  `json:"id" gorm:"column:id"`
	ProjectName  string `json:"project_name" gorm:"column:project_name"`
	Name string `json:"name" gorm:"column:name"`
	Description string `json:"description" gorm:"column:description"`
	Status int `json:"status" gorm:"column:status"`
	UpdatedAt int `json:"-" gorm:"autoUpdateTime"`
	CreatedAt int `json:"-" gorm:"autoCreateTime"`
}

func (Taskz) TableName() string {
	return "tasks"
}

