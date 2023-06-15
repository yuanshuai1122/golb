package services

import "golb/golb-server/models"

type CategoryService struct{}

// GetCategoryList 获取分类列表服务
func (CategoryService) GetCategoryList() interface{} {
	// 获取文章列表
	return models.Category{}.GetCategoryList()
}
