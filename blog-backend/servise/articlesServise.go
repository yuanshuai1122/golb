package servise

import (
	"GIN/models"
	"strings"
)

type ArticleServise struct{}

/**
 * 分页查询 带参数 名称
 * @author lizibin
 */
func (ArticleServise) ShowAllPass(currentPage int64, pageSize int64, keyword string) interface{} {
	// 去查询
	result := models.ArticleInfo{}.ShowAllPass(currentPage, pageSize,keyword)
	if result == 1 {
		return 1
	} else {
		return result
	}
}

/**
 * 查审核通过的文章数量
 * @author lizibin
 */
func (ArticleServise) ShowArticleCount() interface{} {
	// 去查询
	result := models.ArticleInfo{}.ShowArticleCount()
	if result == 1 {
		return 1
	} else {
		return result
	}
}

/**
 * 查审核通过的文章
 * @author lizibin
 */
func (ArticleServise) ShowAll() interface{} {
	// 去查询
	result := models.ArticleInfo{}.ShowAll()
	if result == 1 {
		return 1
	} else {
		return result
	}
}

//Delete
func (ArticleServise) Delete(articleId string) interface{} {
	res := models.SystemSetupInfo{}.ShowAllSystemSetup()
	// 判断置顶
	if res.(models.SystemSetupInfo).StickArticle == articleId {
		return 2
	}

	// 判断左侧精选文章
	allArticle := strings.Contains(res.(models.SystemSetupInfo).AllArticle,articleId)
	if allArticle {
		return 2
	}

	// 判断左侧精选文章
	featuredArticleResult := strings.Contains(res.(models.SystemSetupInfo).FeaturedArticle,articleId)
	if featuredArticleResult {
		return 2
	}

	technologyArticleResult := strings.Contains(res.(models.SystemSetupInfo).TechnologyArticle,articleId)
	// 左侧技术专区文章
	if technologyArticleResult {
		return 2
	}

	// 右侧资源专区文章
	resourceArticle := strings.Contains(res.(models.SystemSetupInfo).ResourceArticle,articleId)
	if resourceArticle {
		return 2
	}
	// 去删除
	result := models.ArticleInfo{}.Delete(articleId)
	if result == 1 {
		return 1
	} else {
		return 0
	}
}