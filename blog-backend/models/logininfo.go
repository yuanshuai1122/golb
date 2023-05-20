package models

import (
	"GIN/utils"
	"time"
)

type LoginInfo struct {
	LoginId        int       `gorm:"column:loginId" json:"loginId"`
	LoginTime      time.Time       `gorm:"column:loginTime" json:"loginTime"`
}

func (LoginInfo) TableName() string  {
	return "logininfo"
}

// 增加登录记录
func (LoginInfo) InsertLoginInfo(json map[string]interface{}) interface{}  {
	if err :=  utils.DB.Model(&LoginInfo{}).Create(json).Error; err != nil {
		return 1
	}else {
		return 0
	}
}

