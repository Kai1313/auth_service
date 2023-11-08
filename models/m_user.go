package models

import "time"

type MUser struct {
	UserId        string    `gorm:"PrimaryKey;type:varchar(50)"`
	Username      string    `gorm:"type:varchar(255);uniqueIndex;"`
	Password      string    `gorm:"type:varchar(255)"`
	Jabatan       string    `gorm:"type:varchar(50)"`
	Email         string    `gorm:"type:varchar(255)"`
	Type          int       `gorm:"type:integer"`
	RememberToken string    `gorm:"type:varchar(255)"`
	ApiToken      string    `gorm:"type:varchar(255)"`
	SessionId     string    `gorm:"type:varchar(255)"`
	IpAddress     string    `gorm:"type:varchar(255)"`
	IsLocked      int       `gorm:"type:integer"`
	CompanyId     string    `gorm:"type:varchar(50)"`
	Notes         string    `gorm:"type:text"`
	DtRecord      time.Time `gorm:"type:timestamp without time zone;"`
	UserRecord    string    `gorm:"type:varchar(255)"`
	DtModified    time.Time `gorm:"type:timestamp without time zone;"`
	UserModified  string    `gorm:"type:varchar(255)"`
}

func (MUser) TableName() string {
	return "m_user"
}