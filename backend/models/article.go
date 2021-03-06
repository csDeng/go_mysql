package models

type Article struct {
	Model
	TagID         int    `json:"tag_id" gorm:"index"`
	Tag           Tag    `json:"tag" gorm:"references:id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Title         string `json:"title" validate:"required"`
	Desc          string `json:"desc" `
	Content       string `json:"content" validate:"required"`
	CreatedBy     string `json:"created_by"`
	ModifiedBy    string `json:"modified_by"`
	State         int    `json:"state" validate:"min=0,max=1"`
	CoverImageUrl string `json:"cover_image_url"`
}

func ExistArticleByID(id int) bool {
	var article Article
	db.Select("id").Where("id = ?", id).First(&article)
	return article.ID > 0

}
func GetArticleTotal(maps interface{}) (count int) {
	db.Model(&Article{}).Where(maps).Count(&count)
	return
}

func GetArticles(pageNum int, pageSize int, maps interface{}) (articles []Article) {
	db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)
	return
}

func GetArticleIndex() (articles []Article) {
	db.Preload("Tag").Order("created_on desc").Offset(0).Limit(5).Find(&articles)
	return
}

func GetArticle(id int) (article Article) {
	db.Where("id = ?", id).First(&article)
	db.Model(&article).Related(&article.Tag)
	return
}
func EditArticle(id int, data interface{}) bool {
	db.Model(&Article{}).Where("id = ?", id).Updates(data)
	return true
}
func AddArticle(data map[string]interface{}) bool {
	db.Create(&Article{
		TagID:         data["tag_id"].(int),
		Title:         data["title"].(string),
		Desc:          data["desc"].(string),
		Content:       data["content"].(string),
		CreatedBy:     data["created_by"].(string),
		State:         data["state"].(int),
		CoverImageUrl: data["cover_image_url"].(string),
	})
	return true
}
func DeleteArticle(id int) bool {
	db.Where("id = ?", id).Delete(Article{})
	return true
}

func CleanAllArticle() bool {
	db.Unscoped().Where("deleted_on != ? ", 0).Delete(&Article{})
	return true
}
