package servise

import (
	"GIN/models"
)

type MessageServise struct{}

/**
 * 根据id删除留言
 * @author lizibin
 */
func (MessageServise) Delete(id string) int {
	isHave := models.MessageInfo{}.SearchMessageByID(id)
	if isHave == 0 {
		// id存在去删除
		result := models.MessageInfo{}.DeleteMessageByID(id)
		if result == 0 {
			return 0
		} else {
			return 1
		}
	} else {
		// id 不存在
		return 2
	}
}

/**
 * 新增留言
 * @author lizibin
 */
func (MessageServise) Insert(json map[string]interface{}) int {
	// 去新增
	result := models.MessageInfo{}.Insert(json)
	if result == 0 {
		return 0
	} else {
		return 1
	}
}

/**
 * 分页查询留言
 * @author lizibin
 */
func (MessageServise) ShowmessageByPage(currentPage int64,pageSize int64) interface{} {
	// 去查询
	result := models.MessageInfo{}.ShowmessageByPage(currentPage,pageSize)
	if result == 1 {
		return 1
	} else {
		return result
	}
}

/**
 * 根据留言id 回复留言  更新留言
 * @author lizibin
 */
func (MessageServise) UpdateMessageReply(json map[string]interface{}) int {
	// 去查询
	result := models.MessageInfo{}.UpdateMessageReply(json)
	if result == 1 {
		return 1
	} else {
		return 0
	}
}

