package models

import (
	"GIN/utils"
	"time"
)

// 返回列表使用
type  CommentInfo struct {
	CommentId int `gorm:"column:commentId" json:"commentId"`
	ArticleId int			`gorm:"column:articleId" json:"articleId"`
	UserId int `gorm:"column:userId" json:"userId"`
	Content string `gorm:"column:content" json:"content"`
	ParentId int 	`gorm:"column:parentId" json:"parentId"`
	CommentDate time.Time `gorm:"column:commentDate" json:"commentDate"`
	ArticleTitle string `gorm:"-" json:"articleTitle"`
	UserName string `gorm:"-" json:"userName"`
}

func (CommentInfo) TableName() string  {
	return "commentinfo"
}

// 新增使用
type  AddCommentInfo struct {
	ArticleId int			`gorm:"column:articleId" json:"articleId"`
	UserId int `gorm:"column:userId" json:"userId"`
	Content string `gorm:"column:content" json:"content"`
	ParentId int 	`gorm:"column:parentId" json:"parentId"`
}

func (AddCommentInfo) TableName() string  {
	return "commentinfo"
}

// 返回所有评论 + 个人信息
type  AllCommentInfoAndUserInfo struct {
	UserId int			`gorm:"column:userId" json:"userId"`
	UserType int `gorm:"column:userType" json:"userType"`
	UserName string `gorm:"column:userName" json:"userName"`
	UserEmail string 	`gorm:"column:userEmail" json:"userEmail"`
	UserRegdate time.Time 	`gorm:"column:userRegdate" json:"userRegdate"`
	UserSignature string 	`gorm:"column:userSignature" json:"userSignature"`
	UserIcon string 	`gorm:"column:userIcon" json:"userIcon"`
	Child []CommentInfo `gorm:"many2many:commentinfo" json:"child"`
	CommentInfo
}

func (AllCommentInfoAndUserInfo) TableName() string  {
	return "userinfo"
}


// 根据用户ID删除评论
func (CommentInfo) DeleteUserCommentByUserID(userID string) interface{}  {
	if err :=  utils.DB.Where("userID = ?",userID).Delete(&CommentInfo{}).Error; err != nil {
		return 1
	}else {
		return 0
	}
}

// 管理员查询所有评论 分页
func (CommentInfo) SearchAll(currentPage int64,pageSize int64, searchContent string) interface{}  {
	offset := (currentPage - 1) * pageSize
	var count int64
	commentInfo := []CommentInfo{}
	// 查总数
	if err := utils.DB.Where("content LIKE ?",searchContent+"%").Find(&commentInfo).Count(&count).Error; err != nil {
		return 1
	}
	if err :=  utils.DB.Offset(int(offset)).Limit(int(pageSize)).Where("content LIKE ?",searchContent+"%").Find(&commentInfo).Error; err != nil {
		return 1

	}else {
		totalPage := utils.ReturnTotalPage(count , pageSize)
		return map[string]interface{}{
			"currentPage": currentPage,
			"list": commentInfo,
			"pageSize": pageSize,
			"total": count,
			"totalPage": totalPage,
		}
	}
}

// 普通用户获取自己的评论 分页
func (CommentInfo) SearchByUserID(currentPage int64,pageSize int64,userID string, searchContent string) interface{}  {
	offset := (currentPage - 1) * pageSize
	var count int64
	commentInfo := []CommentInfo{}
	// 查总数
	if err := utils.DB.Where("content LIKE ? AND userID = ?",searchContent+"%",userID).Find(&commentInfo).Count(&count).Error; err != nil {
		return 1
	}
	if err :=  utils.DB.Offset(int(offset)).Limit(int(pageSize)).Where("content LIKE ? AND userID = ?",searchContent+"%",userID).Find(&commentInfo).Error; err != nil {
		return 1
	}else {
		totalPage := utils.ReturnTotalPage(count , pageSize)
		return map[string]interface{}{
			"currentPage": currentPage,
			"list": commentInfo,
			"pageSize": pageSize,
			"total": count,
			"totalPage": totalPage,
		}
	}
}

// 删除
func (CommentInfo) Delete(id string) int  {
	if err :=  utils.DB.Where("commentId = ?",id).Delete(&UrlInfo{}).Error; err != nil {
		return 1
	}else {
		return 0
	}
}