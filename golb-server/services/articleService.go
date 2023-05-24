package services

import "golb/golb-server/models"

type ArticleService struct{}

// GetArticlesList 查询所有文章列表
func (ArticleService) GetArticlesList() interface{} {
	// 查询所有文章
	result := models.Articles{}.GetArticlesList()
	if result == nil {
		return nil
	}
	return result
}
