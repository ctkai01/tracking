package models

type Config struct {
	Id      uint  `json:"id" gorm:"column:id"`
	Screenshot  int `json:"screenshot" gorm:"column:screenshot"`
	Refresh  int `json:"refresh" gorm:"column:refresh"`
	PresignUrl  int `json:"presign_url" gorm:"column:presign_url"`
	UpdatedAt int `json:"-" gorm:"autoUpdateTime"`
	CreatedAt int `json:"-" gorm:"autoCreateTime"`
}

func (Config) TableName() string {
	return "configs"
}

