package models

import (
	"GIN/utils"
)

// 返回列表使用
type SystemSetupInfo struct {
	Effects02         int    `gorm:"column:effects02" json:"effects02"`
	Effects03         int    `gorm:"column:effects03" json:"effects03"`
	Effects04         int    `gorm:"column:effects04" json:"effects04"`
	Effects01         int    `gorm:"column:effects01" json:"effects01"`
	StickArticle      string `gorm:"column:stickArticle" json:"stickArticle"`
	AllArticle        string `gorm:"column:allArticle" json:"allArticle"`
	FeaturedArticle   string `gorm:"column:featuredArticle" json:"featuredArticle"`
	TechnologyArticle string `gorm:"column:technologyArticle" json:"technologyArticle"`
	ResourceArticle   string `gorm:"column:resourceArticle" json:"resourceArticle"`
	Advertising1      string `gorm:"column:advertising1" json:"advertising1"`
	AdvertisingLink1  string `gorm:"column:advertisingLink1" json:"advertisingLink1"`
	Advertising2      string `gorm:"column:advertising2" json:"advertising2"`
	AdvertisingLink2  string `gorm:"column:advertisingLink2" json:"advertisingLink2"`
	Advertising3      string `gorm:"column:advertising3" json:"advertising3"`
	AdvertisingLink3  string `gorm:"column:advertisingLink3" json:"advertisingLink3"`
	Advertising4      string `gorm:"column:advertising4" json:"advertising4"`
	Advertising5      string `gorm:"column:advertising5" json:"advertising5"`
}

func (SystemSetupInfo) TableName() string {
	return "systemsetup"
}

// 查询系统设置
func (SystemSetupInfo) ShowAllSystemSetup() interface{} {
	systemSetupInfo := SystemSetupInfo{}
	if err := utils.DB.Find(&systemSetupInfo).Error; err != nil {
		return 1
	} else {
		return systemSetupInfo
	}
}

// 更新状态
func (SystemSetupInfo) UpdateSystemSetup(json map[string]interface{}) int {
	if err := utils.DB.Model(&SystemSetupInfo{}).Where("tableName", 1).Updates(json).Error; err != nil {
		return 1
	} else {
		return 0
	}
}
