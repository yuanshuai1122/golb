package servise

import (
	"GIN/models"
	"fmt"
	_ "log"
	"strconv"
	_ "strconv"
)

type CommentServise struct{}

/**
 * 后台评论管理：分页查询  + 搜索框内容  管理员获取全部  普通用户获取自己的 !!! 未完成 ！！！ 没用返回用户名和文章名
 * @author lizibin
 */
func (CommentServise) Findby(currentPage int64, pageSize int64, userType string, userId string, searchContent string) interface{} {
	if userType == "0" {
		// 管理员查全部
		result := models.CommentInfo{}.SearchAll(currentPage, pageSize, searchContent)
		if result == 1 {
			return 1
		} else {
			// 循环去查文章名和用户名添加进去
			for _, v := range result.(map[string]interface{}) {
				if rec, ok := v.([]models.CommentInfo); ok {
					for _, val := range rec {
						fmt.Println("--------------",val.ArticleId)
						val.ArticleTitle = models.ArticleInfo{}.SearchNameByArticleID(strconv.Itoa(val.ArticleId)).(string)
					}
				}
				fmt.Println("result",result)
			}
			return result
		}
	} else {
		// 普通用户获取自己的
		result := models.CommentInfo{}.SearchByUserID(currentPage, pageSize, userId, searchContent)
		if result == 1 {
			return 1
		} else {
			return result
		}
	}
}

