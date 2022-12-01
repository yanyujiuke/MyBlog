package model

import (
	"gorm.io/gorm"
	"myblog/utils/errmsg"
)

type Article struct {
	Category Category `gorm:"foreignKey:Cid"`
	gorm.Model
	Title        string `gorm:"type:varchar(100);not null" json:"title"`
	Cid          int    `gorm:"type:int;not null" json:"cid"`
	Desc         string `gorm:"type:varchar(200)" json:"desc"`
	Content      string `gorm:"type:longtext" json:"content"`
	Img          string `gorm:"type:varchar(100)" json:"img"`
	CommentCount int    `gorm:"type:int;not null" json:"comment_count"`
	ReadCount    int    `gorm:"type:int;not null" json:"read_count"`
}

// GetArtInfo 获取单个文章信息
func GetArtInfo(id int) (Article, int) {
	var art Article

	err = db.Where("id=?", id).Preload("Category").First(&art).Error
	db.Model(&art).Where("id=?", id).UpdateColumn("read_count", gorm.Expr("read_count + ?", 1))
	if err != nil {
		println(err.Error())
		return art, errmsg.ERROR_ART_NOT_EXIST
	}
	return art, errmsg.SUCCSE
}

// GetArts 获取文章列表
func GetArts(pageSize int, pageNum int, title string) (art []Article, total int64, code int) {
	if len(title) > 0 {
		err = db.Where("title LIKE ?", title+"%").
			Limit(pageSize).
			Offset((pageNum - 1) * pageSize).
			Order("Created_At DESC").
			Joins("Category").
			Find(&art).
			Error
		db.Model(&art).Where("title LIKE ?", title).Count(&total)
	} else {
		err = db.Limit(pageSize).
			Offset((pageNum - 1) * pageSize).
			Order("Created_At DESC").
			Joins("Category").
			Find(&art).
			Error
		db.Model(&art).Count(&total)
	}
	code = errmsg.SUCCSE
	if err != nil {
		code = errmsg.ERROR
	}
	return art, total, code
}

// CreateArt 添加文章
func CreateArt(art *Article) int {
	err := db.Create(&art).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// EditArt 编辑文章
func EditArt(id int, art *Article) int {
	err := db.Where("id = ?", id).Updates(&art).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// DeleteArt 删除文章
func DeleteArt(id int) int {
	var art Article
	err := db.Where("id = ?", id).Delete(&art).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// GetArtsByCid 获取某个分类下的所有文章
func GetArtsByCid(cid int, pageSize int, pageNum int) ([]Article, int64, int) {
	var arts []Article
	var total int64

	err := db.
		Preload("Category").
		Where("cid = ?", cid).
		Limit(pageSize).
		Offset((pageNum - 1) * pageSize).
		Find(&arts).
		Error
	db.Model(&arts).Where("cid = ?", cid).Count(&total)
	if err != nil {
		return arts, total, errmsg.ERROR
	}
	return arts, total, errmsg.SUCCSE
}
