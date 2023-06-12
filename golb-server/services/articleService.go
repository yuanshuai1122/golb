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

// GetArticleDetail 根据id获取文章详情
func (ArticleService) GetArticleDetail(id int64) interface{} {
	// 根据id查询文章
	result := models.Articles{}.GetArticleInfoById(id)
	if result == nil {
		return nil
	}
	return result
}

func (ArticleService) PostArticle(articleInfo models.ArticleInfo) interface{} {
	// 新增文章
	result := models.Articles{}.PostArticle(articleInfo)
	if result == nil {
		return nil
	}
	return result
}
