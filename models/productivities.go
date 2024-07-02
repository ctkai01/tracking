package models

type Productivity struct {
	Id      uint  `json:"id" gorm:"column:id"`
	TaskId  int64 `json:"task_id" gorm:"column:task_id"`
	WorkDay string `json:"work_day" gorm:"column:work_day"`
	AppTitle string `json:"app_title" gorm:"column:app_title"`
	AppClass string `json:"app_class" gorm:"column:app_class"`
	Url string `json:"url" gorm:"column:url"`
	Duration  int `json:"duration" gorm:"column:duration"`
	UpdatedAt int `json:"-" gorm:"autoUpdateTime"`
	CreatedAt int `json:"-" gorm:"autoCreateTime"`
}

func (Productivity) TableName() string {
	return "productivities"
}

