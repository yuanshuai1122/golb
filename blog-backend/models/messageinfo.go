package models

import (
	"GIN/utils"
	"time"
)

// 返回列表使用
type  MessageInfo struct {
	MessageId int `gorm:"column:messageId" json:"messageId"`
	MessageName string			`gorm:"column:messageName" json:"messageName"`
	MessageQQ string `gorm:"column:messageQQ" json:"messageQQ"`
	Content string `gorm:"column:content" json:"content"`
	ReplyContent string 	`gorm:"column:replyContent" json:"replyContent"`
	MessageDate time.Time `gorm:"column:messageDate" json:"messageDate"`
}

func (MessageInfo) TableName() string  {
	return "message"
}

// 根据ID查询是否存在
func (MessageInfo) SearchMessageByID(id string) int  {
	if err :=  utils.DB.Where("messageId = ?",id).First(&MessageInfo{}).Error; err != nil {
		return 1
	}else {
		return 0
	}
}


// 根据ID删除留言
func (MessageInfo) DeleteMessageByID(id string) int  {
	if err :=  utils.DB.Where("messageId = ?",id).Delete(&MessageInfo{}).Error; err != nil {
		return 1
	}else {
		return 0
	}
}

// 新增留言
func (MessageInfo) Insert(json map[string]interface{}) int  {
	if err :=  utils.DB.Model(&MessageInfo{}).Create(json).Error; err != nil {
		return 1
	}else {
		return 0
	}
}

// 分页查询留言
func (MessageInfo) ShowmessageByPage(currentPage int64,pageSize int64) interface{}  {
	offset := (currentPage - 1) * pageSize
	var count int64
	messageInfo := []MessageInfo{}
	// 查总数
	if err := utils.DB.Find(&messageInfo).Count(&count).Error; err != nil {
		return 1
	}
	if err :=  utils.DB.Offset(int(offset)).Limit(int(pageSize)).Find(&messageInfo).Error; err != nil {
		return 1
	}else {
		totalPage := utils.ReturnTotalPage(count , pageSize)
		return map[string]interface{}{
			"currentPage": currentPage,
			"list": messageInfo,
			"pageSize": pageSize,
			"total": count,
			"totalPage": totalPage,
		}
	}
}

// 根据留言id 回复留言  更新留言
func (MessageInfo) UpdateMessageReply(json map[string]interface{}) int  {
	if err :=  utils.DB.Model(&MessageInfo{}).Where("messageId = ?",json["messageId"]).Updates(json).Error; err != nil {
		return 1
	}else {
		return 0
	}
}