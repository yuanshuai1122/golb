package servise

import (
	"GIN/models"
)

type TextServise struct{}

/**
 * 根据id删除链接
 * @author lizibin
 */
func (TextServise) Delete(id string) int {
	isHave := models.TextInfo{}.SearchUrlByID(id)
	if isHave == 0 {
		// id存在去删除
		result := models.TextInfo{}.DeleteUrlByID(id)
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
* 模块名: DeleteAll
* 代码描述: 批量删除
* 作者:lizibin
* 创建时间:2023/01/02 15:13:43
 */
func (TextServise) DeleteAll(id string) int {
	result := models.TextInfo{}.DeleteAll(id)
	if result == 0 {
		return 0
	} else {
		return 1
	}

}

/**
 * 新增
 * @author lizibin
 */
func (TextServise) Insert(json map[string]interface{}) int {
	// 去新增
	result := models.TextInfo{}.Insert(json)
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
func (TextServise) AuditScreening(currentPage int64, pageSize int64) interface{} {
	// 去查询
	result := models.TextInfo{}.AuditScreening(currentPage, pageSize)
	if result == 1 {
		return 1
	} else {
		return result
	}
}
