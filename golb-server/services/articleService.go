package services

import "golb/golb-server/models"

type ArticleService struct{}

// GetArticlesList 查询所有文章列表
func (ArticleService) GetArticlesList(pageNum int64, pageSize int64, keywords string) interface{} {
	// 查询所有文章
	result := models.Articles{}.GetArticlesList(pageNum, pageSize, keywords)
	if result == nil {
		return nil
	}
	return result
}
