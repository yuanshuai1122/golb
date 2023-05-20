package models

import "GIN/utils"

// 返回列表使用
type ClassifyInfo struct {
	ClassifyName      string `gorm:"column:classifyName" json:"classifyName"`
	ClassifyIntroduce string `gorm:"column:classifyIntroduce" json:"classifyIntroduce"`
	Color1            string `gorm:"column:color1" json:"color1"`
	Color2            string `gorm:"column:color2" json:"color2"`
	ClassifyId        int    `gorm:"column:classifyId" json:"classifyId"`
	ArticleNumber     int    `gorm:"column:articleNumber" json:"articleNumber"`
}

func (ClassifyInfo) TableName() string {
	return "classifyinfo"
}

// 根据ID查询是否存在
func (ClassifyInfo) SearchClassifyByID(id string) int {
	if err := utils.DB.Where("classifyId = ?", id).First(&ClassifyInfo{}).Error; err != nil {
		return 1
	} else {
		return 0
	}
}

// 根据ID查询是否有文章
func (ClassifyInfo) SearchClassifyHaveActive(id string) int {
	classifyInfo := ClassifyInfo{}
	if err := utils.DB.Where("classifyId = ?", id).First(&classifyInfo).Error; err != nil {
		return 1
	} else {
		if classifyInfo.ArticleNumber != 0 {
			//	有文章 不可删除
			return 2
		} else {
			return 0
		}
	}
}

// 根据ID删除
func (ClassifyInfo) DeleteClassifyByID(id string) int {
	if err := utils.DB.Where("classifyId = ?", id).Delete(&ClassifyInfo{}).Error; err != nil {
		return 1
	} else {
		return 0
	}
}

// 新增
func (ClassifyInfo) Insert(json map[string]interface{}) int {
	if err := utils.DB.Model(&ClassifyInfo{}).Create(json).Error; err != nil {
		return 1
	} else {
		return 0
	}
}

// 查询
func (ClassifyInfo) ShowAllClassifyInfo() interface{} {
	classifyInfo := []ClassifyInfo{}
	if err := utils.DB.Find(&classifyInfo).Error; err != nil {
		return 1
	} else {

		return classifyInfo
	}
}

// 查询详细
func (ClassifyInfo) ShowClassifyInfoByClassifyId(id string) interface{} {
	classifyInfo := ClassifyInfo{}
	if err := utils.DB.Where("classifyId", id).First(&classifyInfo).Error; err != nil {
		return 1
	} else {
		return classifyInfo
	}
}

// +1
func (ClassifyInfo) UpdateArticleNumberByClassifyId(id string) interface{} {
	classifyInfo := ClassifyInfo{}
	if err := utils.DB.Where("classifyId", id).First(&classifyInfo).Update("articleNumber", classifyInfo.ArticleNumber+1).Error; err != nil {
		return 1
	} else {
		return 0
	}
}

// -1
func (ClassifyInfo) UpdateArticleNumberByClassifyId2(id string) interface{} {
	classifyInfo := ClassifyInfo{}
	if err := utils.DB.Where("classifyId", id).First(&classifyInfo).Update("articleNumber", classifyInfo.ArticleNumber-1).Error; err != nil {
		return 1
	} else {
		return 0
	}
}

// 修改
func (ClassifyInfo) Update(json map[string]interface{}) interface{} {
	if err := utils.DB.Model(&ClassifyInfo{}).Where("classifyId = ?", json["classifyId"]).Updates(json).Error; err != nil {
		return 1
	} else {
		return 0
	}
}
