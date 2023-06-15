package models

import "golb/golb-server/utils"

// Category 分类实体
type Category struct {
	// 主键id
	Id int `gorm:"column:id" json:"id"`
	// 分类名称
	CategoryName string `gorm:"column:category_name" json:"categoryName"`
	// 分类描述
	CategoryDesc string `gorm:"column:category_desc" json:"categoryDesc"`
}

// TableName 绑定表名
func (Category) TableName() string {
	return "category"
}

// GetCategoryList 获取分类列表
func (Category) GetCategoryList() interface{} {
	var categoryList []Category
	if err := utils.DB.Find(&categoryList).Order("id desc").Error; err != nil {
		return nil
	}
	return categoryList
}
