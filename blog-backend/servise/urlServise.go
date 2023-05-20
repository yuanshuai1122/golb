package servise

import (
	"GIN/models"
)

type UrlServise struct{}

/**
 * 根据id删除链接
 * @author lizibin
 */
func (UrlServise) Delete(id string) int {
	isHave := models.UrlInfo{}.SearchUrlByID(id)
	if isHave == 0 {
		// id存在去删除
		result := models.UrlInfo{}.DeleteUrlByID(id)
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
 * 新增
 * @author lizibin
 */
func (UrlServise) Insert(json map[string]interface{}) int {
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
func (UrlServise) AuditScreening(currentPage int64, pageSize int64, urlPass string) interface{} {
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
func (UrlServise) Findby(currentPage int64, pageSize int64, searchContent string) interface{} {
	// 去查询
	result := models.UrlInfo{}.Findby(currentPage, pageSize, searchContent)
	if result == 1 {
		return 1
	} else {
		return result
	}
}

/**
 * 根据链接类型传值
 * @author lizibin
 */
func (UrlServise) ShowAllUrlInfo(urlType string) interface{} {
	// 去查询
	result := models.UrlInfo{}.ShowAllUrlInfo(urlType)
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
func (UrlServise) UpdateUrlPass(urlId string, urlPass string) int {
	// 去查询
	result := models.UrlInfo{}.UpdateUrlPass(urlId, urlPass)
	if result == 1 {
		return 1
	} else {
		return 0
	}
}
