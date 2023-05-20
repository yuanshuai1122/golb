package models

import (
	"GIN/utils"
	"time"
)

// 返回列表使用
type UserInfo struct {
	UserId        int       `gorm:"column:userId" json:"userId"`
	UserType      int       `gorm:"column:userType" json:"userType"`
	UserName      string    `gorm:"column:userName" json:"userName"`
	UserPass      string    `gorm:"column:userPass" json:"userPass"`
	UserEmail     string    `gorm:"column:userEmail" json:"userEmail"`
	UserRegdate   time.Time `gorm:"column:userRegdate" json:"userRegdate"`
	UserSignature string    `gorm:"column:userSignature" json:"userSignature"`
	UserIcon      string    `gorm:"column:userIcon" json:"userIcon"`
}

func (UserInfo) TableName() string  {
	return "userinfo"
}

// 登录
func (UserInfo) Login(userName string) interface{}  {
	loginInfo := []UserInfo{}
	if err :=  utils.DB.Where("userName = ?",userName).Find(&loginInfo).Error; err != nil {
		return 1
	}else {
		return map[string]interface{}{
			"data": loginInfo,
			"userType": loginInfo[0].UserType,
			"userPass": loginInfo[0].UserPass,
		}
	}
}

// 删除用户
func (UserInfo) DeleteUseByUserID(userID string) interface{}  {
	if err :=  utils.DB.Where("userID = ?",userID).Delete(&UserInfo{}).Error; err != nil {
		return 1
	}else {
		return 0
	}
}


// 根据ID获取用户信息
func (UserInfo) ShowUserByUserId(userID string) interface{}  {
	userInfo := UserInfo{}
	if err :=  utils.DB.Where("userID = ?",userID).First(&userInfo).Error; err != nil {
		return 1
	}else {
		return userInfo
	}
}

// 查询用户数量
func (UserInfo) ShowUserCount() interface{}  {
	var count int64
	if err :=  utils.DB.Find(&UserInfo{}).Count(&count).Error; err != nil {
		return 1
	}else {
		return count
	}
}

// 分页查询
func (UserInfo) Findby(currentPage int64,pageSize int64,searchContent string) interface{}  {
	offset := (currentPage - 1) * pageSize
	var count int64
	userInfo := []UserInfo{}
	// 查总数
	if err := utils.DB.Where("userName LIKE ?",searchContent+"%").Find(&userInfo).Count(&count).Error; err != nil {
		return 1
	}
	if err :=  utils.DB.Offset(int(offset)).Limit(int(pageSize)).Where("userName LIKE ?",searchContent+"%").Find(&userInfo).Error; err != nil {
		return 1
	}else {
		totalPage := utils.ReturnTotalPage(count , pageSize)
		return map[string]interface{}{
			"currentPage": currentPage,
			"list": userInfo,
			"pageSize": pageSize,
			"total": count,
			"totalPage": totalPage,
		}
	}
}

// 修改个人信息
func (UserInfo) UpdateUserBasicsInfo(json map[string]interface{}) int  {
	if err :=  utils.DB.Model(&UserInfo{}).Where("userId = ?",json["userId"]).Updates(json).Error; err != nil {
		return 1
	}else {
		return 0
	}
}

// UpdateUserPass 修改密码
func (UserInfo) UpdateUserPass(userId string,userPass string) int  {
	if err :=  utils.DB.Model(&UserInfo{}).Where("userId = ?",userId).Update("userPass",userPass).Error; err != nil {
		return 1
	}else {
		return 0
	}
}

//ShowUserName 查询用户名是否重复
func (UserInfo) ShowUserName(userName string) interface{}  {
	userInfo := UserInfo{}
	if err :=  utils.DB.Where("userName = ?",userName).First(&userInfo).Error; err != nil {
		return 1
	}else {
		return 0
	}
}

// SearchEmail
func (UserInfo) SearchEmail(userEmail string) interface{}  {
	userInfo := UserInfo{}
	if err :=  utils.DB.Where("userEmail = ?",userEmail).First(&userInfo).Error; err != nil {
		return 1
	}else {
		return 0
	}
}

// InsertRegisterInfo
func (UserInfo) InsertRegisterInfo(json map[string]interface{}) interface{}  {
	if err :=  utils.DB.Model(&UserInfo{}).Create(json).Error; err != nil {
		return 1
	}else {
		return 0
	}
}