package models

type Activity struct {
	Id      uint  `json:"id" gorm:"column:id"`
	TaskId  int64 `json:"task_id" gorm:"column:task_id"`
	WorkDay string `json:"work_day" gorm:"column:work_day"`
	Hour int `json:"hour" gorm:"column:hour"`
	Minute int `json:"minute" gorm:"column:minute"`
	KeyBoard int `json:"key_board" gorm:"column:key_board"`
	Mouse int `json:"mouse" gorm:"column:mouse"`
	UpdatedAt int `json:"-" gorm:"autoUpdateTime"`
	CreatedAt int `json:"-" gorm:"autoCreateTime"`
}

func (Activity) TableName() string {
	return "activities"
}

