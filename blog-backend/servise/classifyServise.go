package servise

import (
	"GIN/models"
)
type ClassifyServise struct{}
/**
 * 根据id删除
 * @author lizibin
 */
func (ClassifyServise) Delete(id string) int {
	isHave := models.ClassifyInfo{}.SearchClassifyHaveActive(id)
	if isHave == 0 {
		// 不存在文章去删除
		result := models.ClassifyInfo{}.DeleteClassifyByID(id)
		if result == 0 {
			return 0
		} else {
			return 1
		}
	} else if isHave == 1{
		return 1
	} else {
		return 2
	}
}

/**
 * 新增
 * @author lizibin
 */
func (ClassifyServise) Insert(json map[string]interface{}) int {
	// 去新增
	result := models.ClassifyInfo{}.Insert(json)
	if result == 0 {
		return 0
	} else {
		return 1
	}
}


/**
 * 查询
 * @author lizibin
 */
func (ClassifyServise) ShowAllClassifyInfo() interface{} {
	// 去查询
	result := models.ClassifyInfo{}.ShowAllClassifyInfo()
	if result == 1 {
		return 1
	} else {
		return result
	}
}


/**
 * 查询详细信息
 * @author lizibin
 */
func (ClassifyServise) ShowClassifyInfoByClassifyId(id string) interface{} {
	// 去查询
	result := models.ClassifyInfo{}.ShowClassifyInfoByClassifyId(id)
	if result == 1 {
		return 1
	} else {
		return result
	}
}

/**
 * 文章加1
 * @author lizibin
 */
func (ClassifyServise) UpdateArticleNumberByClassifyId(id string) interface{} {
	// 去查询
	result := models.ClassifyInfo{}.UpdateArticleNumberByClassifyId(id)
	if result == 1 {
		return 1
	} else {
		return 0
	}
}

/**
 * 文章-1
 * @author lizibin
 */
func (ClassifyServise) UpdateArticleNumberByClassifyId2(id string) interface{} {
	// 去查询
	result := models.ClassifyInfo{}.UpdateArticleNumberByClassifyId2(id)
	if result == 1 {
		return 1
	} else {
		return 0
	}
}
/**
 * 文章-1
 * @author lizibin
 */
func (ClassifyServise) Updata(json map[string]interface{}) interface{} {
	// 去查询
	result := models.ClassifyInfo{}.Update(json)
	if result == 1 {
		return 1
	} else {
		return 0
	}
}