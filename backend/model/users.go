package model

type Users struct {
	ID       int64  `gorm:"type:bigint(20) NOT NULL auto_increment;primary_key;"`
	Username string `gorm:"type:varchar(50) NOT NULL;"`
	Password string `gorm:"type:varchar(50) NOT NULL;"`
}