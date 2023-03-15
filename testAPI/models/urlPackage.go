package models

import "gorm.io/gorm"

type UrlPackage struct {
	gorm.Model
	Id  int    `gorm:"primaryKey;auto_increment"`
	Url string `gorm:"not null"`
}

func (UrlPackage) TableName() string {
	return "dbo.UrlPackage"
}
