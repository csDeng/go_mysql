package models

import (
	"fmt"
	"log"

	"github.com/csDeng/blog/pkg/setting"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Setup() {

	var err error

	dbConfig := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Name)

	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:               dbConfig,
		DefaultStringSize: 256, //string 类型字段的默认长度
	}), &gorm.Config{})

	if err != nil {
		log.Fatalln("链接数据库错误=> \r\n", err)
	}
	log.Println("=============")
	log.Println(db)
	/* 数据库迁移 */
	db.AutoMigrate(
		&User{},
		&Tag{},
		&Article{},
		&UserArticleRelation{},
		&ArticleTagRelation{},
	)
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&User{},
		&Tag{},
		&Article{},
		&UserArticleRelation{},
		&ArticleTagRelation{},
	)
}
