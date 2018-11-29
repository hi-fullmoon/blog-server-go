package models

import (
	"log"
)

import "github.com/jinzhu/gorm"
import _ "github.com/jinzhu/gorm/dialects/mysql"

type User struct {
	gorm.Model
	UserID   string
	Username string
	Password string
}

type Category struct {
	gorm.Model
	Name     string `gorm:"size:30"`
	Desc     string `gorm:"size:100"`
	Articles []Article
}

type Tag struct {
	gorm.Model
	Name     string    `gorm:"size:30"`
	Desc     string    `gorm:"size:100"`
	Articles []Article `gorm:"many2many:article_tags;"`
}

type Article struct {
	gorm.Model
	Title      string `gorm:"size:30"`
	Desc       string `gorm:"size:200"`
	Content    string `gorm:"size:5000"`
	LikeCount  int
	CategoryID uint
	Category   Category
	Tags       []Tag `gorm:"many2many:article_tags;"`
}

type Comment struct {
	gorm.Model
	Author    string
	Email     string
	Content   string `gorm:"size:200"`
	ArticleID uint
}

type Session struct {
	gorm.Model
	SessionID string
	TTL       int64
	Username  string
	Password  string
}

var (
	db  *gorm.DB
	err error
)

// init db
func InitDB() (*gorm.DB, error) {
	db, err = gorm.Open("mysql", "root:123456@tcp(localhost:3306)/blog?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return nil, err
	}
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.AutoMigrate(&User{}, &Category{}, &Tag{}, &Article{}, &Session{})

	return db, nil
}

// get user
func GetUserInfo(username string) (*User, error) {
	var user User
	if err = db.Where("username = ?", username).First(&user).Error; err != nil {
		return &user, err
	}
	return &user, nil
}

// create category
func CreateCategory(name, desc string) error {
	category := Category{Name: name, Desc: desc}
	if err = db.Create(&category).Error; err != nil {
		return err
	}
	return nil
}

// read category list by name
func ReadCategoryList(name string) ([]*Category, error) {
	var categories []*Category

	if name == "" {
		db.Order("created_at DESC").Preload("Articles").Find(&categories)
	} else {
		db.Where("name = ?", name).Preload("Articles").Order("created_at DESC").Find(&categories)
	}

	if db.Error != nil {
		return []*Category{}, err
	}
	return categories, nil
}

// delete category
func DeleteCategory(id uint) error {
	var category Category
	if err = db.Where("id = ?", id).Delete(&category).Error; err != nil {
		return err
	}
	return nil
}

// update category
func UpdateCategory(id uint, name, desc string) error {
	var category Category
	db.Model(&category).Where("id = ?", id).Updates(map[string]interface{}{
		"id":   id,
		"name": name,
		"desc": desc,
	})
	if err = db.Error; err != nil {
		return err
	}
	return nil
}

// create tag
func CreateTag(name string) error {
	tag := Tag{Name: name}

	if err = db.Create(&tag).Error; err != nil {
		return err
	}
	return nil
}

// read tags by name
func ReadTagList(name string) ([]*Tag, error) {
	var tags []*Tag

	if name == "" {
		db.Order("created_at DESC").Preload("Articles").Find(&tags)
	} else {
		db.Where("name = ?", name).Preload("Articles").Order("created_at DESC").Find(&tags)
	}

	if err = db.Error; err != nil {
		return []*Tag{}, err
	}
	return tags, nil
}

// update tags
func UpdateTag(id uint, name string) error {
	var tag Tag

	if err = db.Model(&tag).Where("id = ?", id).Update("name", name).Error; err != nil {
		return err
	}
	return nil
}

// delete tag
func DeleteTag(id uint) error {
	var tag Tag
	if err = db.Where("id = ?", id).Delete(&tag).Error; err != nil {
		return err
	}
	return nil
}

// create article
func CreateArticle(categoryId uint, title, desc, content string, tagIds []uint) error {
	var article Article

	article = Article{
		Title:      title,
		CategoryID: categoryId,
		Desc:       desc,
		Content:    content,
	}

	if err = db.Create(&article).Error; err != nil {
		return err
	}

	for _, tagId := range tagIds {
		if err = db.Exec("insert into article_tags(article_id, tag_id) values(?, ?)", article.ID, tagId).Error; err != nil {
			return err
		}
	}
	return nil
}

// read article list
func ReadArticleList() ([]*Article, error) {
	var articles []*Article
	if err = db.Preload("Category").Preload("Tags").Find(&articles).Error; err != nil {
		return []*Article{}, err
	}

	return articles, nil
}

// read article information by aid
func ReadArticleInfo(aid uint) (*Article, error) {
	var article Article
	if err = db.Where("id = ?", aid).Preload("Category").Preload("Tags").First(&article).Error; err != nil {
		log.Println(err)
		return &Article{}, err
	}
	return &article, nil
}

// update article
func UpdateArticle(aid, categoryId uint, title, desc, content string, tagIds []uint) error {
	var article Article
	res := db.Where("id = ?", aid).Model(&article).Updates(map[string]interface{}{
		"category_id": categoryId,
		"title":       title,
		"desc":        desc,
		"content":     content,
	})
	if err = res.Error; err != nil {
		return err
	}

	if err = db.Exec("DELETE FROM article_tags WHERE article_id = ?", aid).Error; err != nil {
		return err
	}
	for _, tagId := range tagIds {
		if err = db.Exec("insert into article_tags(article_id, tag_id) values(?, ?)", aid, tagId).Error; err != nil {
			return err
		}
	}
	return nil
}

// delete article
func DeleteArticle(aid uint) error {
	var article Article
	if err = db.Where("id = ?", aid).Delete(&article).Error; err != nil {
		return err
	}
	return nil
}
