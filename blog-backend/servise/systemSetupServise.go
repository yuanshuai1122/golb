package servise

import (
	"GIN/models"
)

type SystemSetupServise struct{}

/**
 * 查询系统设置
 * @author lizibin
 */
func (SystemSetupServise) ShowAllSystemSetup() interface{} {
		result := models.SystemSetupInfo{}.ShowAllSystemSetup()
		if result == 1 {
			return 1
		} else {
			return result
		}
}

/**
 * 跟新系统设置
 * @author lizibin
 */
func (SystemSetupServise) UpdateSystemSetup(json map[string]interface{}) int {
	// 去新增
	result := models.SystemSetupInfo{}.UpdateSystemSetup(json)
	if result == 0 {
		return 0
	} else {
		return 1
	}
}
