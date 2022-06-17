package models

import "gorm.io/gorm"

type ArticleTagRelation struct {
	gorm.Model
	ArticleID int     `json:"article_id"`
	TagID     int     `json:"tag_id"`
	Tag       Tag     `gorm:"foreignKey:TagID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Article   Article `gorm:"foreignKey:ArticleID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
