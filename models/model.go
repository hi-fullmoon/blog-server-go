package models

import (
	"log"
	"time"
)

import "github.com/jinzhu/gorm"
import _ "github.com/jinzhu/gorm/dialects/mysql"

type User struct {
	gorm.Model
	Account     string
	Password    string
	NickName    string `gorm:"size:10"`
	Email       string
	Province    string
	City        string
	County      string
	Website     string
	Profile     string `gorm:"size:200"`
	AvatarImage string `json:"avatar_image"`
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
	Content    string `gorm:"size:6000"`
	LikeCount  int
	ViewCount  int
	CoverImage string
	CategoryID uint
	Category   Category
	Tags       []Tag `gorm:"many2many:article_tags;"`
	YearAt     int
	MonthAt    int
	DayAt      int
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
	UserID    uint
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

// get user information by username
func GetUserByAccount(account string) (*User, error) {
	var user User
	if err = db.Where("account = ?", account).First(&user).Error; err != nil {
		return &user, err
	}
	return &user, nil
}

// get user information by user id
func GetUserById(id uint) (*User, error) {
	var user User
	if err = db.Where("id = ?", id).First(&user).Error; err != nil {
		return &User{}, err
	}
	return &user, nil
}

// update user information
func UpdateUserInfo(user *User) error {
	if err = db.Model(&User{}).Updates(user).Error; err != nil {
		return err
	}
	return nil
}

// update user password
func UpdateUserPassword(id uint, pwd string) error {
	var user User
	if err = db.Model(&user).Where("id = ?", id).Update("password", pwd).Error; err != nil {
		return err
	}
	return nil
}

// judging whether it exists by tag name
func CheckCategoryExistByName(name string) (*Category, bool) {
	var category Category
	if err = db.Where("name = ?", name).First(&category).Error; err != nil {
		return &category, false
	}
	return &category, true
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
func GetCategoryList(name string) ([]*Category, error) {
	var categories []*Category
	db := db

	if name != "" {
		db = db.Where("name LIKE ?", "%"+name+"%")
	}

	if err = db.Order("created_at DESC").Preload("Articles").Find(&categories).Error; err != nil {
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

// get category count
func GetCategoryCount() (int, error) {
	var count int
	if err = db.Model(&Category{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// judging whether it exists by tag name
func CheckTagExistByName(name string) (*Tag, bool ){
	var tag Tag
	if err = db.Where("name = ?", name).First(&tag).Error; err != nil {
		return &tag, false
	}
	return &tag, true
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
func GetTagList(name string, pageSize, pageNum int) ([]*Tag, int, error) {
	var tags []*Tag
	db := db

	if name != "" {
		db = db.Where("name LIKE ?", "%"+name+"%")
	}

	if pageSize == 0 {
		pageSize = 10
	}
	if pageNum == 0 {
		pageNum = 1
	}
	offset := pageSize * (pageNum - 1)
	db = db.Limit(pageSize).Offset(offset).Preload("Articles").Order("created_at DESC").Find(&tags)

	var pageTotal int
	db = db.Limit(-1).Count(&pageTotal)

	if err = db.Error; err != nil {
		return []*Tag{}, 0, err
	}

	return tags, pageTotal, nil
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
	if err = db.Exec("DELETE FROM article_tags WHERE tag_id = ?", id).Error; err != nil {
		return err
	}
	return nil
}

// get tag count
func GetTagCount() (int, error) {
	var count int
	if err = db.Model(&Tag{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// create article
func CreateArticle(categoryId uint, title, desc, content, image string, tagIds []uint) error {
	var article Article

	nowTime := time.Now()

	article = Article{
		Title:      title,
		CategoryID: categoryId,
		Desc:       desc,
		Content:    content,
		CoverImage: image,
		YearAt:     nowTime.Year(),
		MonthAt:    int(nowTime.Month()),
		DayAt:      nowTime.Day(),
	}

	if err = db.Create(&article).Error; err != nil {
		return err
	}

	for _, tagId := range tagIds {
		if err = db.Exec("INSERT INTO article_tags(article_id, tag_id) VALUES(?, ?)", article.ID, tagId).Error; err != nil {
			return err
		}
	}
	return nil
}

// read article list
func GetArticleList(title, createdStartAt, createdEndAt, updatedStartAt, updatedEndAt string,
	categoryId, tagId uint, pageSize, pageNum int) ([]*Article, int, error) {
	var articles []*Article
	db := db
	if tagId != 0 {
		db = db.Table("tags").
			Select("articles.*").
			Joins("INNER JOIN article_tags ON tags.id = article_tags.tag_id").
			Joins("INNER JOIN articles ON article_tags.article_id = articles.id").
			Where("tags.id = ? AND articles.deleted_at IS NULL", tagId)
	}

	if title != "" {
		db = db.Where("articles.title LIKE ?", "%"+title+"%")
	}

	if categoryId != 0 {
		db = db.Where("articles.category_id = ?", categoryId)
	}

	if createdStartAt != "" && createdEndAt != "" {
		db = db.Where("articles.created_at BETWEEN ? AND ?", createdStartAt, createdEndAt)
	}

	if updatedStartAt != "" && updatedEndAt != "" {
		db = db.Where("articles.updated_at BETWEEN ? AND ?", updatedStartAt, updatedEndAt)
	}

	if pageSize == 0 {
		pageSize = 10
	}
	if pageNum == 0 {
		pageNum = 1
	}
	offset := pageSize * (pageNum - 1)

	db = db.Preload("Category").Preload("Tags").Order("created_at DESC").
		Limit(pageSize).Offset(offset).Find(&articles)

	var pageTotal int
	db = db.Limit(-1).Count(&pageTotal)

	if err = db.Error; err != nil {
		return []*Article{}, 0, err
	}

	return articles, pageTotal, nil
}

// get article list by article name
func GetArticleListByArticleTitle(articleTitle string) ([]*Article, error) {
	var articles []*Article
	db := db
	db = db.Where("title LIKE ?", "%"+articleTitle+"%").Select("id, title").Order("created_at DESC").Find(&articles)

	if err = db.Error; err != nil {
		return []*Article{}, err
	}

	return articles, nil
}

// get article list by tag name
func GetArticleListByTagName(tagName string) ([]*Article, error) {
	var articles []*Article
	db := db
	if tagName != "" {
		db = db.Table("tags").
			Select("articles.*").
			Joins("INNER JOIN article_tags ON tags.id = article_tags.tag_id").
			Joins("INNER JOIN articles ON article_tags.article_id = articles.id").
			Where("tags.name = ? AND articles.deleted_at IS NULL", tagName)
	}

	db = db.Preload("Tags").Order("created_at DESC").Find(&articles)

	if err = db.Error; err != nil {
		return []*Article{}, err
	}

	return articles, nil
}

// get article list by category name
func GetArticleListByCategoryName(categoryName string) ([]*Article, error) {
	var articles []*Article
	db := db

	db = db.Table("categories").
		Where("name = ?", categoryName).
		Select("articles.*").
		Joins("INNER JOIN articles ON articles.category_id = categories.id").
		Order("created_at DESC").
		Find(&articles)

	if err = db.Error; err != nil {
		return []*Article{}, err
	}

	return articles, nil
}

// get article list by group
// TODO：need to be modified.
func GetArticleByGroup() ([]map[string]interface{}, error) {
	var articles []*Article
	var yearAt, monthAt int

	rows, err := db.Table("articles").Select("year_at, month_at").Group("year_at, month_at").Order("year_at, month_at DESC").Rows()
	if err = db.Error; err != nil {
		return nil, err
	}

	out := make([]map[string]interface{}, 0, 10)

	for rows.Next() {
		rows.Scan(&yearAt, &monthAt)

		res := db.Where("year_at = ? AND month_at = ?", yearAt, monthAt).Select("id, year_at, month_at, day_at, title").Order("created_at DESC").Find(&articles)
		if res.Error != nil {
			return nil, err
		}
		m := map[string]interface{}{
			"YearAt":   yearAt,
			"MonthAt":  monthAt,
			"Articles": articles,
		}
		out = append(out, m)
	}

	return out, nil
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
func UpdateArticle(aid, categoryId uint, title, desc, content, image string, tagIds []uint) error {
	var article Article
	res := db.Where("id = ?", aid).Model(&article).Updates(map[string]interface{}{
		"category_id": categoryId,
		"title":       title,
		"desc":        desc,
		"content":     content,
		"cover_image": image,
	})
	if err = res.Error; err != nil {
		return err
	}

	if err = db.Exec("DELETE FROM article_tags WHERE article_id = ?", aid).Error; err != nil {
		return err
	}
	for _, tagId := range tagIds {
		if err = db.Exec("INSERT INTO article_tags(article_id, tag_id) VALUES(?, ?)", aid, tagId).Error; err != nil {
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

// get article count
func GetArticleCount() (int, error) {
	var count int
	if err = db.Model(&Article{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// update article view count
func UpdateArticleViewCount(aid uint) (int, error) {
	var article Article
	if err = db.Model(&article).Where("id = ?", aid).Update("view_count", gorm.Expr("view_count + ?", 1)).Error; err != nil {
		return 0, err
	}
	return article.ViewCount, err
}
