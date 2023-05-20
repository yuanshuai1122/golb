package servise

import (
	"GIN/models"
)

type IpServise struct{}

/**
 * 根据id删除链接
 * @author lizibin
 */
func (IpServise) Delete(id string) int {
	isHave := models.IpInfo{}.SearchUrlByID(id)
	if isHave == 0 {
		// id存在去删除
		result := models.IpInfo{}.DeleteUrlByID(id)
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
func (IpServise) DeleteAll(id string) int {
	result := models.IpInfo{}.DeleteAll(id)
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
func (IpServise) Insert(json map[string]interface{}) int {
	// 去新增
	result := models.IpInfo{}.Insert(json)
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
func (IpServise) GetPage(currentPage int64, pageSize int64) interface{} {
	// 去查询
	result := models.IpInfo{}.GetPage(currentPage, pageSize)
	if result == 1 {
		return 1
	} else {
		return result
	}
}
