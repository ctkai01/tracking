package models

type IdleTime struct {
	Id      uint  `json:"id" gorm:"column:id"`
	TaskId  int64 `json:"task_id" gorm:"column:task_id"`
	WorkDay string `json:"work_day" gorm:"column:work_day"`
	StartAt   int `json:"start_at" gorm:"column:start_at"`
	EndAt     int `json:"end_at" gorm:"column:end_at"`
	Duration  int `json:"duration" gorm:"column:duration"`
	UpdatedAt int `json:"-" gorm:"autoUpdateTime"`
	CreatedAt int `json:"-" gorm:"autoCreateTime"`
}

func (IdleTime) TableName() string {
	return "idle_times"
}

