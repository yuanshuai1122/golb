package servise

import (
	"GIN/models"
	"GIN/utils"
)

type LoginServise struct{}

/**
 * 登录
 * @author lizibin
 */
func (LoginServise) Login(userName string, userPass string) interface{} {
	result := models.UserInfo{}.Login(userName)
	if result == 1 {
		return 1 // 无此账号
	} else {
		// 查菜单
		userType := result.(map[string]interface{})["userType"]
		dbPass := result.(map[string]interface{})["userPass"]
		if utils.BcryptCheck(userPass, dbPass.(string)) {
			//	密码正确
			menu := models.MenuInfo{}.SearchMenuByUserID(userType.(int))
			if menu != 1 {
				for k, v := range menu.([]models.MenuInfo) {
					if v.Id == 2 {
						menu.([]models.MenuInfo)[k].Children = models.MenuInfo{}.SearchMenuByParentID(userType.(int)).([]models.MenuInfo)
					} else {
						v.Children = []models.MenuInfo{}
					}
				}
				result.(map[string]interface{})["menuInfo"] = menu
				return result
			} else {
				return 3 // 查询失败
			}
		} else {
			return 2 // 密码错误
		}

	}
}

/**
 * 增加登录登录 用LoginModel
 * @author lizibin
 */
func (LoginServise) InsertLoginInfo(json map[string]interface{}) interface{} {
	result := models.LoginInfo{}.InsertLoginInfo(json)
	if result == 1 {
		return 1
	} else {
		return 0
	}
}
