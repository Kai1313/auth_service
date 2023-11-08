package models

import "time"

type MUserAccess struct {
	UserId           string    `gorm:"PrimaryKey;type:varchar(50)"`
	WarehouseId      string    `gorm:"PrimaryKey;type:varchar(50)"`
	ModuleFunctionId string    `gorm:"PrimaryKey;type:varchar(50)"`
	DtRecord         time.Time `gorm:"type:timestamp without time zone;"`
	UserRecord       string    `gorm:"type:varchar(255)"`
	DtModified       time.Time `gorm:"type:timestamp without time zone;"`
	UserModified     string    `gorm:"type:varchar(255)"`
}

func (MUserAccess) TableName() string {
	return "m_user_access"
}
