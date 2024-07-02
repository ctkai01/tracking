package models

type Sync struct {
	Id      uint  `json:"id" gorm:"column:id"`
	Productivity  int `json:"productivity" gorm:"column:productivity"`
	WorkingTime  int `json:"working_time" gorm:"column:working_time"`
}

func (Sync) TableName() string {
	return "sync"
}

