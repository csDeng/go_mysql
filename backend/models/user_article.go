package models

import "gorm.io/gorm"

type UserArticleRelation struct {
	gorm.Model
	UserID    int     `json:"user_id"`
	ArticleID int     `json:"article_id"`
	User      User    `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Article   Article `gorm:"foreignKey:ArticleID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
