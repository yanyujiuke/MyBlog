package model

import "myblog/utils/errmsg"

type Profile struct {
	ID        int    `gorm:"primary_key" json:"id"`
	Name      string `gorm:"type:varchar(20)" json:"name"`
	Desc      string `gorm:"type:varchar(200)" json:"desc"`
	Qqchat    string `gorm:"type:varchar(200)" json:"qq_chat"`
	Wecaht    string `gorm:"type:varchar(100)" json:"wecaht"`
	Weibo     string `gorm:"type:varchar(200)" json:"weibo"`
	Bili      string `gorm:"type:varchar(200)" json:"bili"`
	Email     string `gorm:"type:varchar(200)" json:"email"`
	Img       string `gorm:"type:varchar(200)" json:"img"`
	Avatar    string `gorm:"type:varchar(200)" json:"avatar"`
	IcpRecord string `gorm:"type:varchar(200)" json:"icp_record"`
}

// GetProfile 获取个人信息
func GetProfile(id int) (pro Profile, code int) {
	err = db.Where("id = ?", id).First(&pro).Error
	if err != nil {
		return pro, errmsg.ERROR
	}
	return pro, errmsg.SUCCSE
}

// UpdateProfile 更新个人信息
func UpdateProfile(id int, pro *Profile) int {
	var profile Profile
	err = db.Model(&profile).Where("id = ?", id).Updates(&pro).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
