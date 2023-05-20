package servise

import (
	"GIN/models"
)

type UserServise struct{}

/**
 *  根据用户id删除用户    同时删除用户的所有评论以及文章
 * @author lizibin
 */
func (UserServise) Delete(id string) int {
	// 删除用户
	deleteUserResult := models.UserInfo{}.DeleteUseByUserID(id)
	// 删除文章
	deleteArticleResult := models.ArticleInfo{}.DeleteUserArticByUserID(id)
	// 删除评论
	deleteCommentResult := models.CommentInfo{}.DeleteUserCommentByUserID(id)
	if deleteUserResult == 0 && deleteArticleResult == 0 && deleteCommentResult == 0 {
		return 0
	} else {
		return 1
	}
}

/**
 * 新增
 * @author lizibin
 */
func (UserServise) Insert(json map[string]interface{}) int {
	// 去新增
	result := models.UrlInfo{}.Insert(json)
	if result == 0 {
		return 0
	} else {
		return 1
	}
}

/**
 * 分页查询 带参数 审核类型
 * @author lizibin
 */
func (UserServise) AuditScreening(currentPage int64, pageSize int64, urlPass string) interface{} {
	// 去查询
	result := models.UrlInfo{}.AuditScreening(currentPage, pageSize, urlPass)
	if result == 1 {
		return 1
	} else {
		return result
	}
}

/**
 * 分页查询 带参数 名称
 * @author lizibin
 */
func (UserServise) Findby(currentPage int64, pageSize int64, searchContent string) interface{} {
	// 去查询
	result := models.UserInfo{}.Findby(currentPage, pageSize, searchContent)
	if result == 1 {
		return 1
	} else {
		return result
	}
}

/**
 * 根据id获取用户信息
 * @author lizibin
 */
func (UserServise) ShowUserByUserId(userID string) interface{} {
	// 去查询
	result := models.UserInfo{}.ShowUserByUserId(userID)
	if result == 1 {
		return 1
	} else {
		return result
	}
}

// 查询用户数量
func (UserServise) ShowUserCount() interface{} {
	// 去查询
	result := models.UserInfo{}.ShowUserCount()
	if result == 1 {
		return 1
	} else {
		return result
	}
}

/**
 * 改变状态
 * @author lizibin
 */
func (UserServise) UpdateUrlPass(urlId string, urlPass string) int {
	// 去查询
	result := models.UrlInfo{}.UpdateUrlPass(urlId, urlPass)
	if result == 1 {
		return 1
	} else {
		return 0
	}
}

/**
 * 修改个人信息
 * @author lizibin
 */
func (UserServise) UpdateUserBasicsInfo(json map[string]interface{}) int {
	// 去新增
	result := models.UserInfo{}.UpdateUserBasicsInfo(json)
	if result == 0 {
		return 0
	} else {
		return 1
	}
}

/**
 * 修改密码
 * @author lizibin
 */
func (UserServise) UpdateUserPass(userId string,userPass string) int {
	// 去新增
	result := models.UserInfo{}.UpdateUserPass(userId,userPass)
	if result == 0 {
		return 0
	} else {
		return 1
	}
}


/**
 * 查询用户名是否重复
 * @author lizibin
 */
func (UserServise) ShowUserName(userName string) interface{} {
	// 去新增
	result := models.UserInfo{}.ShowUserName(userName)
	if result == 1 {
		return 1
		// 查询失败
	} else {
		return 0
	}
}

/**
 * 查询邮箱是否重复
 * @author lizibin
 */
func (UserServise) SearchEmail(userEmail string) interface{} {
	// 去新增
	result := models.UserInfo{}.SearchEmail(userEmail)
	if result == 1 {
		return 1
		// 查询失败
	} else {
		return 0
	}
}
/**
 * 查询邮箱是否重复
 * @author lizibin
 */
//InsertRegisterInfo
func (UserServise) InsertRegisterInfo(json map[string]interface{}) interface{} {
	// 去新增
	result := models.UserInfo{}.InsertRegisterInfo(json)
	if result == 1 {
		return 1
	} else {
		return 0
	}
}