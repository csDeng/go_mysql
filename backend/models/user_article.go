package models

type UserArticleRelation struct {
	Model
	UserId    int     `json:"user_id"`
	ArticleId int     `json:"article_id"`
	User      User    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;ForeignKey:UserId"`
	Article   Article `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;ForeignKey:ArticleId"`
}
