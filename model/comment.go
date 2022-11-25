package model

import (
	"gorm.io/gorm"
	"myblog/utils/errmsg"
)

type Comment struct {
	gorm.Model
	UserId    uint   `json:"user_id"`
	ArticleId uint   `json:"article_id"`
	Title     string `json:"title"`
	Username  string `json:"username"`
	Content   string `gorm:"type:varchar(500);not null" json:"content"`
	Status    int8   `gorm:"type:tinyint;default:2" json:"status"`
}

// GetComments 获取所有评论
func GetComments(pageSize int, pageNum int) ([]Comment, int64, int) {
	var commentList []Comment
	var total int64
	db.Find(&commentList).Count(&total)
	err = db.Model(&commentList).
		Limit(pageSize).
		Offset((pageNum - 1) * pageSize).
		Order("Created_At DESC").
		Select("comment.id, article.title,user_id,article_id, user.username, comment.content, comment.status,comment.created_at,comment.deleted_at").
		Joins("LEFT JOIN article ON comment.article_id = article.id").
		Joins("LEFT JOIN user ON comment.user_id = user.id").
		Scan(&commentList).
		Error
	if err != nil {
		return commentList, 0, errmsg.ERROR
	}
	return commentList, total, errmsg.SUCCSE
}

// DeleteComment 删除评论
func DeleteComment(id int) int {
	var comment Comment
	err = db.Where("id = ?", id).Delete(&comment).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// CheckComment 通过评论
// isCheck true 通过评论, isCheck false 撤下评论
func CheckComment(id int, data *Comment, isCheck bool) int {
	var comment Comment
	var res Comment
	var art Article
	var maps = make(map[string]interface{})

	maps["status"] = data.Status

	err = db.Model(&comment).
		Where("id	= ?", id).
		Updates(maps).
		First(&res).
		Error
	if isCheck {
		db.Model(&art).
			Where("id = ?", res.ArticleId).
			UpdateColumn("comment_count", gorm.Expr("comment_count + ?", 1))
	} else {
		db.Model(&art).
			Where("id = ?", res.ArticleId).
			UpdateColumn("comment_count", gorm.Expr("comment_count - ?", 1))
	}

	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
