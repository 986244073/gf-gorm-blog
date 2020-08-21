package model

import "github.com/jinzhu/gorm"

type Article struct {
	Category Category
	gorm.Model
	Title   string `gorm:"type:varchar(20);not null" json:"title"`
	Cid     int    `gorm:"type:int(11);not null" json:"cid"`
	Desc    string `gorm:"type:varchar(200);" json:"desc"`
	Content string `gorm:"type:longtext;not null" json:"content"`
	Img     string `gorm:"type:varchar(200);" json:"img"`

}
